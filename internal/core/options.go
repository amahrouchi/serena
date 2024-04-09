package core

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"go.uber.org/fx"
)

// Options registers the core package FX options.
var Options = fx.Options(
	// Declare core deps
	fx.Provide(NewConfig),
	fx.Invoke(LoadConfig), // Loads the config
	fx.Provide(NewLogger),

	// Declare and start the HTTP server
	fx.Provide(http.NewEchoServer), // provide the echo server
	fx.Invoke(RegisterHooks),       // register the hooks starting/stopping the server
)
