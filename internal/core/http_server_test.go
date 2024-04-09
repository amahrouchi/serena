package core

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
	server := NewEchoServer()

	s.NotNil(server)
	s.IsType(&echo.Echo{}, server)
}

// TestHttpServerSuite tests the HttpServerSuite.
func TestHttpServerSuite(t *testing.T) {
	suite.Run(t, new(HttpServerSuite))
}
