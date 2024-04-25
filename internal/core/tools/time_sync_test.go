package tools_test

import (
	"github.com/amahrouchi/serena/internal/core/tools"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TimeSyncTestSuite is the test suite for the TimeSync struct.
type TimeSyncTestSuite struct {
	suite.Suite
}

// TestNewTimeSync tests the NewTimeSync method.
func (s *TimeSyncTestSuite) TestNewTimeSync() {
	timeSync := tools.NewTimeSync()
	s.Assert().NotNil(timeSync)
}

// TestCurrent tests the Current method.
func (s *TimeSyncTestSuite) TestCurrent() {
	timeSync := tools.NewTimeSync()

	currTime, err := timeSync.Current()

	s.Assert().NoError(err)
	s.Assert().NotNil(currTime)
}

// TestTimeSyncTestSuite tests the TimeSyncTestSuite.
func TestTimeSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TimeSyncTestSuite))
}
