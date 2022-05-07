package clock

import (
	"time"
)

// Clock provides the sub set of methods in time.Time that this package provides.
type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration
}

type clock struct {
	offset time.Duration
}

// Now returns the current time relative to the configured start time.
func (c *clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// Since returns the time elapsed since t.
func (c *clock) Since(t time.Time) time.Duration {
	return c.Now().Sub(t)

}

// Until returns the duration until t.
func (c *clock) Until(t time.Time) time.Duration {
	return t.Sub(c.Now())
}

// Start creates a new Clock starting at t.
func Start(t time.Time) Clock {
	return &clock{
		offset: t.Sub(time.Now()),
	}
}
