package blockchain

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"go.uber.org/fx"
)

var Options = fx.Options(
	fx.Provide(
		fx.Annotate(services.NewBlockWorker, fx.As(new(services.BlockWorkerInterface))),
		fx.Annotate(services.NewBlockProducer, fx.As(new(services.BlockProducerInterface))),
		fx.Annotate(repositories.NewBlockRepository, fx.As(new(repositories.BlockRepositoryInterface))),
	),
	fx.Invoke(func(worker services.BlockWorkerInterface, config *configuration.Config) {
		if config.BlockWorkerEnabled {
			go func() {
				err := worker.Start()
				if err != nil {
					panic(err)
				}
			}()
		}
	}),
)
