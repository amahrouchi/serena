package utils

import (
	"github.com/amahrouchi/serena/internal/utils/infrastructure/handlers"
	"go.uber.org/fx"
)

// Options registers the utils package FX options.
var Options = fx.Options(
	fx.Invoke(RegisterRoutes),
	fx.Provide(handlers.NewHealthzHandler),
)
