package http

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HttpServerSuite tests the HttpServer.
type HttpServerSuite struct {
	suite.Suite
}

// TestNewEchoServer tests the NewEchoServer method.
func (s *HttpServerSuite) TestNewEchoServer() {
	// Prepare the mocked handler
	handler := new(mockHandler)
	handler.On("Route").Return(Route{Method: echo.GET, Path: "/mock"})
	handler.On("Handle").Return(nil)

	// Create the server
	server := NewEchoServer([]Handler{handler})

	// Assert the server
	s.NotNil(server)
	s.IsType(&echo.Echo{}, server)
}

// TestHttpServerSuite tests the HttpServerSuite.
func TestHttpServerSuite(t *testing.T) {
	suite.Run(t, new(HttpServerSuite))
}
