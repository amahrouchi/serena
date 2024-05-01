package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/rs/zerolog"
)

// HashGenInterface is an interface for a hash generator.
type HashGenInterface interface {
	FromBlock(block *models.Block) string
}

// HashGen is responsible for generating hashes.
type HashGen struct {
	logger *zerolog.Logger
}

// NewHashGen creates a new HashGen.
func NewHashGen(logger *zerolog.Logger) *HashGen {
	return &HashGen{
		logger: logger,
	}
}

// FromBlock generates a hash from a block.
func (hg *HashGen) FromBlock(block *models.Block) string {
	// return lo.RandomString(64, lo.AlphanumericCharset)

	// Marshal payload
	jsonBlock, err := json.Marshal(block)
	if err != nil {
		hg.logger.Error().
			Err(err).
			Interface("block", block).
			Msg("Failed to marshal block payload")

		panic(err) // TODO: see how to recover/handle this
	}
	hg.logger.Debug().Str("jsonBlock", string(jsonBlock)).Msg("JSON block before hashing")

	// Calculate hash
	hash := sha256.New()
	hash.Write(jsonBlock)

	return hex.EncodeToString(hash.Sum(nil))
}
