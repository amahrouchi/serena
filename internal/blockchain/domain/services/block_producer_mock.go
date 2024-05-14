package services

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/stretchr/testify/mock"
)

// BlockProducerMock is a mock for a block producer.
type BlockProducerMock struct {
	mock.Mock
}

func (bpm *BlockProducerMock) GetActiveBlock() (*models.Block, error) {
	args := bpm.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Block), args.Error(1)
}

func (bpm *BlockProducerMock) CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error) {
	if prevHash == nil {
		args := bpm.Called(nil, status)
		return args.Get(0).(*models.Block), args.Error(1)
	}

	args := bpm.Called(*prevHash, status)
	return args.Get(0).(*models.Block), args.Error(1)
}

func (bpm *BlockProducerMock) CreateGenesisBlock() (*models.Block, error) {
	args := bpm.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Block), args.Error(1)
}

func (bpm *BlockProducerMock) SwitchActiveBlock() error {
	args := bpm.Called()
	return args.Error(0)
}
