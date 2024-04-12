package services

import (
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
	currTime, err := ntp.Time("time.google.com")
	if err != nil {
		return nil, err
	}

	return &currTime, nil
}
