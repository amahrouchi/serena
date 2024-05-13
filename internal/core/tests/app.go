package tests

import (
	"github.com/amahrouchi/serena/internal/app"
	"github.com/amahrouchi/serena/internal/core/database"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

// TestApp is a test application.
type TestApp struct {
	EnableFxLogs bool
}

// NewTestApp creates a new test application.
func NewTestApp(enableFxLogs bool) *TestApp {
	return &TestApp{
		EnableFxLogs: enableFxLogs,
	}
}

// Run runs the test application.
func (ta *TestApp) Run(t *testing.T, opts ...fx.Option) *fxtest.App {
	if !ta.EnableFxLogs {
		opts = append(opts, fx.NopLogger)
	}

	return fxtest.New(
		t,
		app.Options,
		fx.Options(opts...),
		fx.Invoke(func(migrator *database.Migrator, logger *zerolog.Logger) {
			// Reset the test database (Postgres specific)
			logger.Info().Msg("Resetting the test database...")
			migrator.ResetDatabase()
			logger.Info().Msg("Test database successfully reset")
		}),
	).RequireStart()
}
