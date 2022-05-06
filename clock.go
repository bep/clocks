package clock

import (
	"time"
)

// Clock provides the sub set of methods in time.Time that this package provides.
type Clock interface {
	Now() time.Time
}

type clock struct {
	offset time.Duration
}

// Now returns the current time relative to the configured start time.
func (c *clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// Start creates a new Clock starting at t.
func Start(t time.Time) Clock {
	return &clock{
		offset: t.Sub(time.Now()),
	}
}
