package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// BlockRepositoryInterface is an interface for a block repository
type BlockRepositoryInterface interface {
	CreateEmptyBlock() error
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
func (br *BlockRepository) CreateEmptyBlock() error {
	br.logger.Debug().Msg("Creating an empty block")

	// TODO: implement
	return errors.New("not implemented")
}

// GetLastBlock gets the last block
func (br *BlockRepository) GetLastBlock() (*models.Block, error) {
	br.logger.Debug().Msg("Getting the last block with hash...")

	// Loading last finalized block
	block := models.Block{}
	result := br.db.Not(&models.Block{Hash: nil}).
		Order("created_at").
		Last(&block)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	br.logger.Info().
		Uint("lastBlockId", block.ID).
		Str("lastBlockHash", *block.Hash).
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

	// Payload
	payload, err := json.Marshal(make(map[string]any))
	if err != nil {
		br.logger.Error().Err(err).Msg("Cannot marshal payload while creating the genesis block")
		return nil, err
	}

	// Block construction
	block := models.Block{
		PreviousHash: "",
		CreatedAt:    *now,
		Hash:         lo.ToPtr(hex.EncodeToString(hash.Sum(nil))),
		Payload:      string(payload),
	}

	// Save block to DB
	result := br.db.Create(&block)
	if result.Error != nil {
		br.logger.Error().Msg("Genesis block failed to be created")
		return nil, errors.New("cannot create genesis block")
	}

	return &block, nil
}
