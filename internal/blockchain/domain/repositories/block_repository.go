package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/amahrouchi/serena/internal/core"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// BlockRepositoryInterface is an interface for a block repository
type BlockRepositoryInterface interface {
	CreateEmptyBlock()
	GetLastBlock() *models.Block
	CreateGenesisBlock() *models.Block
}

// BlockRepository is a repository for blocks
type BlockRepository struct {
	timeSync core.TimeSyncInterface
	db       *gorm.DB
	logger   *zerolog.Logger
}

// NewBlockRepository creates a new BlockRepository
func NewBlockRepository(
	timeSync core.TimeSyncInterface,
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
func (br *BlockRepository) CreateEmptyBlock() {
	br.logger.Debug().Msg("Creating an empty block")
	// TODO: implement
}

// GetLastBlock gets the last block
func (br *BlockRepository) GetLastBlock() *models.Block {
	br.logger.Debug().Msg("Getting the last block with hash...")
	// TODO: implement

	return nil
}

// CreateGenesisBlock creates the genesis block
func (br *BlockRepository) CreateGenesisBlock() *models.Block {
	// Getting current time from NTP
	now, err := br.timeSync.Current()
	if err != nil {
		br.logger.Error().Err(err).Msg("Cannot get current time while creating the genesis block")
		panic(err)
	}

	br.logger.Debug().Msg("Creating genesis block")

	// Hash
	hash := sha256.New()
	hash.Write([]byte(lo.RandomString(40, lo.LettersCharset)))

	// Payload
	payload, err := json.Marshal(make(map[string]any))
	if err != nil {
		panic(err)
	}

	// Block construction
	block := models.Block{
		Header: &models.BlockHeader{
			PreviousHash: "",
			CreationDate: uint64(now.UnixMilli()),
		},
		Hash:    hex.EncodeToString(hash.Sum(nil)),
		Payload: string(payload),
	}

	// Save block to DB
	br.db.Create(&block)
	if block.ID == 0 {
		panic(errors.New("cannot create genesis block"))
	}

	return &block
}
