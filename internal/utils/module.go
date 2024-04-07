package utils

import (
	"github.com/amahrouchi/serena/internal/utils/handlers"
	"go.uber.org/fx"
)

// Modules is used to register the current module dependencies.
var Modules = fx.Options(
	fx.Invoke(RegisterRoutes),
	fx.Provide(handlers.NewHealthzHandler),
)
