package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type CascadeLFO struct {
	cv               units.BipolarCV
	lfo              [8]lfo
	rate             units.Hertz
	rateAttenuverter units.BipolarCV
}

func (m *CascadeLFO) Init(config Config) error {
	m.cv = 0.5
	m.rate = config.Rate
	m.rateAttenuverter = config.RateAttenuverter

	for i, fn := range config.LFO {
		o := &m.lfo[i]
		if fn == nil {
			fn = noopLFO
		}
		o.out = fn
	}

	return nil
}

func noopLFO(cv units.BipolarCV) {
}

func (m *CascadeLFO) SetCV(cv units.BipolarCV) {
	m.cv = cv
}

func (m *CascadeLFO) Reset() {
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Reset()
	}
}

func (m *CascadeLFO) SetRate(freq units.Hertz) {
	m.rate = freq
}

func (m *CascadeLFO) Rate() units.Hertz {
	return m.rate
}

func (m *CascadeLFO) SetRateAttenuverter(cv units.BipolarCV) {
	m.rateAttenuverter = cv
}

func (m *CascadeLFO) RateAttenuverter() units.BipolarCV {
	return m.rateAttenuverter
}

func (m *CascadeLFO) Tick(deltaTime time.Duration) {
	period := AdjustRate(m.rate, m.cv, m.rateAttenuverter).ToPeriod()
	delta := float32(period.Seconds() * deltaTime.Seconds())
	for i := range m.lfo {
		o := &m.lfo[i]
		o.Update(delta)
		delta /= 2.0
	}
}
