package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type CascadeLFO struct {
	cv     units.CV
	lfo    [8]lfo
	rate   units.Hertz
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
	m.cv = 0.5
	return nil
}

func (m *CascadeLFO) SetCV(cv units.CV) {
	m.cv = cv
}

func (m *CascadeLFO) SetAttenuverter(av float32) {
	m.rateAV = av
}

func (m *CascadeLFO) Attenuverter() float32 {
	return m.rateAV
}

func (m *CascadeLFO) SetRate(rate units.Hertz) {
	m.rate = rate
}

func (m *CascadeLFO) Rate() units.Hertz {
	return m.rate
}

func (m *CascadeLFO) Tick(deltaTime time.Duration) {
	period := AdjustRate(m.rate, m.cv, m.rateAV).ToPeriod()
	delta := float32(period.Seconds() * deltaTime.Seconds())
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Update(delta)
		delta /= 2.0
	}
}

func (m *CascadeLFO) Reset() {
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Reset()
	}
}
