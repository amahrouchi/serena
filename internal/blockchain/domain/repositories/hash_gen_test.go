package repositories_test

import (
	"fmt"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HashGenSuite is the test suite for the HashGen struct.
type HashGenSuite struct {
	suite.Suite
}

// TestNewHashGen tests the NewHashGen method.
func (hgs *HashGenSuite) TestNewHashGen() {
	hashGen := repositories.NewHashGen(
		tests.NewEmptyLogger(),
	)

	hgs.NotNil(hashGen)
}

// TestFromBlock tests the FromBlock method.
func (hgs *HashGenSuite) TestFromBlock() {
	hgs.Run("nil block", func() {
		hashGen := repositories.NewHashGen(
			tests.NewEmptyLogger(),
		)

		_, err := hashGen.FromBlock(nil)

		hgs.Error(err)
	})

	hgs.Run("valid block", func() {
		hashGen := repositories.NewHashGen(
			tests.NewEmptyLogger(),
		)

		hash, err := hashGen.FromBlock(&models.Block{
			ID:           1,
			PreviousHash: lo.ToPtr("previous_hash"),
			Payload:      "{\"data\":\"test\"}",
			Status:       models.BlockStatusPending,
		})

		fmt.Printf("hash: %s\n", hash)

		hgs.NoError(err)
		hgs.Equal("5f0d96cd38ace8977d488c7b3b8dcef2414f08c6c869371f269255c9af50bd10", hash)
	})
}

// TestHashGenSuite tests the HashGenSuite.
func TestHashGenSuite(t *testing.T) {
	suite.Run(t, new(HashGenSuite))
}
