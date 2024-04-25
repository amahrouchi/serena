package repositories_test

import (
	"errors"
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
	var repo repositories.BlockRepositoryInterface
	app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&repo))
	defer app.RequireStop()

	// Create empty block
	err := repo.CreateEmptyBlock()

	// Assert
	brs.Error(err)
	brs.Equal("not implemented", err.Error())
}

// TestGetLastBlock tests the GetActiveBlock method
func (brs *BlockRepositorySuite) TestGetLastBlock() {
	// Test get last block (no errors)
	brs.Run("test get last block (no errors)", func() {
		var db *gorm.DB
		var repo repositories.BlockRepositoryInterface
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&db, &repo))
		defer app.RequireStop()

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
		block, err := repo.GetActiveBlock()

		brs.NoError(err)
		brs.NotNil(block)
		brs.Equal(uint(1), block.ID)
		brs.Equal("hash", *block.Hash)
		brs.Equal("{\"key\": \"value\"}", block.Payload)
		brs.Equal("previous_hash", block.PreviousHash)
		brs.Equal(now.Unix(), block.CreatedAt.Unix())

	})

	// Test get last block (no block)
	brs.Run("test get last block (no block)", func() {
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
		app := tests.NewTestApp(false).Run(brs.T(), fx.Populate(&repo))
		defer app.RequireStop()

		// Create genesis block
		block, err := repo.CreateGenesisBlock()

		// Assert
		brs.NoError(err)
		brs.NotNil(block)
		brs.Greater(block.ID, uint(0))
		brs.Equal("", block.PreviousHash)
		brs.NotNil(block.Hash)
		brs.Equal("{}", block.Payload)
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
