package services

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
)

// BlockWorkerInterface is an interface for a block worker.
type BlockWorkerInterface interface {
	Start() error
}

// BlockWorker is a worker that processes blocks.
type BlockWorker struct {
	producer BlockProducerInterface
	timeSync tools.TimeSyncInterface
	logger   *zerolog.Logger
	config   *configuration.Config
}

// NewBlockWorker creates a new BlockWorker.
func NewBlockWorker(
	producer BlockProducerInterface,
	timeSync tools.TimeSyncInterface,
	logger *zerolog.Logger,
	config *configuration.Config,
) *BlockWorker {
	return &BlockWorker{
		producer: producer,
		timeSync: timeSync,
		logger:   logger,
		config:   config,
	}
}

// Start starts the block worker.
func (bw *BlockWorker) Start() error {
	bw.logger.Info().Msg("Starting block worker...")

	// Try to get the reference time 5 times
	refTime, err := bw.timeSync.Current()
	if err != nil {
		bw.logger.Error().Err(err).Msg("Failed to get reference time while starting block worker")
		return err
	}

	// Load the last block
	lastBlock, err := bw.producer.GetActiveBlock()
	if lastBlock == nil && err == nil {
		// Create the genesis block
		_, err := bw.producer.CreateGenesisBlock()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	for {
		// Get the current time
		currTime, err := bw.timeSync.Current()
		if err != nil {
			bw.logger.Warn().Err(err).Msg("Failed to get current time")
			continue
		}

		// Create the block after the block duration has passed
		diff := currTime.UnixMilli() - refTime.UnixMilli()
		if diff >= int64(bw.config.App.BlockChain.Interval*1000) {
			// Close the current block and create a new one
			err = bw.producer.SwitchActiveBlock()
			if err != nil {
				bw.logger.Error().Err(err).Msg("Failed to switch active block")
				panic(err)
			}

			// Get the current block time
			blockTime, err := bw.timeSync.Current()
			if err != nil {
				bw.logger.Warn().Err(err).Msg("Failed to get block time")
				panic(err)
			}

			// Update the reference time
			refTime = blockTime
		}
	}
}
