package services

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/rs/zerolog"
)

// BlockProducerInterface is an interface for a block producer.
type BlockProducerInterface interface {
	GetActiveBlock() (*models.Block, error)
	CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error)
	CreateGenesisBlock() (*models.Block, error)
	SwitchActiveBlock() error
}

// BlockProducer is responsible for producing blocks.
type BlockProducer struct {
	blockRepo repositories.BlockRepositoryInterface
	logger    *zerolog.Logger
}

// NewBlockProducer creates a new BlockProducer.
func NewBlockProducer(
	blockRepo repositories.BlockRepositoryInterface,
	logger *zerolog.Logger,
) *BlockProducer {
	return &BlockProducer{
		blockRepo: blockRepo,
		logger:    logger,
	}
}

// GetActiveBlock Gets the last created block
func (bp *BlockProducer) GetActiveBlock() (*models.Block, error) {
	return bp.blockRepo.GetActiveBlock()
}

// CreateEmptyBlock produces a block.
func (bp *BlockProducer) CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error) {
	return bp.blockRepo.CreateEmptyBlock(prevHash, status)
}

// CreateGenesisBlock create the genesis block
func (bp *BlockProducer) CreateGenesisBlock() (*models.Block, error) {
	return bp.blockRepo.CreateGenesisBlock()
}

// SwitchActiveBlock close the active block, activate the pending one and create a new pending block
func (bp *BlockProducer) SwitchActiveBlock() error {
	return bp.blockRepo.SwitchActiveBlock()
}
