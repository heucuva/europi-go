package module

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type CascadeLFO struct {
	delta  float32
	lfo    [8]lfo
	rate   float32
	rateAV float32
}

func noop(_ units.CV) {
}

func (m *CascadeLFO) Init(config Config) error {
	m.rate = config.Rate
	m.SetAttenuverter(config.RateAttenuverter)
	for i := range m.lfo {
		o := &m.lfo[i]
		f := config.LFO[i]
		if f == nil {
			f = noop
		}
		o.out = f
	}
	m.SetCV(0.5)
	return nil
}

func (m *CascadeLFO) SetCV(cv units.CV) {
	ai := cv.ToFloat32()*2.0 - 1.0
	rate := RateToCV(m.rate).ToFloat32() + ai*m.rateAV
	m.delta = CVToRate(units.CV(rate)) * float32(len(lfoTriangle)) * (2.0 * math.Pi)
}

func (m *CascadeLFO) SetAttenuverter(av float32) {
	m.rateAV = av
}

func (m *CascadeLFO) Attenuverter() float32 {
	return m.rateAV
}

func (m *CascadeLFO) SetRate(rate float32) {
	m.rate = rate
}

func (m *CascadeLFO) Rate() float32 {
	return m.rate
}

func (m *CascadeLFO) Tick(deltaTime time.Duration) {
	d := m.delta * float32(deltaTime.Seconds())
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Update(d)
		d /= 2.0
	}
}

func (m *CascadeLFO) Reset() {
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Reset()
	}
}
