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
	CreateGenesisBlock() (*models.Block, error)
	CreateEmptyBlock(prevHash *string, status models.BlockStatus) (*models.Block, error)
	SwitchActiveBlock() error
	GetActiveBlock() (*models.Block, error)
	GetPendingBlock() (*models.Block, error)
	Update(block *models.Block) error
	AppendDataToActiveBlock(author string, data map[string]any) error
	Activate(block *models.Block) error
	Close(block *models.Block) error
}

// BlockRepository is a repository for blocks
type BlockRepository struct {
	timeSync tools.TimeSyncInterface
	hashGen  HashGenInterface
	db       *gorm.DB
	logger   *zerolog.Logger
}

// NewBlockRepository creates a new BlockRepository
func NewBlockRepository(
	timeSync tools.TimeSyncInterface,
	hashGen HashGenInterface,
	db *gorm.DB,
	logger *zerolog.Logger,
) *BlockRepository {
	return &BlockRepository{
		timeSync: timeSync,
		hashGen:  hashGen,
		db:       db,
		logger:   logger,
	}
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
		Payload:      "[]",
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
		Payload:      "[]",
		CreatedAt:    *now,
	}
	result := br.db.Create(&block)

	if result.Error != nil {
		br.logger.Error().Msg("Empty block failed to be created")
		return nil, errors.New("cannot create empty block")
	}

	return &block, nil
}

// SwitchActiveBlock closes the active block, activates the pending one and creates a new pending block
func (br *BlockRepository) SwitchActiveBlock() error {
	return br.db.Transaction(func(tx *gorm.DB) error {
		// Ensure following queries are run in the transaction
		db := br.db
		br.db = tx
		defer func() {
			br.db = db
		}()

		// Load the active block
		activeBlock, err := br.GetActiveBlock()
		if err != nil {
			return err
		}

		// If no active block, return an error
		if activeBlock == nil {
			return errors.New("no active block found")
		}

		br.logger.Debug().Interface("block", activeBlock).Msg("Active block loaded")

		// Close the active block
		err = br.Close(activeBlock)
		if err != nil {
			return err
		}
		br.logger.Debug().Interface("block", activeBlock).Msg("Active block closed")

		// Calculate hash of the active block
		hash, err := br.hashGen.FromBlock(activeBlock)
		if err != nil {
			return err
		}
		activeBlock.Hash = lo.ToPtr(hash)
		err = br.Update(activeBlock)
		if err != nil {
			return err
		}
		br.logger.Debug().Interface("block", activeBlock).Msg("Active block hash calculated")

		// Load the pending block
		newActiveBlock, err := br.GetPendingBlock()
		if err != nil {
			return err
		}
		br.logger.Debug().Interface("block", newActiveBlock).Msg("Pending block loaded")

		// Activate the pending block
		if newActiveBlock != nil {
			newActiveBlock.PreviousHash = activeBlock.Hash
			err = br.Activate(newActiveBlock)
			if err != nil {
				return err
			}
		} else {
			newActiveBlock, err = br.CreateEmptyBlock(activeBlock.Hash, models.BlockStatusActive)
			if err != nil {
				return err
			}
		}
		br.logger.Debug().Interface("block", newActiveBlock).Msg("Pending block activated")

		// Create new pending block
		pendingBlock, err := br.CreateEmptyBlock(nil, models.BlockStatusPending)
		if err != nil {
			return err
		}
		br.logger.Debug().Interface("block", pendingBlock).Msg("New pending block created")

		return nil
	})
}

// GetActiveBlock gets the last block
func (br *BlockRepository) GetActiveBlock() (*models.Block, error) {
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

	return &block, nil
}

// GetPendingBlock gets the pending block
func (br *BlockRepository) GetPendingBlock() (*models.Block, error) {
	// Loading last pending block
	block := models.Block{}
	result := br.db.Where(&models.Block{Status: models.BlockStatusPending}).
		Order("created_at").
		First(&block)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &block, nil
}

// Update updates a block
func (br *BlockRepository) Update(block *models.Block) error {
	result := br.db.Save(block)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// AppendDataToActiveBlock appends data to the active block
func (br *BlockRepository) AppendDataToActiveBlock(author string, data map[string]any) error {
	// Load the active block
	activeBlock, err := br.GetActiveBlock()
	if err != nil {
		return err
	}

	// If no active block, return an error
	if activeBlock == nil {
		return errors.New("no active block found to append data to")
	}

	// Append data to the active block
	var payload []models.BlockPayloadItem
	err = json.Unmarshal([]byte(activeBlock.Payload), &payload)
	if err != nil {
		return err
	}

	// Getting current time
	createdAt, err := br.timeSync.Current()
	if err != nil {
		return err
	}

	// Append data to the payload
	payload = append(payload, models.BlockPayloadItem{
		Author:    author,
		Data:      data,
		CreatedAt: *createdAt,
	})

	// Update the active block
	payloadToSave, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	activeBlock.Payload = string(payloadToSave)

	return br.Update(activeBlock)
}

// Activate activates a block
func (br *BlockRepository) Activate(block *models.Block) error {
	block.Status = models.BlockStatusActive

	return br.Update(block)
}

// Close closes a block
func (br *BlockRepository) Close(block *models.Block) error {
	block.Status = models.BlockStatusClosed

	return br.Update(block)
}
