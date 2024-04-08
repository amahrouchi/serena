package app

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/amahrouchi/serena/internal/utils"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	core.Modules,
	utils.Modules,
)
