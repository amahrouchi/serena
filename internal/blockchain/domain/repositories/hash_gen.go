package repositories

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/samber/lo"
)

// HashGenInterface is an interface for a hash generator.
type HashGenInterface interface {
	FromBlock(block *models.Block) string
}

// HashGen is responsible for generating hashes.
type HashGen struct {
}

// NewHashGen creates a new HashGen.
func NewHashGen() *HashGen {
	return &HashGen{}
}

// FromBlock generates a hash from a block.
func (hg *HashGen) FromBlock(block *models.Block) string {
	return lo.RandomString(64, lo.AlphanumericCharset)
}
