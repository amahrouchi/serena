package services

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/rs/zerolog"
)

// PayloadWriterInterface defines the interface for the payload writer service.
type PayloadWriterInterface interface {
	Write(author string, data map[string]any) error
}

// PayloadWriter provides a service to write data to the blockchain.
type PayloadWriter struct {
	blockRepo repositories.BlockRepositoryInterface
	logger    *zerolog.Logger
}

// NewPayloadWriter creates a new instance of PayloadWriter.
func NewPayloadWriter(logger *zerolog.Logger, blockRepo repositories.BlockRepositoryInterface) *PayloadWriter {
	return &PayloadWriter{
		blockRepo: blockRepo,
		logger:    logger,
	}
}

// Write writes data to the blockchain.
func (w *PayloadWriter) Write(author string, data map[string]any) error {
	w.logger.Info().
		Str("author", author).
		Interface("data", data).
		Msgf("Writing data to the blockchain")

	err := w.blockRepo.AppendDataToActiveBlock(author, data)
	if err != nil {
		return err
	}

	return nil
}
