package http_test

import (
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HandlerTestSuite tests the Handler struct.
type HandlerTestSuite struct {
	suite.Suite
}

// TestHandler is a test handler.
type TestHandler struct {
}

// Route returns the route of the handler.
func (th *TestHandler) Route() http.Route {
	return http.Route{
		Method: "GET",
		Path:   "/test",
	}
}

// Handle handles the test request.
func (th *TestHandler) Handle(_ echo.Context) error {
	return nil
}

// TestAsHandler tests the AsHandler function.
func (s *HandlerTestSuite) TestAsHandler() {
	handler := TestHandler{}
	annotation := http.AsHandler(handler)
	s.NotNil(annotation)
}

// TestHandlerTestSuite tests the HandlerTestSuite.
func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
