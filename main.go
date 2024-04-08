package main

import (
	"github.com/amahrouchi/serena/internal/app"
	"go.uber.org/fx"
)

// TODO:
//   - organize utils and core? in a ddd way
//   - see how to handle env vars and config files (Viper?)
//   - use .env for docker-compose and app config
//   - tests

func main() {
	fx.New(app.Modules).Run()
}
