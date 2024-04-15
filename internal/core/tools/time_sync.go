package tools

import (
	"errors"
	"github.com/beevik/ntp"
	"time"
)

// TimeSyncInterface is the interface for the TimeSync service.
type TimeSyncInterface interface {
	Current() (*time.Time, error)
}

// TimeSync is a service for syncing time.
type TimeSync struct{}

// NewTimeSync creates a new TimeSync service.
func NewTimeSync() *TimeSync {
	return &TimeSync{}
}

// Current returns the current time.
func (ts *TimeSync) Current() (*time.Time, error) {
	for i := 0; i < 5; i++ {
		currTime, err := ntp.Time("time.google.com") // TODO: use several sources
		if err == nil {
			return &currTime, nil
		}
	}

	return nil, errors.New("unable to get date from NTP")
}
