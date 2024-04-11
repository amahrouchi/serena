package blockworker

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/rs/zerolog"
	"time"
)

// BlockWorker is a worker that processes blocks.
type BlockWorker struct {
	logger *zerolog.Logger
	config *core.Config
}

// NewBlockWorker creates a new BlockWorker.
func NewBlockWorker(logger *zerolog.Logger, config *core.Config) *BlockWorker {
	return &BlockWorker{
		logger: logger,
		config: config,
	}
}

// Start starts the block worker.
func (bw *BlockWorker) Start() {
	bw.logger.Info().Msg("Starting block worker...")
	for {
		bw.logger.Info().Msg("Working on a block...")

		// wait between blocks
		// TODO:
		//   - sync with an external time provider
		//   - store the creation time inside the block
		//   - use the creation time of prev block to calculate the wait time
		//   - create blocks using a block producer service
		//   - use a channel provided through FX to pass payload from API to block producer
		time.Sleep(time.Duration(bw.config.BlockDuration) * time.Second)
	}
}
