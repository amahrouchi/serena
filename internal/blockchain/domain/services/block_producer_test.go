package services_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// BlockProducerSuite is the test suite for the BlockProducer struct.
type BlockProducerSuite struct {
	suite.Suite
}

// TestCalculateHash tests the CalculateHash method.
func (bps *BlockProducerSuite) TestCalculateHash() {
	producer := services.NewBlockProducer(nil, nil)

	// Calculate hash
	hash := producer.CalculateHash(nil)

	// Assert
	bps.Equal("", hash)
}

// TestGetLastBlock tests the GetLastBlock method.
func (bps *BlockProducerSuite) TestGetLastBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("GetLastBlock").Return(&models.Block{}, nil)

	// Get last block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	block, err := producer.GetLastBlock()

	// Assert
	bps.NoError(err)
	bps.NotNil(block)
}

// TestProduceBlock tests the CreateEmptyBlock method.
func (bps *BlockProducerSuite) TestProduceBlock() {
	repo := new(repositories.BlockRepositoryMock)
	repo.On("CreateEmptyBlock").Return(nil)

	// Produce block
	producer := services.NewBlockProducer(repo, tests.NewEmptyLogger())
	producer.CreateEmptyBlock()

	// Assert
	repo.AssertExpectations(bps.T())
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

// TestGetLastBlock runs the BlockProducerSuite.-
func TestBlockProducerSuite(t *testing.T) {
	suite.Run(t, new(BlockProducerSuite))
}
