package http

import "github.com/labstack/echo/v4"

// Route represents a struct that contains the method, path and handler of a route.
type Route interface {
	Method() string
	Path() string
	Handle(c echo.Context) error
}
