package main

import (
	"github.com/amahrouchi/serena/internal/core"
	"go.uber.org/fx"
)

// TODO:
//   - install FX and create HTTP handlers through FX
//   - organize utils and core? in a ddd way
//   - see what is spf13/cobra (some king of command line tool?)
//   - see how to handle env vars and config files (Viper?)
//   - use .env for docker-compose and app config
//   - tests

func main() {
	fx.New(core.Modules).Run()
}
