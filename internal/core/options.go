package core

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/amahrouchi/serena/internal/core/tools"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Options registers the core package FX options.
var Options = fx.Options(
	// Declare core deps
	fx.Provide(
		newConfig,
		newLogger,
		newDbConnection,
		fx.Annotate(http.NewEchoServer, fx.ParamTags(`group:"handlers"`)),
		fx.Annotate(tools.NewTimeSync, fx.As(new(tools.TimeSyncInterface))),
	),
	fx.Invoke(
		loadConfig,
		registerHooks,
		func(db *gorm.DB) {}, // force the DB connection creation
	),
)
