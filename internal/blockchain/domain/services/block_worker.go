package services

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
)

// BlockWorkerInterface is an interface for a block worker.
type BlockWorkerInterface interface {
	Start()
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
func (bw *BlockWorker) Start() {
	bw.logger.Info().Msg("Starting block worker...")

	// Try to get the reference time 5 times
	refTime, err := bw.timeSync.Current()
	if err != nil {
		bw.logger.Error().Err(err).Msg("Failed to get reference time")
		panic(err)
	}

	// Load the last block
	lastBlock, err := bw.producer.GetLastBlock()
	if lastBlock == nil && err == nil {
		lastBlock = bw.producer.CreateGenesisBlock()
	} else if err != nil {
		bw.logger.Error().Err(err).Msgf("Cannot retrieve last block")
		panic(err)
	}

	for {
		// wait between blocks
		// TODO:
		//   - sync with an external time provider
		//   - store the creation time inside the block
		//   - use the creation time of prev block to calculate the wait time
		//   - create blocks using a block producer service
		//   - use a channel provided through FX to pass payload from API to block producer

		// TODO:
		//   - producer empty block to the chain
		//   - generate an empty previous hash for the first block
		//   - persist the block
		//     - into dummy json file
		//     - into a real DB (a Mongo seems adequate)
		//   - add a channel property to the producer
		//   - provide values in the channel using an API endpoint to save data into the chain
		//   - implement an auth system to write to the chain
		//   - implement a read API to display the chain
		//     - last N blocks
		//     - first N blocks
		//     - from block H1 to H2, from H1, to H2
		//     - care because displaying to much could become costly

		// Get the current time
		currTime, err := bw.timeSync.Current()
		if err != nil {
			bw.logger.Warn().Err(err).Msg("Failed to get current time")
			continue
		}

		// Create the block after the block duration has passed
		diff := currTime.UnixMilli() - refTime.UnixMilli()
		if diff >= int64(bw.config.BlockDuration*1000) {
			bw.logger.Info().Msg("Closing current block: generating hash...")
			bw.logger.Info().Msg("Switching current block: adding prev hash...")
			bw.logger.Info().Msg("Creating a new next block...")
			bw.producer.ProduceBlock()

			// Get the current block time
			blockTime, err := bw.timeSync.Current()
			if err != nil {
				bw.logger.Warn().Err(err).Msg("Failed to get block time")
				continue
			}

			// Update the reference time
			refTime = blockTime
		}
	}
}
