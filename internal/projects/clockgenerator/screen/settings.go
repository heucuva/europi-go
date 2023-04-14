package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km    *knobmenu.KnobMenu
	Clock *module.ClockGenerator
}

func (m *Settings) bpmString() string {
	return module.BPMToString(m.Clock.BPM())
}

func (m *Settings) bpmValue() units.CV {
	return module.BPMToCV(m.Clock.BPM())
}

func (m *Settings) setBpmValue(value units.CV) {
	m.Clock.SetBPM(module.CVToBPM(value))
}

func (m *Settings) gateString() string {
	return module.GateDurationToString(m.Clock.GateDuration())
}

func (m *Settings) gateValue() units.CV {
	return module.GateDurationToCV(m.Clock.GateDuration())
}

func (m *Settings) setGateValue(value units.CV) {
	m.Clock.SetGateDuration(module.CVToGateDuration(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("bpm", "BPM", m.bpmString, m.bpmValue, m.setBpmValue),
		knobmenu.WithItem("gate", "Gate", m.gateString, m.gateValue, m.setGateValue),
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
