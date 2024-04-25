package tools

import (
	"github.com/samber/lo"
	"time"
)

// TimeSyncInterface is the interface for the TimeSync service.
type TimeSyncInterface interface {
	Current() (*time.Time, error)
}

// TimeSync is a service for syncing time.
type TimeSync struct {
}

// NewTimeSync creates a new TimeSync service.
func NewTimeSync() *TimeSync {
	return &TimeSync{}
}

// Current returns the current time.
func (ts *TimeSync) Current() (*time.Time, error) {
	/*
	 * I wanted to work with NTP servers
	 * but it creates to much network delay.
	 * Maybe later we can declare a node responsible
	 * for time sync and query it here?
	 */

	return lo.ToPtr(time.Now()), nil
}
