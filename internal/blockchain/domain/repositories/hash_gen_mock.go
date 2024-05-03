package repositories

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/models"
	"github.com/stretchr/testify/mock"
)

// HashGenMock is a mock for the HashGen struct.
type HashGenMock struct {
	mock.Mock
}

// FromBlock is a mock for the FromBlock method.
func (hgm *HashGenMock) FromBlock(block *models.Block) (string, error) {
	args := hgm.Called(block)
	return args.String(0), args.Error(1)
}
