package utils

import (
	"github.com/amahrouchi/serena/internal/utils/infrastructure/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers the routes of the current module.
func RegisterRoutes(
	e *echo.Echo,
	health *handlers.HealthzHandler,
) {
	e.GET("/healthz", health.Handle)
}
