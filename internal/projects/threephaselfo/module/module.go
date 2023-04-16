package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type ThreePhaseLFO struct {
	degree0   func(cv units.BipolarCV)
	degree120 func(cv units.BipolarCV)
	degree240 func(cv units.BipolarCV)
	waveMode  WaveMode
	phi3Rate  units.Hertz
	skewRate  units.Hertz
	skewShape units.CV

	t        time.Duration
	interval time.Duration
	wave     wave
}

func (m *ThreePhaseLFO) Init(config Config) error {
	fnDegree0 := config.Degree0
	if fnDegree0 == nil {
		fnDegree0 = noopDegree0
	}
	m.degree0 = fnDegree0

	fnDegree120 := config.Degree120
	if fnDegree120 == nil {
		fnDegree120 = noopDegree120
	}
	m.degree120 = fnDegree120

	fnDegree240 := config.Degree240
	if fnDegree240 == nil {
		fnDegree240 = noopDegree240
	}
	m.degree240 = fnDegree240

	m.skewRate = config.SkewRate
	m.skewShape = config.SkewShape

	m.SetPhi3Rate(config.Phi3Rate)

	var err error
	m.waveMode = config.WaveMode + 1
	m.wave, err = getWaveMode(config.WaveMode)
	return err
}

func noopDegree0(cv units.BipolarCV) {
}

func noopDegree120(cv units.BipolarCV) {
}

func noopDegree240(cv units.BipolarCV) {
}

func (m *ThreePhaseLFO) Reset() {
	panic("unimplemented")
}

func (m *ThreePhaseLFO) SetWaveMode(mode WaveMode) {
	if mode == m.waveMode {
		// no change
		return
	}

	wave, err := getWaveMode(mode)
	if err != nil {
		panic(err)
	}

	m.wave = wave
	m.waveMode = mode
}

func (m *ThreePhaseLFO) WaveMode() WaveMode {
	return m.waveMode
}

func (m *ThreePhaseLFO) SetPhi3Rate(freq units.Hertz) {
	m.phi3Rate = freq
	m.interval = freq.ToPeriod()
}

func (m *ThreePhaseLFO) Phi3Rate() units.Hertz {
	return m.phi3Rate
}

func (m *ThreePhaseLFO) SetSkewRate(freq units.Hertz) {
	m.skewRate = freq
}

func (m *ThreePhaseLFO) SkewRate() units.Hertz {
	return m.skewRate
}

func (m *ThreePhaseLFO) SetSkewShape(cv units.CV) {
	m.skewShape = cv
}

func (m *ThreePhaseLFO) SkewShape() units.CV {
	return m.skewShape
}

func (m *ThreePhaseLFO) Tick(deltaTime time.Duration) {
	t := (m.t + deltaTime) % m.interval
	cv0, cv120, cv240 := m.wave.Get(t, m.interval)
	m.t = t
	m.degree0(cv0)
	m.degree120(cv120)
	m.degree240(cv240)
}
