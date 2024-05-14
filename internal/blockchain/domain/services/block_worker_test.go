package services_test

import (
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"testing"
	"time"
)

// BlockWorkerSuite tests the block worker service.
type BlockWorkerSuite struct {
	suite.Suite
}

// TestAppStartWithWorker tests starting the app with the worker enabled.
func (s *BlockWorkerSuite) TestStart() {
	// Prepare the test app
	var worker services.BlockWorkerInterface
	app := tests.NewTestApp(false).Run(
		s.T(),
		fx.Decorate(func(config *configuration.Config) *configuration.Config {
			config.App.BlockChain.Interval = 1
			return config
		}),
		fx.Populate(&worker),
	)
	defer app.RequireStop()

	// Start the worker
	err := worker.Start()

	// Wait for the worker to process a few blocks
	time.Sleep(3 * time.Second)
	worker.(*services.BlockWorker).QuitChan <- true

	// Assert no error occurred
	s.NoError(err)
}

// TestStartFailures tests starting the worker with failures.
func (s *BlockWorkerSuite) TestStartFailures() {
	s.Run("Failed to get reference time", func() {
		// Prepare the test app
		var worker services.BlockWorkerInterface
		app := tests.NewTestApp(false).Run(
			s.T(),
			fx.Decorate(func() tools.TimeSyncInterface {
				mockTimeSync := new(tools.TimeSyncMock)
				mockTimeSync.On("Current").Return(nil, errors.New("error"))
				return mockTimeSync
			}),
			fx.Populate(&worker),
		)
		defer app.RequireStop()

		// Start the worker
		err := worker.Start()

		// Assert no error occurred
		s.Error(err)
		s.Equal("error", err.Error())
	})

	s.Run("Failed to get active block", func() {
		// Prepare the test app
		var worker services.BlockWorkerInterface
		app := tests.NewTestApp(false).Run(
			s.T(),
			fx.Decorate(func() services.BlockProducerInterface {
				mockBlockProducer := new(services.BlockProducerMock)
				mockBlockProducer.On("GetActiveBlock").Return(nil, errors.New("error"))
				return mockBlockProducer
			}),
			fx.Populate(&worker),
		)
		defer app.RequireStop()

		// Start the worker
		err := worker.Start()

		// Assert no error occurred
		s.Error(err)
		s.Equal("error", err.Error())
	})

	s.Run("Failed to create genesis block", func() {
		// Prepare the test app
		var worker services.BlockWorkerInterface
		app := tests.NewTestApp(false).Run(
			s.T(),
			fx.Decorate(func() services.BlockProducerInterface {
				mockBlockProducer := new(services.BlockProducerMock)
				mockBlockProducer.On("GetActiveBlock").Return(nil, nil)
				mockBlockProducer.On("CreateGenesisBlock").Return(nil, errors.New("error"))
				return mockBlockProducer
			}),
			fx.Populate(&worker),
		)
		defer app.RequireStop()

		// Start the worker
		err := worker.Start()

		// Assert no error occurred
		s.Error(err)
		s.Equal("error", err.Error())
	})
}

// TestBlockWorkerSuite runs the block worker suite.
func TestBlockWorkerSuite(t *testing.T) {
	suite.Run(t, new(BlockWorkerSuite))
}
