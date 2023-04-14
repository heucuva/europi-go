package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/threephaselfo/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	LFO *module.ThreePhaseLFO
	km  *knobmenu.KnobMenu
}

func (m *Settings) waveString() string {
	return module.WaveModeToString(m.LFO.Wave())
}

func (m *Settings) waveValue() units.CV {
	return module.WaveModeToCV(m.LFO.Wave())
}

func (m *Settings) setWaveValue(value units.CV) {
	m.LFO.SetWave(module.CVToWaveMode(value))
}

func (m *Settings) rateString() string {
	return module.RateToString(m.LFO.Rate())
}

func (m *Settings) rateValue() units.CV {
	return module.RateToCV(m.LFO.Rate())
}

func (m *Settings) setRateValue(value units.CV) {
	m.LFO.SetRate(module.CVToRate(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("wave", "Wave", m.waveString, m.waveValue, m.setWaveValue),
		knobmenu.WithItem("rate", "Rate", m.rateString, m.rateValue, m.setRateValue),
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
