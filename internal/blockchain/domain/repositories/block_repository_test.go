package repositories_test

import (
	"errors"
	"fmt"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
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

	brs.Run("test create empty block (fail to get time)", func() {
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

		// Create empty block
		block, err := repo.CreateEmptyBlock(nil, models.BlockStatusPending)

		// Assert
		brs.Nil(block)
		brs.Error(err)
		brs.Equal("error", err.Error())
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

// TestSwitchActiveBlock tests the SwitchActiveBlock method
func (brs *BlockRepositorySuite) TestSwitchActiveBlock() {
	brs.Run("test switch (no errors)", func() {
		// Run the test app
		var repo repositories.BlockRepositoryInterface
		var db *gorm.DB
		app := tests.NewTestApp(false).Run(
			brs.T(),
			fx.Populate(&repo, &db),
		)
		defer app.RequireStop()

		// Create test data
		activeBlock := &models.Block{
			Status:       models.BlockStatusActive,
			Hash:         nil,
			PreviousHash: lo.ToPtr("active_previous_hash"),
			Payload:      "{}",
			CreatedAt:    time.Now(),
		}
		db.Create(activeBlock)

		pendingBlock := &models.Block{
			Status:       models.BlockStatusPending,
			Hash:         nil,
			PreviousHash: nil,
			Payload:      "{}",
			CreatedAt:    time.Now(),
		}
		db.Create(pendingBlock)

		// Switch active block
		err := repo.SwitchActiveBlock()

		// Load block data from updated db
		var closedBlock, newActiveBlock models.Block
		db.Find(&closedBlock, activeBlock.ID)
		db.Find(&newActiveBlock, pendingBlock.ID)
		newPendingBlock, _ := repo.GetPendingBlock()

		// Assert closed block
		brs.NoError(err)
		brs.Equal(models.BlockStatusClosed, closedBlock.Status)
		brs.NotNil(closedBlock.Hash)
		brs.NotNil(closedBlock.PreviousHash)

		// Assert active block
		brs.Equal(models.BlockStatusActive, newActiveBlock.Status)
		brs.Equal(closedBlock.Hash, newActiveBlock.PreviousHash)
		brs.Nil(newActiveBlock.Hash)
		brs.Equal(newActiveBlock.Payload, "{}")

		// Assert pending block
		brs.NotEqual(pendingBlock.ID, newPendingBlock.ID)
		brs.Equal(models.BlockStatusPending, newPendingBlock.Status)
		brs.Equal(newPendingBlock.Payload, "{}")
		brs.Nil(newPendingBlock.Hash)
		brs.Nil(newPendingBlock.PreviousHash)
	})

	brs.Run("test switch (no pending block)", func() {
		var repo repositories.BlockRepositoryInterface
		var db *gorm.DB
		app := tests.NewTestApp(false).Run(
			brs.T(),
			fx.Populate(&repo, &db),
		)
		defer app.RequireStop()

		// Create test data
		activeBlock := &models.Block{
			Status:       models.BlockStatusActive,
			Hash:         nil,
			PreviousHash: lo.ToPtr("active_previous_hash"),
			Payload:      "{}",
			CreatedAt:    time.Now(),
		}
		db.Create(activeBlock)

		// Switch active block
		err := repo.SwitchActiveBlock()

		// Load block data from updated db
		var closedBlock models.Block
		db.Find(&closedBlock, activeBlock.ID)
		newPendingBlock, _ := repo.GetPendingBlock()
		newActiveBlock, _ := repo.GetActiveBlock()

		// Assert closed block
		brs.NoError(err)
		brs.Equal(models.BlockStatusClosed, closedBlock.Status)
		brs.NotNil(closedBlock.Hash)
		brs.NotNil(closedBlock.PreviousHash)

		// Assert active block
		brs.Equal(models.BlockStatusActive, newActiveBlock.Status)
		brs.Equal(closedBlock.Hash, newActiveBlock.PreviousHash)
		brs.Nil(newActiveBlock.Hash)
		brs.Equal(newActiveBlock.Payload, "{}")

		// Assert pending block
		brs.Equal(models.BlockStatusPending, newPendingBlock.Status)
		brs.Equal(newPendingBlock.Payload, "{}")
		brs.Nil(newPendingBlock.Hash)
		brs.Nil(newPendingBlock.PreviousHash)
	})

	brs.Run("test switch (fail to hash)", func() {
		{
			// Context data
			hashErr := errors.New("hash error")

			// Run the test app
			var repo repositories.BlockRepositoryInterface
			var db *gorm.DB
			app := tests.NewTestApp(false).Run(
				brs.T(),
				fx.Populate(&repo, &db),
				fx.Decorate(func() repositories.HashGenInterface {
					// duplicate active block
					hashGen := repositories.HashGenMock{}
					hashGen.On("FromBlock", mock.AnythingOfType("*models.Block")).
						Return("", hashErr)

					return &hashGen
				}),
			)
			defer app.RequireStop()

			// Create test data
			activeBlock := &models.Block{
				Status:       models.BlockStatusActive,
				Hash:         nil,
				PreviousHash: lo.ToPtr("active_previous_hash"),
				Payload:      "{}",
				CreatedAt:    time.Now(),
			}
			db.Create(activeBlock)

			// Switch active block
			err := repo.SwitchActiveBlock()

			// Assert
			activeBlockAfterFailedSwitch, _ := repo.GetActiveBlock()
			brs.ErrorIs(err, hashErr)
			brs.Equal(activeBlock.ID, activeBlockAfterFailedSwitch.ID)
		}
	})
}

func (brs *BlockRepositorySuite) TestUpdate() {

	brs.Run("test update (no errors)", func() {
		// Run the test app
		var repo repositories.BlockRepositoryInterface
		var db *gorm.DB
		app := tests.NewTestApp(false).Run(
			brs.T(),
			fx.Populate(&repo, &db),
		)
		defer app.RequireStop()

		// Create test data
		block := &models.Block{
			ID:           1,
			Status:       models.BlockStatusActive,
			Hash:         nil,
			PreviousHash: lo.ToPtr("active_previous_hash"),
			Payload:      "{}",
			CreatedAt:    time.Now(),
		}

		// Update the block
		block.Status = models.BlockStatusClosed
		err := repo.Update(block)
		brs.NoError(err)
	})

	brs.Run("test update (duplicate key error)", func() {
		// Run the test app
		var repo repositories.BlockRepositoryInterface
		var db *gorm.DB
		app := tests.NewTestApp(false).Run(
			brs.T(),
			fx.Populate(&repo, &db),
		)
		defer app.RequireStop()

		// Create test data
		block := &models.Block{
			ID:           1,
			Status:       models.BlockStatusActive,
			Hash:         nil,
			PreviousHash: lo.ToPtr("active_previous_hash"),
			Payload:      "{}",
			CreatedAt:    time.Now(),
		}

		// Update the block
		block.Status = models.BlockStatusClosed
		err := repo.Update(block)
		brs.NoError(err)

		// Change the block ID to trigger a duplicate key error
		block.ID = 2
		err2 := repo.Update(block)
		brs.Error(err2)
	})
}

// TestBlockRepositorySuite launches the test suite
func TestBlockRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlockRepositorySuite))
}
