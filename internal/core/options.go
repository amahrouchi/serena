package core

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"go.uber.org/fx"
)

// Options registers the core package FX options.
var Options = fx.Options(
	// Declare core deps
	fx.Provide(
		newConfig,
		newLogger,
		fx.Annotate(
			http.NewEchoServer,
			fx.ParamTags(`group:"handlers"`),
		),
	),
	fx.Invoke(
		loadConfig,
		registerHooks,
		dbConnection,
	),
)
