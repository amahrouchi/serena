package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// Handler represents a struct that contains the method, path and handler of a route.
type Handler interface {
	Route() Route
	Handle(c echo.Context) error
}

// Route represents an HTTP route configuration.
type Route struct {
	Method string
	Path   string
}

// AsHandler is an FX constructor that annotates a struct as a Handler.
func AsHandler(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}
