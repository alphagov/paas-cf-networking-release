package store_test

import (
	"fmt"
	"policy-server/db"
	"policy-server/store"
	testhelpers "test-helpers"
	"time"

	dbHelper "code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport"
	"code.cloudfoundry.org/lager"

	uuid "github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EgressDestination", func() {
	var (
		dbConf dbHelper.Config
		realDb *db.ConnWrapper

		terminalsTable         *store.TerminalsTable
		egressDestinationTable *store.EgressDestinationTable

		terminalId string
	)

	BeforeEach(func() {
		dbConf = testsupport.GetDBConfig()
		dbConf.DatabaseName = fmt.Sprintf("egress_destination_test_node_%d", time.Now().UnixNano())
		dbConf.Timeout = 30
		testhelpers.CreateDatabase(dbConf)

		logger := lager.NewLogger("Egress Destination Test")
		realDb = db.NewConnectionPool(dbConf, 200, 200, 5*time.Minute, "Egress Destination Test", "Egress Destination Test", logger)

		migrate(realDb)

		egressDestinationTable = &store.EgressDestinationTable{}

		terminalsTable = &store.TerminalsTable{
			Guids: &store.GuidGenerator{},
		}

		tx, err := realDb.Beginx()
		Expect(err).NotTo(HaveOccurred())

		terminalId, err = terminalsTable.Create(tx)
		Expect(err).NotTo(HaveOccurred())

		_, err = egressDestinationTable.CreateIPRange(tx, terminalId, "1.1.1.1", "2.2.2.2", "tcp", 8080, 8081, -1, -1)
		Expect(err).NotTo(HaveOccurred())

		err = tx.Commit()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if realDb != nil {
			Expect(realDb.Close()).To(Succeed())
		}
		testhelpers.RemoveDatabase(dbConf)
	})

	Context("when a destination metadata doesn't exist for destination", func() {
		It("returns empty strings for name/description", func() {
			tx, err := realDb.Beginx()
			Expect(err).NotTo(HaveOccurred())

			destinations, err := egressDestinationTable.All(tx)
			Expect(err).NotTo(HaveOccurred())

			err = tx.Rollback()
			Expect(err).NotTo(HaveOccurred())

			_, err = uuid.ParseHex(destinations[0].GUID)
			Expect(err).NotTo(HaveOccurred())

			Expect(destinations[0].GUID).To(Equal(terminalId))
			Expect(destinations[0].Name).To(Equal(""))
			Expect(destinations[0].Description).To(Equal(""))
			Expect(destinations[0].Protocol).To(Equal("tcp"))
			Expect(destinations[0].IPRanges).To(Equal([]store.IPRange{{Start: "1.1.1.1", End: "2.2.2.2"}}))
			Expect(destinations[0].Ports).To(Equal([]store.Ports{{Start: 8080, End: 8081}}))
			Expect(destinations[0].ICMPType).To(Equal(-1))
			Expect(destinations[0].ICMPCode).To(Equal(-1))
		})
	})

	FContext("GetIpRange by GUID", func() {
		It("return IpRange for given GUID", func () {
			tx, err := realDb.Beginx()
			Expect(err).NotTo(HaveOccurred())

			destinations, err := egressDestinationTable.All(tx)
			Expect(err).NotTo(HaveOccurred())

			destination, err := egressDestinationTable.GetIPRange(tx ,destinations[0].GUID)
			Expect(err).NotTo(HaveOccurred())
			Expect(destination.GUID).To(Equal(terminalId))
			Expect(destination.Name).To(Equal(""))
			Expect(destination.Description).To(Equal(""))
			Expect(destination.Protocol).To(Equal("tcp"))
			Expect(destination.IPRanges).To(Equal([]store.IPRange{{Start: "1.1.1.1", End: "2.2.2.2"}}))
			Expect(destination.Ports).To(Equal([]store.Ports{{Start: 8080, End: 8081}}))
			Expect(destination.ICMPType).To(Equal(-1))
			Expect(destination.ICMPCode).To(Equal(-1))
		})

		It("does not error for unknown GUID", func(){
			tx, err := realDb.Beginx()
			Expect(err).NotTo(HaveOccurred())

			destination, err := egressDestinationTable.GetIPRange(tx ,"dfsfsdfsfsdf4534534535345")
			Expect(err).NotTo(HaveOccurred())
			Expect(destination).To(Equal(store.EgressDestination{}))
		})
	})

	Context("DeleteIpRange", func() {
		Context("when the destination does not exist", func() {
			It("does not error", func() {
				tx, err := realDb.Beginx()
				Expect(err).NotTo(HaveOccurred())

				err = egressDestinationTable.DeleteIPRange(tx, "dfsfsdfsfsdf4534534535345")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the destination does exist", func() {
			It("deletes the destination", func() {
				tx, err := realDb.Beginx()
				Expect(err).NotTo(HaveOccurred())

				destinations, err := egressDestinationTable.All(tx)
				Expect(err).NotTo(HaveOccurred())

				err = egressDestinationTable.DeleteIPRange(tx ,destinations[0].GUID)
				Expect(err).NotTo(HaveOccurred())

				destinations, err = egressDestinationTable.All(tx)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(destinations)).To(Equal(0))


			})
		})
	})
})
