package handlers_test

import (
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/blockchain/infrastructure/handlers"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// WriteHandlerTestSuite is the test suite for the WriteHandler struct.
type WriteHandlerTestSuite struct {
	suite.Suite
}

// TestRoute tests the Route method.
func (s *WriteHandlerTestSuite) TestRoute() {
	// Build service
	serviceAndMocks := buildService()

	route := serviceAndMocks.writeHandler.Route()

	s.Equal(echo.POST, route.Method)
	s.Equal("/write", route.Path)
}

// TestHandle tests the Handle method.
func (s *WriteHandlerTestSuite) TestHandle() {
	s.Run("handle successfully", func() {
		// Build service
		serviceAndMocks := buildService()

		// Mock the payload writer
		serviceAndMocks.payloadWriter.
			On("Write", "author", map[string]interface{}{"data": "test"}).
			Return(nil)

		// Create the Echo server
		e := echo.New()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			echo.POST,
			"/write",
			strings.NewReader(`{"author": "author", "data": {"data": "test"}}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		context := e.NewContext(request, response)

		// Call the handler
		err := serviceAndMocks.writeHandler.Handle(context)

		// Assert
		s.NoError(err)
		s.Equal(http.StatusOK, response.Code)
	})

	s.Run("handle with bad request", func() {
		// Build service
		serviceAndMocks := buildService()

		// Create the Echo server
		e := echo.New()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(echo.POST, "/write", nil)
		context := e.NewContext(request, response)

		// Call the handler
		err := serviceAndMocks.writeHandler.Handle(context)

		// Assert
		s.NoError(err)
		s.Equal(http.StatusBadRequest, response.Code)
	})

	s.Run("handle with write failure", func() {
		// Build service
		serviceAndMocks := buildService()

		// Mock the payload writer
		serviceAndMocks.payloadWriter.
			On("Write", "author", map[string]interface{}{"data": "test"}).
			Return(errors.New("write error"))

		// Create the Echo server
		e := echo.New()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			echo.POST,
			"/write",
			strings.NewReader(`{"author": "author", "data": {"data": "test"}}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		context := e.NewContext(request, response)

		// Call the handler
		err := serviceAndMocks.writeHandler.Handle(context)

		// Assert
		s.NoError(err)
		s.Equal(http.StatusInternalServerError, response.Code)
	})
}

// TestWriteHandlerTestSuite runs the WriteHandlerTestSuite test suite.
func TestWriteTestSuite(t *testing.T) {
	suite.Run(t, new(WriteHandlerTestSuite))
}

// writerAndMockHolder holds the write handler and its mock.
type writerAndMockHolder struct {
	writeHandler  *handlers.WriteHandler
	payloadWriter *services.PayloadWriterMock
}

// buildService creates a new instance of the write handler and its mock.
func buildService() *writerAndMockHolder {
	payloadWriter := &services.PayloadWriterMock{}
	logger := tests.NewEmptyLogger()
	writeHandler := handlers.NewWriteHandler(payloadWriter, logger)

	return &writerAndMockHolder{
		writeHandler:  writeHandler,
		payloadWriter: payloadWriter,
	}
}
