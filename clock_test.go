package clock

import (
	"math"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp"
)

const timeLayout = "2006-01-02-15:04:05"

var durationEq = qt.CmpEquals(
	cmp.Comparer(func(x, y time.Duration) bool {
		xs := math.RoundToEven(float64(x) / float64(time.Second))
		ys := math.RoundToEven(float64(y) / float64(time.Second))
		return xs == ys
	}),
)

func TestClock(t *testing.T) {
	c := qt.New(t)

	c.Run("Past", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)

		c.Assert(toString(clock.Now()), qt.Equals, "2019-10-11-02:50:01")
		time.Sleep(1 * time.Second)
		c.Assert(toString(clock.Now()), qt.Equals, "2019-10-11-02:50:02")
	})

	c.Run("Future", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2053-10-11-02:50:01")
		clock := Start(start)

		c.Assert(toString(clock.Now()), qt.Equals, "2053-10-11-02:50:01")
		time.Sleep(1 * time.Second)
		c.Assert(toString(clock.Now()), qt.Equals, "2053-10-11-02:50:02")
	})

	c.Run("Offset", func(c *qt.C) {
		c.Parallel()

		clock := Start(time.Now().Add(5010 * time.Millisecond))
		c.Assert(clock.Offset(), durationEq, time.Duration(5*time.Second))
	})

	c.Run("Since", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)

		time.Sleep(1 * time.Second)
		c.Assert(clock.Since(start), durationEq, time.Duration(1*time.Second))
	})

	c.Run("Until", func(c *qt.C) {
		c.Parallel()

		start, _ := time.Parse(timeLayout, "2019-10-11-02:50:01")
		clock := Start(start)
		then := clock.Now().Add(3010 * time.Millisecond)
		c.Assert(clock.Until(then), durationEq, time.Duration(3*time.Second))
	})
}

func TestSystemClock(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	c.Assert(toString(System().Now()), qt.Equals, toString(time.Now()))
	c.Assert(System().Since(time.Now().Add(-10*time.Hour)), durationEq, time.Since(time.Now().Add(-10*time.Hour)))
	c.Assert(System().Until(time.Now().Add(10*time.Hour)), durationEq, time.Until(time.Now().Add(10*time.Hour)))
	c.Assert(System().Offset(), qt.Equals, time.Duration(0))
}

func TestFixedClock(t *testing.T) {
	t.Parallel()

	c := qt.New(t)

	fixed := Fixed(TimeCupFinalNorway1976)
	fiveSecondsLater := Fixed(TimeCupFinalNorway1976.Add(5 * time.Second))
	fiveSecondsBefore := Fixed(TimeCupFinalNorway1976.Add(-5 * time.Second))

	c.Assert(toString(fixed.Now()), qt.Equals, toString(TimeCupFinalNorway1976))
	c.Assert(fixed.Offset() > 0, qt.IsTrue)
	c.Assert(fixed.Since(fiveSecondsBefore.Now()), durationEq, time.Duration(5*time.Second))
	c.Assert(fixed.Until(fiveSecondsLater.Now()), durationEq, time.Duration(5*time.Second))

}

func toString(t time.Time) string {
	return t.UTC().Format(timeLayout)
}
