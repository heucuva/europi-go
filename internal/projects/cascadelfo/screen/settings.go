package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/cascadelfo/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km  *knobmenu.KnobMenu
	LFO *module.CascadeLFO
}

func (m *Settings) rateString() string {
	return module.RateString(m.LFO.Rate())
}

func (m *Settings) rateValue() units.CV {
	return module.RateToCV(m.LFO.Rate())
}

func (m *Settings) setRateValue(value units.CV) {
	m.LFO.SetRate(module.CVToRate(value))
}

func (m *Settings) rateAttenuverterString() string {
	return module.RateAttenuverterString(m.LFO.RateAttenuverter())
}

func (m *Settings) rateAttenuverterValue() units.CV {
	return module.RateAttenuverterToCV(m.LFO.RateAttenuverter())
}

func (m *Settings) setRateAttenuverterValue(value units.CV) {
	m.LFO.SetRateAttenuverter(module.CVToRateAttenuverter(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("rate", "Rate", m.rateString, m.rateValue, m.setRateValue),
		knobmenu.WithItem("rateAttenuverter", "R.AV", m.rateAttenuverterString, m.rateAttenuverterValue, m.setRateAttenuverterValue),
	)
	if err != nil {
		panic(err)
	}

	m.km = km
}

func (m *Settings) Button1(e *europi.EuroPi, p machine.Pin) {
	m.km.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
}
