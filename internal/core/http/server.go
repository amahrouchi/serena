package http

import (
	"github.com/labstack/echo/v4"
)

// NewEchoServer creates a new Echo server.
func NewEchoServer(handlers []Handler) *echo.Echo {
	// Create an Echo server and add the handlers
	e := echo.New()
	for _, handler := range handlers {
		e.Add(handler.Route().Method, handler.Route().Path, handler.Handle)
	}

	return e
}
