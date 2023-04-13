package module

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type waveSine struct{}

const (
	twoPi      = 2.0 * math.Pi
	twoPiThird = twoPi / 3.0
)

func (waveSine) Get(t, interval time.Duration) (units.CV, units.CV, units.CV) {
	x0 := float32(t.Seconds() / interval.Seconds() * twoPi)
	x120 := x0 + twoPiThird
	x240 := x0 - twoPiThird

	cv0 := units.CV((math.Sin(float64(x0)) + 1.0) / 2.0)
	cv120 := units.CV((math.Sin(float64(x120)) + 1.0) / 2.0)
	cv240 := units.CV((math.Sin(float64(x240)) + 1.0) / 2.0)
	return cv0, cv120, cv240
}

func (waveSine) Mode() WaveMode {
	return WaveModeSine
}
