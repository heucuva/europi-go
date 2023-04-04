package cascadelfo

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type Module struct {
	delta  float32
	lfo    [8]lfo
	rate   units.CV
	rateAV units.CV
}

func (m *Module) Init(config Config) error {
	m.rate = config.Rate
	m.rateAV = config.RateAttenuverter
	for i := range m.lfo {
		o := &m.lfo[i]
		f := config.LFO[i]
		if f == nil {
			f = func(cv units.CV) {}
		}
		o.out = f
	}
	m.SetCV(0)
	return nil
}

func (m *Module) SetCV(cv units.CV) {
	rate := ((m.rate + cv) * m.rateAV).ToFloat32()
	// best guess on calculating step rate of final lfo
	m.delta = 1500.0 * float32(math.Pow(1875.0, float64(-rate)))
}

func (m *Module) Tick(deltaTime time.Duration) {
	d := m.delta * float32(deltaTime.Seconds())
	for i := range m.lfo {
		o := &m.lfo[len(m.lfo)-i-1]
		o.Update(d)
		d /= 2.0
	}
}
