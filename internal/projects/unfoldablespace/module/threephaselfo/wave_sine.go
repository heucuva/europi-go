package threephaselfo

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type waveSine struct{}

func (waveSine) Get(t, interval time.Duration) units.CV {
	x0 := float32(2.0 * math.Pi * t.Seconds() / interval.Seconds())
	return units.CV((math.Sin(float64(x0)) + 1.0) / 2.0)
}
