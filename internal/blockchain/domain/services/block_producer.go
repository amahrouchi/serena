package services

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/rs/zerolog"
)

// BlockProducerInterface is an interface for a block producer.
type BlockProducerInterface interface {
	CalculateHash(block *models.Block) string
	GetLastBlock() *models.Block
	ProduceBlock()
	CreateGenesisBlock() *models.Block
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

// CalculateHash calculates the hash of the block.
func (bp *BlockProducer) CalculateHash(block *models.Block) string {
	// Marshal headers
	jsonHeader, err := json.Marshal(block.Header)
	if err != nil {
		bp.logger.Error().
			Err(err).
			Interface("block", block).
			Msg("Failed to marshal block header")

		panic(err) // TODO: see how to recover/handle this
	}

	// Marshal payload
	jsonPayload, err := json.Marshal(block.Payload)
	if err != nil {
		bp.logger.Error().
			Err(err).
			Interface("block", block).
			Msg("Failed to marshal block payload")

		panic(err) // TODO: see how to recover/handle this
	}

	// Calculate hash
	preHash := string(jsonHeader) + string(jsonPayload)
	hash := sha256.New()
	hash.Write([]byte(preHash))

	return string(hash.Sum(nil))
}

// GetLastBlock Gets the last created block
func (bp *BlockProducer) GetLastBlock() *models.Block {
	return bp.blockRepo.GetLastBlock()
}

// ProduceBlock produces a block.
func (bp *BlockProducer) ProduceBlock() {
	bp.logger.Debug().Msg("Producing a block...")
	bp.blockRepo.CreateEmptyBlock()
}

// CreateGenesisBlock create the genesis block
func (bp *BlockProducer) CreateGenesisBlock() *models.Block {
	return bp.blockRepo.CreateGenesisBlock()
}