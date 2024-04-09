package main

import (
	"github.com/amahrouchi/serena/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Options).Run()
}
