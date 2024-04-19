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

// RunTestApp runs the test app
func RunTestApp(t *testing.T, opts ...fx.Option) {
	fxtest.New(
		t,
		app.Options,
		fx.Invoke(func(config *configuration.Config, db *gorm.DB, logger *zerolog.Logger) {
			// Reset the test database
			logger.Info().Msg("Resetting test database")
			db.Exec("DROP SCHEMA public CASCADE")
			db.Exec("CREATE SCHEMA public")
			database.AutoMigrate(db)
		}),
		fx.Options(opts...),
	).RequireStart().RequireStop()
}
