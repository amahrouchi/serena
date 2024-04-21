package tools

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// TimeSyncMock is a mock for the TimeSync service.
type TimeSyncMock struct {
	mock.Mock
}

// Current returns the current time.
func (ts *TimeSyncMock) Current() (*time.Time, error) {
	args := ts.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*time.Time), args.Error(1)
}
