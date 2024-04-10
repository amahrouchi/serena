package handlers

import (
	"github.com/amahrouchi/serena/internal/core"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

// HealthzHandlerSuite tests the HealthzHandler.
type HealthzHandlerSuite struct {
	suite.Suite

	healthzHandler *HealthzHandler
}

// SetupTest sets up the test suite.
func (s *HealthzHandlerSuite) SetupTest() {
	// Create the logger & config
	logger := zerolog.New(nil).Level(zerolog.Disabled)
	config := core.Config{
		Env:  core.EnvDev,
		Port: 8080,
	}

	// Create the handler
	s.healthzHandler = &HealthzHandler{
		logger: &logger,
		config: &config,
	}
}

// TestNewHealthzHandler tests the NewHealthzHandler method.
func (s *HealthzHandlerSuite) TestNewHealthzHandler() {
	// Create the handler
	handler := NewHealthzHandler(s.healthzHandler.logger, s.healthzHandler.config)

	// Assert handler is created
	s.NotNil(handler)
	s.IsType(&HealthzHandler{}, handler)
	s.Equal(s.healthzHandler.logger, handler.logger)
	s.Equal(s.healthzHandler.config, handler.config)
}

// TestHandle tests the Handle method.
func (s *HealthzHandlerSuite) TestHandle() {
	s.Run("handle successfully", func() {
		// Create the Echo server
		e := echo.New()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(echo.GET, "/healthz", nil)
		context := e.NewContext(request, response)

		// Call the handler
		err := s.healthzHandler.Handle(context)

		// Assert response status and body
		s.NoError(err)
		s.Equal(200, response.Code)
		s.JSONEq(`{"status":"ok"}`, response.Body.String())
	})
}

// TestHealthzHandlerSuite runs the HealthzHandlerSuite.
func TestHealthzHandlerSuite(t *testing.T) {
	suite.Run(t, new(HealthzHandlerSuite))
}