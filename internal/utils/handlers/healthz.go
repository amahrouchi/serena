package handlers

import "github.com/labstack/echo/v4"

// HealthzHandler provides a health check endpoint.
type HealthzHandler struct {
}

// NewHealthzHandler creates a new instance of HealthzHandler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// Handle handles the health check endpoint.
func (h *HealthzHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	}
}
