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
type TimeSync struct {
	NtpServer string
}

// NewTimeSync creates a new TimeSync service.
func NewTimeSync() *TimeSync {
	return &TimeSync{
		NtpServer: "time.google.com", // TODO: use several sources
	}
}

// Current returns the current time.
func (ts *TimeSync) Current() (*time.Time, error) {
	for i := 0; i < 10; i++ {
		currTime, err := ntp.Time(ts.NtpServer)
		if err == nil {
			return &currTime, nil
		}
	}

	return nil, errors.New("unable to get date from NTP")
}
