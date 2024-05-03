package repositories_test

import (
	"github.com/amahrouchi/serena/internal/blockchain/domain/repositories"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

// HashGenSuite is the test suite for the HashGen struct.
type HashGenSuite struct {
	suite.Suite
}

// TestNewHashGen tests the NewHashGen method.
func (hgs *HashGenSuite) TestNewHashGen() {
	hashGen := repositories.NewHashGen(
		tests.NewEmptyLogger(),
	)

	hgs.NotNil(hashGen)
}

// TestHashGenSuite tests the HashGenSuite.
func TestHashGenSuite(t *testing.T) {
	suite.Run(t, new(HashGenSuite))
}
