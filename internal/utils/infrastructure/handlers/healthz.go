package handlers

import (
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// HealthzHandler provides a health check endpoint.
type HealthzHandler struct {
	logger *zerolog.Logger
	config *configuration.Config
}

// NewHealthzHandler creates a new instance of HealthzHandler.
func NewHealthzHandler(logger *zerolog.Logger, config *configuration.Config) *HealthzHandler {
	return &HealthzHandler{
		logger: logger,
		config: config,
	}
}

// Route sets the http route configuration.
func (h *HealthzHandler) Route() http.Route {
	return http.Route{
		Method: echo.GET,
		Path:   "/healthz",
	}
}

// Handle handles the health check endpoint.
func (h *HealthzHandler) Handle(c echo.Context) error {
	h.logger.Info().Msg("Health check!")
	h.logger.Info().Msgf("Config from handler: %+v", h.config)

	return c.JSON(200, map[string]string{"status": "ok"})
}
