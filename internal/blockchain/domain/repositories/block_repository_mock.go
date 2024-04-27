package repositories

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/stretchr/testify/mock"
)

// BlockRepositoryMock is a repository for blocks
type BlockRepositoryMock struct {
	mock.Mock
}

func (brm *BlockRepositoryMock) CreateGenesisBlock() (*models.Block, error) {
	args := brm.Called()
	return args.Get(0).(*models.Block), args.Error(1)
}

func (brm *BlockRepositoryMock) CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error) {
	if prevHash == nil {
		args := brm.Called(nil, status)
		return args.Get(0).(*models.Block), args.Error(1)
	}

	args := brm.Called(*prevHash, status)
	return args.Get(0).(*models.Block), args.Error(1)
}

func (brm *BlockRepositoryMock) SwitchActiveBlock() error {
	args := brm.Called()
	return args.Error(0)
}

func (brm *BlockRepositoryMock) GetActiveBlock() (*models.Block, error) {
	args := brm.Called()
	return args.Get(0).(*models.Block), args.Error(1)
}

func (brm *BlockRepositoryMock) GetPendingBlock() (*models.Block, error) {
	args := brm.Called()
	return args.Get(0).(*models.Block), args.Error(1)
}

func (brm *BlockRepositoryMock) Update(block *models.Block) error {
	args := brm.Called(block)
	return args.Error(0)
}

func (brm *BlockRepositoryMock) Activate(block *models.Block) error {
	args := brm.Called(block)
	return args.Error(0)
}

func (brm *BlockRepositoryMock) Close(block *models.Block) error {
	args := brm.Called(block)
	return args.Error(0)
}
