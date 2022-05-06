package clock

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

const timeLayout = "2006-01-02-15:04:05"

func TestClock(t *testing.T) {
	c := qt.New(t)

	c.Run("Past", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)

		c.Assert(toString(clock.Now()), qt.Equals, "2019-10-11-02:50:01")
		time.Sleep(3 * time.Second)
		c.Assert(toString(clock.Now()), qt.Equals, "2019-10-11-02:50:04")
	})

	c.Run("Future", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2053-10-11-02:50:01")
		clock := Start(start)

		c.Assert(toString(clock.Now()), qt.Equals, "2053-10-11-02:50:01")
		time.Sleep(3 * time.Second)
		c.Assert(toString(clock.Now()), qt.Equals, "2053-10-11-02:50:04")
	})

}

func toString(t time.Time) string {
	return t.UTC().Format(timeLayout)
}
