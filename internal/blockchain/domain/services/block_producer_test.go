package services_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
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

// TestGetLastBlock runs the BlockProducerSuite.-
func TestBlockProducerSuite(t *testing.T) {
	suite.Run(t, new(BlockProducerSuite))
}
