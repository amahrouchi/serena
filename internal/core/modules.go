package core

import (
	"github.com/amahrouchi/serena/internal/utils"
	"go.uber.org/fx"
)

// Modules registers the application dependencies.
var Modules = fx.Options(
	// Declare the app modules here
	utils.Modules,

	// Declare core deps
	fx.Provide(NewConfig),
	fx.Invoke(func(config *Config) {}), // Loads the config
	fx.Provide(NewLogger),

	// Declare and start the HTTP server
	fx.Provide(NewEchoServer), // provide the echo server
	fx.Invoke(RegisterHooks),  // register the hooks starting/stopping the server
)
