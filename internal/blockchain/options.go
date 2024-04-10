package blockchain

import (
	"github.com/amahrouchi/serena/internal/blockchain/blockworker"
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		blockworker.NewBlockWorker,
	),
	fx.Invoke(func(worker *blockworker.BlockWorker) {
		go worker.Start()
	}),
)
