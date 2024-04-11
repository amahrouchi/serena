package domain

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/rs/zerolog"
)

// BlockProducer is responsible for producing blocks.
type BlockProducer struct {
	logger zerolog.Logger
}

// NewBlockProducer creates a new BlockProducer.
func NewBlockProducer(logger zerolog.Logger) *BlockProducer {
	return &BlockProducer{
		logger: logger,
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

		panic(err) // TODO: see how to recover from this
	}

	// Marshal payload
	jsonPayload, err := json.Marshal(block.Payload)
	if err != nil {
		bp.logger.Error().
			Err(err).
			Interface("block", block).
			Msg("Failed to marshal block payload")

		panic(err) // TODO: see how to recover from this
	}

	// Calculate hash
	preHash := string(jsonHeader) + string(jsonPayload)
	hash := sha256.New()
	hash.Write([]byte(preHash))

	return string(hash.Sum(nil))
}
