package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type CascadeLFO struct {
	cv     units.BipolarCV
	lfo    [8]lfo
	rate   units.Hertz
	rateAV units.BipolarCV
}

func noop(_ units.BipolarCV) {
}

func (m *CascadeLFO) Init(config Config) error {
	m.rate = config.Rate
	m.rateAV = config.RateAttenuverter
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

func (m *CascadeLFO) SetCV(cv units.BipolarCV) {
	m.cv = cv
}

func (m *CascadeLFO) SetAttenuverter(av units.BipolarCV) {
	m.rateAV = av
}

func (m *CascadeLFO) Attenuverter() units.BipolarCV {
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
