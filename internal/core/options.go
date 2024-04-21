package core

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/database"
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Options registers the core package FX options.
var Options = fx.Options(
	// Declare core deps
	fx.Provide(
		configuration.NewConfig,
		configuration.NewConfigYaml,
		tools.NewLogger,
		database.NewPostgresDbConnection,
		database.NewMigrator,
		fx.Annotate(http.NewEchoServer, fx.ParamTags(`group:"handlers"`)),
		fx.Annotate(tools.NewTimeSync, fx.As(new(tools.TimeSyncInterface))),
	),
	fx.Invoke(
		configuration.LoadConfig,
		configuration.RegisterHooks,
		// Auto-migrate the schema
		func(db *gorm.DB, config *configuration.Config, migrator *database.Migrator, logger *zerolog.Logger) {
			if config.Env != configuration.EnvTest {
				// TODO: handle migrations properly (test env can keep this behaviour, see RunTestApp func)
				logger.Info().Msg("Auto-migration of the schema...")
				migrator.AutoMigrate()
			}
		},
	),
)
