package module

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type waveSine struct{}

func (waveSine) Get(t, interval time.Duration) (units.CV, units.CV, units.CV) {
	const period float32 = 2.0 * math.Pi
	intv := float32(interval.Seconds())

	phasePos := intv / 3.0

	tPos := float32(t.Seconds()) / intv

	x0 := tPos * period
	x120 := x0 + phasePos
	x240 := x120 + phasePos

	cv0 := units.CV((math.Sin(float64(x0)) + 1.0) / 2.0)
	cv120 := units.CV((math.Sin(float64(x120)) + 1.0) / 2.0)
	cv240 := units.CV((math.Sin(float64(x240)) + 1.0) / 2.0)
	return cv0, cv120, cv240
}

func (waveSine) Mode() WaveMode {
	return WaveModeSine
}
