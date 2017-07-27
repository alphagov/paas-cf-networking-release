package integration_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"policy-server/config"
	"policy-server/integration/helpers"
	"strings"

	"code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/cf-networking-helpers/metrics"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Policy Cleanup", func() {
	var (
		sessions          []*gexec.Session
		conf              config.Config
		policyServerConfs []config.Config
		dbConf            db.Config
		headers           map[string]string

		fakeMetron metrics.FakeMetron
	)

	BeforeEach(func() {
		headers = map[string]string{"Accept": "1.0.0"}
		fakeMetron = metrics.NewFakeMetron()

		dbConf = testsupport.GetDBConfig()
		dbConf.DatabaseName = fmt.Sprintf("test_node_%d", GinkgoParallelNode())
		testsupport.CreateDatabase(dbConf)

		template := helpers.DefaultTestConfig(dbConf, fakeMetron.Address(), "fixtures")
		template.CleanupInterval = 1
		template.CCAppRequestChunkSize = 1

		policyServerConfs = configurePolicyServers(template, 2)
		sessions = startPolicyServers(policyServerConfs)
		conf = policyServerConfs[0]
	})

	AfterEach(func() {
		for _, session := range sessions {
			session.Interrupt()
			Eventually(session, helpers.DEFAULT_TIMEOUT).Should(gexec.Exit())
		}

		testsupport.RemoveDatabase(dbConf)

		Expect(fakeMetron.Close()).To(Succeed())
	})

	Describe("Cleanup policies endpoint", func() {
		BeforeEach(func() {
			body := strings.NewReader(`{ "policies": [
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 8080, "end": 8080 } } },
				{"source": { "id": "live-app-2-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 9999, "end": 9999 } } },
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "dead-app", "protocol": "tcp", "ports": { "start": 3333, "end": 3333 }} }
				]} `)

			resp := helpers.MakeAndDoRequest(
				"POST",
				fmt.Sprintf("http://%s:%d/networking/v0/external/policies", conf.ListenHost, conf.ListenPort),
				headers,
				body,
			)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

		})

		It("responds with a 200 and lists all stale policies", func() {
			resp := helpers.MakeAndDoRequest(
				"POST",
				fmt.Sprintf("http://%s:%d/networking/v0/external/policies/cleanup", conf.ListenHost, conf.ListenPort),
				headers,
				nil,
			)

			stalePoliciesStr := `{
				"total_policies":1,
				"policies": [
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "dead-app", "protocol": "tcp", "ports": { "start": 3333, "end": 3333 } } }
				 ]}
				`

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			Expect(bodyBytes).To(MatchJSON(stalePoliciesStr))
			Eventually(fakeMetron.AllEvents, "5s").Should(ContainElement(
				HaveName("CleanupRequestTime"),
			))
			Eventually(fakeMetron.AllEvents, "5s").Should(ContainElement(
				HaveName("StoreDeleteSuccessTime"),
			))
		})
	})

	Describe("Automatic Stale Policy Cleanup", func() {
		BeforeEach(func() {
			body := strings.NewReader(`{ "policies": [
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 8080, "end": 8080 } } },
				{"source": { "id": "live-app-2-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 9999, "end": 9999 } } },
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "dead-app", "protocol": "tcp", "ports": { "start": 3333, "end": 3333 } } }
				]} `)

			resp := helpers.MakeAndDoRequest(
				"POST",
				fmt.Sprintf("http://%s:%d/networking/v0/external/policies", conf.ListenHost, conf.ListenPort),
				headers,
				body,
			)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		It("eventually cleans up stale policies", func() {
			listPolicies := func() []byte {
				resp := helpers.MakeAndDoRequest(
					"GET",
					fmt.Sprintf("http://%s:%d/networking/v0/external/policies", conf.ListenHost, conf.ListenPort),
					headers,
					nil,
				)
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				bodyBytes, _ := ioutil.ReadAll(resp.Body)
				return bodyBytes
			}

			activePolicies := `{ "total_policies": 2,
			"policies": [
				{"source": { "id": "live-app-1-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 8080, "end": 8080 } } },
				{"source": { "id": "live-app-2-guid" }, "destination": { "id": "live-app-2-guid", "protocol": "tcp", "ports": { "start": 9999, "end": 9999 } } }
				]} `
			Eventually(listPolicies, "5s").Should(MatchJSON(activePolicies))

			By("emitting store metrics")
			Eventually(fakeMetron.AllEvents, "5s").Should(ContainElement(
				HaveName("StoreDeleteSuccessTime"),
			))
		})
	})
})
