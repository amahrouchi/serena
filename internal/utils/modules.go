package utils

import (
	"github.com/amahrouchi/serena/internal/utils/infrastructure/handlers"
	"go.uber.org/fx"
)

// Modules registers the current module dependencies.
var Modules = fx.Options(
	fx.Invoke(RegisterRoutes),
	fx.Provide(handlers.NewHealthzHandler),
)
