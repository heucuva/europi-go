package module

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type CascadeLFO struct {
	delta  float32
	lfo    [8]lfo
	rate   units.CV
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

const (
	lfoConstA = 218.45333333333333333333333333248
	lfoConstB = 5.23377884541046550578029111514
)

func (m *CascadeLFO) SetCV(cv units.CV) {
	ai := cv.ToFloat32()*2.0 - 1.0
	rate := m.rate.ToFloat32() + ai*m.rateAV
	// best guess on calculating step rate of final lfo
	m.delta = float32(lfoConstA * math.Exp(lfoConstB*float64(rate)))
}

func (m *CascadeLFO) SetAttenuverter(cv units.CV) {
	m.rateAV = cv.ToFloat32()*2.0 - 1.0
}

func (m *CascadeLFO) SetRate(cv units.CV) {
	m.rate = cv
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
