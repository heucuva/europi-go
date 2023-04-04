package threephaselfo

import (
	"math"
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type Module struct {
	t        time.Duration
	interval time.Duration
	degree0  func(cv units.CV)
}

func (m *Module) Init(config Config) error {
	m.degree0 = config.Degree0
	if m.degree0 == nil {
		m.degree0 = func(cv units.CV) {}
	}

	m.interval = europim.Lerp(config.Phi3Rate.ToFloat32(), 1, time.Second/4)
	return nil
}

func (m *Module) Tick(deltaTime time.Duration) {
	t := (m.t + deltaTime) % m.interval

	x0 := float32(2.0 * math.Pi * t.Seconds() / m.interval.Seconds())
	cv0 := units.CV((math.Sin(float64(x0)) + 1.0) / 2.0)

	m.degree0(cv0)
	m.t = t
}
