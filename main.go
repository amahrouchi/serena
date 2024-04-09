package main

import (
	"github.com/amahrouchi/serena/internal/app"
	"go.uber.org/fx"
)

// TODO:
//   - tests

func main() {
	fx.New(app.Modules).Run()
}
