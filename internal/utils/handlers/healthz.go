package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// HealthzHandler provides a health check endpoint.
type HealthzHandler struct {
	logger *zerolog.Logger
}

// NewHealthzHandler creates a new instance of HealthzHandler.
func NewHealthzHandler(logger *zerolog.Logger) *HealthzHandler {
	return &HealthzHandler{
		logger: logger,
	}
}

// Handle handles the health check endpoint.
func (h *HealthzHandler) Handle(c echo.Context) error {
	h.logger.Info().Msg("Health check!")

	return c.JSON(200, map[string]string{"status": "ok"})
}
