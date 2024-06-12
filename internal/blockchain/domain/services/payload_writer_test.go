package services_test

import (
	"errors"
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/blockchain/domain/services"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// PayloadWriterSuite is the test suite for the PayloadWriter struct.
type PayloadWriterSuite struct {
	suite.Suite
}

// TestWrite tests the Write method.
func (s *PayloadWriterSuite) TestWrite() {
	s.Run("should write data to the blockchain", func() {
		// Mock the repo
		repo := new(repositories.BlockRepositoryMock)
		repo.On(
			"AppendDataToActiveBlock",
			"author",
			map[string]interface{}{"data": "test"},
		).Return(nil)

		// Write data
		writer := services.NewPayloadWriter(
			tests.NewEmptyLogger(),
			repo,
		)
		err := writer.Write("author", map[string]interface{}{"data": "test"})

		// Assert
		s.NoError(err)
	})

	s.Run("should return an error if the repo fails", func() {
		// Mock the repo
		repo := new(repositories.BlockRepositoryMock)
		repo.On(
			"AppendDataToActiveBlock",
			"author",
			map[string]interface{}{"data": "test"},
		).Return(errors.New("repo error"))

		// Write data
		writer := services.NewPayloadWriter(
			tests.NewEmptyLogger(),
			repo,
		)
		err := writer.Write("author", map[string]interface{}{"data": "test"})

		// Assert
		s.Error(err)
		s.Equal("repo error", err.Error())
	})
}

// TestPayloadWriterSuite runs the PayloadWriterSuite.
func TestPayloadWriterSuite(t *testing.T) {
	suite.Run(t, new(PayloadWriterSuite))
}
