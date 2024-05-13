package services_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tests"
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
func (s *BlockWorkerSuite) TestAppStartWithWorker() {
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

// TestBlockWorkerSuite runs the block worker suite.
func TestBlockWorkerSuite(t *testing.T) {
	suite.Run(t, new(BlockWorkerSuite))
}
