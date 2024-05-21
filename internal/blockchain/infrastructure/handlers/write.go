package handlers

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/blockchain/infrastructure/requests"
	"github.com/amahrouchi/serena/internal/core/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	nethttp "net/http"
)

// WriteHandler provides a write endpoint to write data to the blockchain.
type WriteHandler struct {
	PayloadWriter services.PayloadWriterInterface
	Logger        *zerolog.Logger
}

// NewWriteHandler creates a new instance of WriteHandler.
func NewWriteHandler(
	payloadWriter services.PayloadWriterInterface,
	logger *zerolog.Logger,
) *WriteHandler {
	return &WriteHandler{
		PayloadWriter: payloadWriter,
		Logger:        logger,
	}
}

// Route sets the http route configuration.
func (h *WriteHandler) Route() http.Route {
	return http.Route{
		Method: echo.POST,
		Path:   "/write",
	}
}

// Handle handles the write endpoint.
func (h *WriteHandler) Handle(c echo.Context) error {
	// Bind request payload
	data := new(requests.WriteRequest)
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, err.Error())
	}

	// Validate payload
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		h.Logger.Error().Err(err).Msg("Write process validation error")

		return c.JSON(nethttp.StatusBadRequest, map[string]any{"message": "Bad request"})
	}

	// Write data to the blockchain
	err = h.PayloadWriter.Write(data.Author, data.Data)
	if err != nil {
		h.Logger.Error().Err(err).Msg("Unable to write data to the blockchain")

		return c.JSON(
			nethttp.StatusInternalServerError,
			map[string]any{"message": "Internal server error"},
		)
	}

	return c.JSON(nethttp.StatusOK, data)
}
