package repositories

import "github.com/rs/zerolog"

// BlockRepositoryInterface is an interface for a block repository
type BlockRepositoryInterface interface {
	CreateEmptyBlock()
}

// BlockRepository is a repository for blocks
type BlockRepository struct {
	logger *zerolog.Logger
}

// NewBlockRepository creates a new BlockRepository
func NewBlockRepository(logger *zerolog.Logger) *BlockRepository {
	return &BlockRepository{
		logger: logger,
	}
}

// CreateEmptyBlock creates an empty block
func (br *BlockRepository) CreateEmptyBlock() {
	br.logger.Debug().Msg("Creating an empty block")
}
