package threephaselfo

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type Module struct {
	t        time.Duration
	interval time.Duration
	wave     wave
	degree0  func(cv units.CV)
}

func (m *Module) Init(config Config) error {
	m.degree0 = config.Degree0
	if m.degree0 == nil {
		m.degree0 = func(cv units.CV) {}
	}

	m.interval = europim.Lerp(config.Phi3Rate.ToFloat32(), 1, time.Second/4)

	var err error
	m.wave, err = m.getWaveMode(config.WaveMode)
	return err
}

func (m *Module) Tick(deltaTime time.Duration) {
	t := (m.t + deltaTime) % m.interval
	cv0 := m.wave.Get(t, m.interval)
	m.degree0(cv0)
	m.t = t
}
