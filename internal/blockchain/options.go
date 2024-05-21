package blockchain

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/blockchain/infrastructure/handlers"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/http"
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		// handlers
		http.AsHandler(handlers.NewWriteHandler),

		// services
		fx.Annotate(services.NewBlockWorker, fx.As(new(services.BlockWorkerInterface))),
		fx.Annotate(services.NewBlockProducer, fx.As(new(services.BlockProducerInterface))),
		fx.Annotate(services.NewPayloadWriter, fx.As(new(services.PayloadWriterInterface))),
		fx.Annotate(repositories.NewBlockRepository, fx.As(new(repositories.BlockRepositoryInterface))),
		fx.Annotate(repositories.NewHashGen, fx.As(new(repositories.HashGenInterface))),
	),
	fx.Invoke(func(worker services.BlockWorkerInterface, config *configuration.Config) {
		// Start the worker
		if config.App.BlockChain.WorkerEnabled {
			err := worker.Start()
			if err != nil {
				panic(err)
			}
		}
	}),
)
