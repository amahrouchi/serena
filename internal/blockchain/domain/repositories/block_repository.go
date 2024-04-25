package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// BlockRepositoryInterface is an interface for a block repository
type BlockRepositoryInterface interface {
	CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error)
	GetLastBlock() (*models.Block, error)
	CreateGenesisBlock() (*models.Block, error)
}

// BlockRepository is a repository for blocks
type BlockRepository struct {
	timeSync tools.TimeSyncInterface
	db       *gorm.DB
	logger   *zerolog.Logger
}

// NewBlockRepository creates a new BlockRepository
func NewBlockRepository(
	timeSync tools.TimeSyncInterface,
	db *gorm.DB,
	logger *zerolog.Logger,
) *BlockRepository {
	return &BlockRepository{
		timeSync: timeSync,
		db:       db,
		logger:   logger,
	}
}

// CreateEmptyBlock creates an empty block
func (br *BlockRepository) CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error) {
	// Getting current time from NTP
	now, err := br.timeSync.Current()
	if err != nil {
		br.logger.Error().Err(err).Msg("Cannot get current time while creating an empty block")
		return nil, err
	}

	// Block construction
	block := models.Block{
		Status:       status,
		PreviousHash: prevHash,
		Payload:      "{}",
		CreatedAt:    *now,
	}
	result := br.db.Create(&block)

	if result.Error != nil {
		br.logger.Error().Msg("Empty block failed to be created")
		return nil, errors.New("cannot create empty block")
	}

	br.logger.Info().
		Interface("block", block).
		Msg("Empty block created")

	return &block, nil
}

// GetLastBlock gets the last block
func (br *BlockRepository) GetLastBlock() (*models.Block, error) {
	br.logger.Debug().Msg("Getting the last block with hash...")

	// Loading last finalized block
	block := models.Block{}
	result := br.db.Not(&models.Block{Hash: nil}).
		Where(&models.Block{Status: models.BlockStatusActive}).
		Order("created_at").
		Last(&block)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	br.logger.Info().
		Interface("block", block).
		Msg("Last block loaded")

	return &block, nil
}

// CreateGenesisBlock creates the genesis block
func (br *BlockRepository) CreateGenesisBlock() (*models.Block, error) {
	// Getting current time from NTP
	now, err := br.timeSync.Current()
	if err != nil {
		br.logger.Error().Err(err).Msg("Cannot get current time while creating the genesis block")
		return nil, err
	}

	// Hash
	hash := sha256.New()
	hash.Write([]byte(lo.RandomString(40, lo.LettersCharset)))

	// Block construction
	block := models.Block{
		Status:       models.BlockStatusClosed,
		PreviousHash: lo.ToPtr("genesis"),
		Payload:      "{}",
		Hash:         lo.ToPtr(hex.EncodeToString(hash.Sum(nil))),
		CreatedAt:    *now,
	}

	// Create the first blocks into a transaction
	tErr := br.db.Transaction(func(tx *gorm.DB) error {
		// Save genesis block to DB
		result := tx.Create(&block)
		if result.Error != nil {
			br.logger.Error().Msg("Genesis block failed to be created")
			return result.Error
		}

		// Ensure following blocks are created in the same transaction
		db := br.db
		br.db = tx
		defer func() {
			br.db = db
		}()

		// Create the active block
		_, err = br.CreateEmptyBlock(block.Hash, models.BlockStatusActive)
		if err != nil {
			return err
		}

		// Create the next block
		_, err = br.CreateEmptyBlock(nil, models.BlockStatusPending)
		if err != nil {
			return err
		}

		return nil
	})

	// Handle transaction error
	if tErr != nil {
		return nil, tErr
	}

	return &block, nil
}
