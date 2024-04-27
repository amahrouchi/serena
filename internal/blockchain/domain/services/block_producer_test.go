package services_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TODO: fix these tests

// BlockProducerSuite is the test suite for the BlockProducer struct.
type BlockProducerSuite struct {
	suite.Suite
}

// TestGetLastBlock tests the GetActiveBlock method.
func (bps *BlockProducerSuite) TestGetActiveBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("GetActiveBlock").Return(&models.Block{}, nil)

	// Get last block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	block, err := producer.GetActiveBlock()

	// Assert
	bps.NoError(err)
	bps.NotNil(block)
}

// TestCreateEmptyBlock tests the CreateEmptyBlock method.
func (bps *BlockProducerSuite) TestCreateEmptyBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("CreateEmptyBlock", nil, models.BlockStatusPending).Return(&models.Block{}, nil)

	// Create empty block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	block, err := producer.CreateEmptyBlock(nil, models.BlockStatusPending)

	// Assert
	bps.NoError(err)
	bps.NotNil(block)
}

// TestCreateGenesisBlock tests the CreateGenesisBlock method.
func (bps *BlockProducerSuite) TestCreateGenesisBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("CreateGenesisBlock").Return(&models.Block{}, nil)

	// Create genesis block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	block, err := producer.CreateGenesisBlock()

	// Assert
	bps.NoError(err)
	bps.NotNil(block)
}

// TestSwitchActiveBlock tests the SwitchActiveBlock method.
func (bps *BlockProducerSuite) TestSwitchActiveBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("SwitchActiveBlock").Return(nil)

	// Switch active block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	err := producer.SwitchActiveBlock()

	// Assert
	bps.NoError(err)
}

// TestGetLastBlock runs the BlockProducerSuite.-
func TestBlockProducerSuite(t *testing.T) {
	suite.Run(t, new(BlockProducerSuite))
}
