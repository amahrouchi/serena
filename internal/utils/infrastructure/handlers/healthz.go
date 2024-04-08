package handlers

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// HealthzHandler provides a health check endpoint.
type HealthzHandler struct {
	logger *zerolog.Logger
	config *core.Config
}

// NewHealthzHandler creates a new instance of HealthzHandler.
func NewHealthzHandler(logger *zerolog.Logger, config *core.Config) *HealthzHandler {
	return &HealthzHandler{
		logger: logger,
		config: config,
	}
}

// Handle handles the health check endpoint.
func (h *HealthzHandler) Handle(c echo.Context) error {
	h.logger.Info().Msg("Health check!")
	h.logger.Info().Msgf("Config from handler: %+v", h.config)

	return c.JSON(200, map[string]string{"status": "ok"})
}
