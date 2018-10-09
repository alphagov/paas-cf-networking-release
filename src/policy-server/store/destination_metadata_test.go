package store_test

import (
	"errors"
	"policy-server/store"

	dbfakes "code.cloudfoundry.org/cf-networking-helpers/db/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DestinationMetadata", func() {
	var (
		tx                       *dbfakes.Transaction
		destinationMetadataTable *store.DestinationMetadataTable
	)

	BeforeEach(func() {
		tx = new(dbfakes.Transaction)

		destinationMetadataTable = &store.DestinationMetadataTable{}
	})

	Context("when the db fails to insert", func() {
		Context("on mysql", func() {
			BeforeEach(func() {
				tx.DriverNameReturns("mysql")
				tx.ExecReturns(nil, errors.New("failed to insert"))
			})

			It("returns an error", func() {
				_, err := destinationMetadataTable.Create(tx, "term-guid", "some-name", "some-desc")
				Expect(err).To(MatchError("failed to insert"))
			})
		})
	})

	Context("when the db fails to update", func() {
		It("returns the error", func() {
			tx.ExecReturns(nil, errors.New("not right now"))
			err := destinationMetadataTable.Update(tx, "term-guid", "name", "desc")
			Expect(err).To(MatchError("not right now"))
		})
	})

	Context("update", func() {
		It("passes an error from RowsAffected if RowsAffected fails", func() {
			tx.ExecReturns(&fakeSqlResultWithError{}, nil)
			err := destinationMetadataTable.Update(tx, "term-guid", "name", "desc")
			Expect(err).To(MatchError("not right now"))
		})

		It("passes an error from Exec if Exec fails", func() {
			tx.ExecReturns(&fakeSqlResultWithError{}, errors.New("bigger error"))
			err := destinationMetadataTable.Update(tx, "term-guid", "name", "desc")
			Expect(err).To(MatchError("bigger error"))
		})

		It("returns error when zero rows affected", func() {
			tx.ExecReturns(&fakeSqlResultWithZeroRowsAffected{}, nil)
			err := destinationMetadataTable.Update(tx, "term-guid", "name", "desc")
			Expect(err).To(MatchError("destination GUID not found"))
		})
	})
})
