package repositories_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"testing"
	"time"
)

// BlockRepositorySuite the test suite for the BlockRepository struct
type BlockRepositorySuite struct {
	suite.Suite

	repositories.BlockRepository
}

// TestGetLastBlock tests the GetLastBlock method
func (brs *BlockRepositorySuite) TestGetLastBlock() {
	var db *gorm.DB
	var repo repositories.BlockRepositoryInterface
	tests.NewTestApp(false).Run(brs.T(), fx.Populate(&db, &repo))

	// Create a block
	now := time.Now()
	db.Create(&models.Block{
		ID:           1,
		Hash:         lo.ToPtr("hash"),
		Payload:      "{\"key\": \"value\"}",
		PreviousHash: "previous_hash",
		CreatedAt:    now,
	})

	// Get the last block
	block, err := repo.GetLastBlock()

	brs.NoError(err)
	brs.NotNil(block)
	brs.Equal(uint(1), block.ID)
	brs.Equal("hash", *block.Hash)
	brs.Equal("{\"key\": \"value\"}", block.Payload)
	brs.Equal("previous_hash", block.PreviousHash)
	brs.Equal(now.Unix(), block.CreatedAt.Unix())
}

func (brs *BlockRepositorySuite) TestCreateGenesisBlock() {
	// TODO: test for real
}

// TestBlockRepositorySuite launches the test suite
func TestBlockRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlockRepositorySuite))
}
