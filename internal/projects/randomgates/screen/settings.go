package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/randomgates/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	RandomGates *module.RandomGates
	km          *knobmenu.KnobMenu
}

func (m *Settings) modeString() string {
	return module.ModeToString(m.RandomGates.Mode())
}

func (m *Settings) modeValue() units.CV {
	return module.ModeToCV(m.RandomGates.Mode())
}

func (m *Settings) setModeValue(value units.CV) {
	m.RandomGates.SetMode(module.CVToMode(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("mode", "Mode", m.modeString, m.modeValue, m.setModeValue),
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
