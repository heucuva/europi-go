package module

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type ThreePhaseLFO struct {
	t         time.Duration
	interval  time.Duration
	wave      wave
	degree0   func(cv units.CV)
	degree120 func(cv units.CV)
	degree240 func(cv units.CV)
}

func noop(_ units.CV) {
}

func (m *ThreePhaseLFO) Init(config Config) error {
	m.degree0 = config.Degree0
	if m.degree0 == nil {
		m.degree0 = noop
	}

	m.degree120 = config.Degree120
	if m.degree120 == nil {
		m.degree120 = noop
	}

	m.degree240 = config.Degree240
	if m.degree240 == nil {
		m.degree240 = noop
	}

	m.SetRate(config.Phi3Rate)

	var err error
	m.wave, err = m.getWaveMode(config.WaveMode)
	return err
}

func (m *ThreePhaseLFO) SetRate(cv units.CV) {
	m.interval = europim.Lerp(cv.ToFloat32(), 1, time.Second/4)
}

func (m *ThreePhaseLFO) SetWaveCV(cv units.CV) {
	mode := europim.Lerp(cv.ToFloat32(), WaveModeSine, cWaveModeCount-1)
	if mode == m.wave.Mode() {
		// no change
		return
	}

	wave, err := m.getWaveMode(mode)
	if err != nil {
		panic(err)
	}

	m.wave = wave
}

func (m *ThreePhaseLFO) Wave() WaveMode {
	return m.wave.Mode()
}

func (m *ThreePhaseLFO) Tick(deltaTime time.Duration) {
	t := (m.t + deltaTime) % m.interval
	cv0, cv120, cv240 := m.wave.Get(t, m.interval)
	m.t = t
	m.degree0(cv0)
	m.degree120(cv120)
	m.degree240(cv240)
}
