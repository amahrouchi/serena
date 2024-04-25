package repositories_test

import (
	"errors"
	"fmt"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/amahrouchi/serena/internal/core/tools"
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
}

// TestCreateEmptyBlock tests the CreateEmptyBlock method
func (brs *BlockRepositorySuite) TestCreateEmptyBlock() {
	brs.Run("test create empty block (no prev hash)", func() {
		var repo repositories.BlockRepositoryInterface
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&repo))
		defer app.RequireStop()

		// Create empty block
		block, err := repo.CreateEmptyBlock(lo.ToPtr("previous_hash"), models.BlockStatusPending)

		// Assert
		brs.NoError(err)
		brs.NotNil(block)
		brs.Greater(block.ID, uint(0))
		brs.Equal("previous_hash", *block.PreviousHash)
		brs.Nil(block.Hash)
		brs.Equal("{}", block.Payload)
		brs.Equal(models.BlockStatusPending, block.Status)
		brs.IsType(time.Time{}, block.CreatedAt)
	})

	brs.Run("test create empty block (with prev hash)", func() {
		var repo repositories.BlockRepositoryInterface
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&repo))
		defer app.RequireStop()

		// Create empty block
		block, err := repo.CreateEmptyBlock(nil, models.BlockStatusPending)

		// Assert
		brs.NoError(err)
		brs.NotNil(block)
		brs.Greater(block.ID, uint(0))
		brs.Nil(block.PreviousHash)
		brs.Nil(block.Hash)
		brs.Equal("{}", block.Payload)
		brs.Equal(models.BlockStatusPending, block.Status)
		brs.IsType(time.Time{}, block.CreatedAt)
	})
}

// TestGetLastBlock tests the GetActiveBlock method
func (brs *BlockRepositorySuite) TestGetActiveBlock() {
	// Test get last block (no errors)
	brs.Run("test get active block (no errors)", func() {
		var db *gorm.DB
		var repo repositories.BlockRepositoryInterface
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&db, &repo))
		defer app.RequireStop()

		// Create a block
		now := time.Now()
		db.Create(&models.Block{
			ID:           1,
			Status:       models.BlockStatusPending,
			Hash:         lo.ToPtr("hash"),
			Payload:      "{\"key\": \"value\"}",
			PreviousHash: lo.ToPtr("previous_hash"),
			CreatedAt:    now,
		})

		// Get the last block
		block, err := repo.GetPendingBlock()
		fmt.Printf("block: %v\n", block)

		brs.NoError(err)
		brs.NotNil(block)
		brs.Equal(uint(1), block.ID)
		brs.Equal(models.BlockStatusPending, block.Status)
		brs.Equal("hash", *block.Hash)
		brs.Equal("{\"key\": \"value\"}", block.Payload)
		brs.Equal("previous_hash", *block.PreviousHash)
		brs.Equal(now.Unix(), block.CreatedAt.Unix())
	})

	// Test get last block (no block)
	brs.Run("test get active block (no block)", func() {
		var db *gorm.DB
		var repo repositories.BlockRepositoryInterface
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&db, &repo))
		defer app.RequireStop()

		// Get the last block
		block, err := repo.GetActiveBlock()

		// Assert
		brs.NoError(err)
		brs.Nil(block)
	})
}

// TestCreateGenesisBlock tests the CreateGenesisBlock method
func (brs *BlockRepositorySuite) TestCreateGenesisBlock() {
	// Test create genesis block (no errors)
	brs.Run("test create genesis block (no errors)", func() {
		var repo repositories.BlockRepositoryInterface
		var db *gorm.DB
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&repo, &db))
		defer app.RequireStop()

		// Create genesis block
		block, err := repo.CreateGenesisBlock()

		// Assert
		brs.NoError(err)
		brs.NotNil(block)
		brs.Greater(block.ID, uint(0))
		brs.Equal(models.BlockStatusClosed, block.Status)
		brs.Equal("genesis", *block.PreviousHash)
		brs.NotNil(block.Hash)
		brs.Equal("{}", block.Payload)

		// Test the creation of the active and pending blocks
		var blocks []models.Block
		allBlock := db.Order("id ASC").Find(&blocks)
		brs.NoError(allBlock.Error)
		brs.Equal(3, len(blocks))
		brs.Equal(models.BlockStatusActive, blocks[1].Status)
		brs.Equal(models.BlockStatusPending, blocks[2].Status)
	})

	// Test create genesis block (fail to get time)
	brs.Run("test create genesis block (fail to get time)", func() {
		// Prepare deps to populate
		var repo repositories.BlockRepositoryInterface

		// Run the test app
		app := tests.NewTestApp(false).Run(
			brs.T(),
			fx.Populate(&repo),
			fx.Decorate(func() tools.TimeSyncInterface {
				mockTimeSync := new(tools.TimeSyncMock)
				mockTimeSync.On("Current").Return(nil, errors.New("error"))

				return mockTimeSync
			}),
		)
		defer app.RequireStop()

		// Create genesis block
		block, err := repo.CreateGenesisBlock()

		// Assert
		brs.Nil(block)
		brs.Error(err)
		brs.Equal("error", err.Error())
	})
}

// TestBlockRepositorySuite launches the test suite
func TestBlockRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlockRepositorySuite))
}
