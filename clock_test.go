package clock

import (
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp"
)

const timeLayout = "2006-01-02-15:04:05"

var durationEq = qt.CmpEquals(
	cmp.Comparer(func(x, y time.Duration) bool {
		return x.Truncate(1*time.Second) == y.Truncate(1*time.Second)
	}),
)

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

	c.Run("Since", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)

		time.Sleep(3 * time.Second)
		c.Assert(clock.Since(start), durationEq, time.Duration(3*time.Second))
	})

	c.Run("Until", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)
		then := clock.Now().Add(3010 * time.Millisecond)
		c.Assert(clock.Until(then), durationEq, time.Duration(3*time.Second))
	})

}

func toString(t time.Time) string {
	return t.UTC().Format(timeLayout)
}
