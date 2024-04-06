package main

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// TODO:
//   - install FX and create HTTP handlers through FX
//   - see how to handle env vars and config files (Viper?)
//   - use .env for docker-compose and app config
//   - tests

func main() {
	fx.New(
		fx.Provide(core.NewHTTPServer),
		fx.Invoke(func(e *echo.Echo) {}),
	).Run()
}
