package utils

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/amahrouchi/serena/internal/utils/infrastructure/handlers"
	"go.uber.org/fx"
)

// Options registers the utils package FX options.
var Options = fx.Options(
	fx.Provide(
		fx.Annotate(
			handlers.NewHealthzHandler,
			fx.As(new(http.Handler)),
			fx.ResultTags(`group:"handlers"`),
		),
	),
)
