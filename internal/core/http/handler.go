package http

import "github.com/labstack/echo/v4"

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
