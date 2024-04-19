package repositories

import (
	"fmt"
	"github.com/amahrouchi/serena/internal/core/configuration"
	"github.com/amahrouchi/serena/internal/core/tests"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"testing"
)

// BlockRepositorySuite the test suite for the BlockRepository struct
type BlockRepositorySuite struct {
	suite.Suite
}

func (brs *BlockRepositorySuite) TestGetLastBlock() {
	var config *configuration.Config
	tests.RunTestApp(brs.T(), fx.Populate(&config))
	fmt.Printf("config: %+v\n", config)
}

func (brs *BlockRepositorySuite) TestCreateBlock() {
}

// TestBlockRepositorySuite launches the test suite
func TestBlockRepositorySuite(t *testing.T) {
	suite.Run(t, new(BlockRepositorySuite))
}
