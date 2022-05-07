package clock

import (
	"time"
)

// Clock provides the sub set of methods in time.Time that this package provides.
type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration

	// Offset returns the offset of this clock relative to the system clock.
	Offset() time.Duration
}

// Start creates a new Clock starting at t.
func Start(t time.Time) Clock {
	return &clock{
		offset: t.Sub(time.Now()),
	}
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

// Offset returns the offset of this clock relative to the system clock.
// This can be used to convert to/from system time.
func (c *clock) Offset() time.Duration {
	return c.offset
}

var goClock = &systemClock{}

// System is a Clock that uses the system clock, meaning it just delegates to time.Now() etc.
func System() Clock {
	return goClock
}

type systemClock struct {
}

func (c *systemClock) Now() time.Time {
	return time.Now()
}

func (c *systemClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (c *systemClock) Until(t time.Time) time.Duration {
	return time.Until(t)
}

func (c *systemClock) Offset() time.Duration {
	return 0
}
