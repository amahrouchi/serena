package main

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/amahrouchi/serena/internal/utils"
	"go.uber.org/fx"
)

// TODO:
//   - install FX and create HTTP handlers through FX
//   - see how to handle env vars and config files (Viper?)
//   - use .env for docker-compose and app config
//   - tests

func main() {
	fx.New(
		// Declare the app modules here
		utils.Modules,

		// Declare and start the HTTP server
		fx.Provide(core.NewEchoServer), // provide the echo server
		fx.Invoke(core.RegisterHooks),  // register the hooks starting/stopping the server
	).Run()
}
