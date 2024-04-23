package repositories

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/stretchr/testify/mock"
)

// BlockRepositoryMock is a repository for blocks
type BlockRepositoryMock struct {
	mock.Mock
}

func (brm *BlockRepositoryMock) CreateEmptyBlock() error {
	args := brm.Called()
	return args.Error(0)
}

func (brm *BlockRepositoryMock) GetLastBlock() (*models.Block, error) {
	args := brm.Called()
	return args.Get(0).(*models.Block), args.Error(1)
}

func (brm *BlockRepositoryMock) CreateGenesisBlock() (*models.Block, error) {
	args := brm.Called()
	return args.Get(0).(*models.Block), args.Error(1)
}
