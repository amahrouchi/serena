package core

import (
	"github.com/amahrouchi/serena/internal/utils"
	"go.uber.org/fx"
)

// Modules is used to register the application dependencies.
var Modules = fx.Options(
	// Declare the app modules here
	utils.Modules,

	// Declare core deps
	fx.Provide(NewLogger),

	// Declare and start the HTTP server
	fx.Provide(NewEchoServer), // provide the echo server
	fx.Invoke(RegisterHooks),  // register the hooks starting/stopping the server
)
