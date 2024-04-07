package core

import (
	"github.com/labstack/echo/v4"
)

// NewEchoServer creates a new Echo server.
func NewEchoServer() *echo.Echo {
	return echo.New()
}
