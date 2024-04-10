package http

import (
	"github.com/labstack/echo/v4"
)

// NewEchoServer creates a new Echo server.
func NewEchoServer(routes []Route) *echo.Echo {
	// Create an Echo server and add the routes
	e := echo.New()
	for _, route := range routes {
		e.Add(route.Method(), route.Path(), route.Handle)
	}

	return e
}
