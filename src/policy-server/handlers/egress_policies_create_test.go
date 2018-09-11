package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"policy-server/handlers"
	"policy-server/handlers/fakes"
	"policy-server/uaa_client"

	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"policy-server/store"
)

var _ = Describe("EgressPoliciesCreate", func() {
	var (
		expectedStoreEgressPolicies []store.EgressPolicy

		fakeMapper *fakes.EgressPolicyMapper
		fakeStore  *fakes.EgressPolicyStore
		logger     *lagertest.TestLogger

		handler *handlers.EgressPolicyCreate

		responseWriter *httptest.ResponseRecorder
		request        *http.Request
		requestBody    string
		tokenData      uaa_client.CheckTokenResponse
	)

	BeforeEach(func() {
		fakeStore = &fakes.EgressPolicyStore{}
		fakeMapper = &fakes.EgressPolicyMapper{}

		handler = &handlers.EgressPolicyCreate{
			Store:  fakeStore,
			Mapper: fakeMapper,
		}

		logger = lagertest.NewTestLogger("test")

		var err error
		requestBody = `{
			"egress_policies": [
				{
					"source": { "id": "AN-APP-GUID", "type": "app" },
					"destination": {"id": "A-DEST-GUID" }
				}
			]
		}`

		expectedStoreEgressPolicies = []store.EgressPolicy{
			{
				Source:      store.EgressSource{ID: "AN-APP-GUID"},
				Destination: store.EgressDestination{GUID: "A-DEST-GUID"},
			},
		}

		fakeMapper.AsStoreEgressPolicyReturns(expectedStoreEgressPolicies, nil)

		request, err = http.NewRequest("POST", "/networking/v1/external/egress_policies", bytes.NewBuffer([]byte(requestBody)))
		Expect(err).NotTo(HaveOccurred())

		responseWriter = httptest.NewRecorder()

		tokenData = uaa_client.CheckTokenResponse{
			Scope: []string{"some-scope"},
		}

	})

	Describe("ServeHTTP", func() {
		It("creates an egress policy", func() {
			MakeRequestWithLoggerAndAuth(handler.ServeHTTP, responseWriter, request, logger, tokenData)

			Expect(fakeMapper.AsStoreEgressPolicyCallCount()).To(Equal(1))
			policies := fakeMapper.AsStoreEgressPolicyArgsForCall(0)
			Expect(string(policies)).To(Equal(requestBody))

			Expect(fakeStore.CreateCallCount()).To(Equal(1))
			storePolicies := fakeStore.CreateArgsForCall(0)
			Expect(storePolicies).To(Equal(expectedStoreEgressPolicies))
		})
	})
})
