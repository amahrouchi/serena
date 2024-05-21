package services

import "github.com/rs/zerolog"

// PayloadWriterInterface defines the interface for the payload writer service.
type PayloadWriterInterface interface {
	Write(author string, data map[string]any) error
}

// PayloadWriter provides a service to write data to the blockchain.
type PayloadWriter struct {
	logger *zerolog.Logger
}

// NewPayloadWriter creates a new instance of PayloadWriter.
func NewPayloadWriter(logger *zerolog.Logger) *PayloadWriter {
	return &PayloadWriter{
		logger: logger,
	}
}

// Write writes data to the blockchain.
func (w *PayloadWriter) Write(author string, data map[string]any) error {
	w.logger.Info().
		Str("author", author).
		Interface("data", data).
		Msgf("Writing data to the blockchain")

	// TODO: Implement the write logic here

	return nil
}
