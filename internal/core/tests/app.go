package tests

import (
	"github.com/amahrouchi/serena/internal/app"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/database"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
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
func (ta *TestApp) Run(t *testing.T, opts ...fx.Option) {
	if !ta.EnableFxLogs {
		opts = append(opts, fx.NopLogger)
	}

	fxtest.New(
		t,
		app.Options,
		fx.Invoke(func(config *configuration.Config, db *gorm.DB, logger *zerolog.Logger) {
			// Reset the test database
			logger.Info().Msg("Resetting test database")
			db.Exec("DROP SCHEMA IF EXISTS public CASCADE")
			db.Exec("CREATE SCHEMA public")
			database.AutoMigrate(db)
		}),
		fx.Options(opts...),
	).RequireStart().RequireStop()
}
