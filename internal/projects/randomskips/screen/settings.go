package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/internal/projects/randomskips/module"
)

type Settings struct {
	RandomSkips *module.RandomSkips
	kb          *knobbank.KnobBank
}

func (m *Settings) Start(e *europi.EuroPi) {
	chanceT := m.RandomSkips.Chance()

	var err error
	m.kb, err = knobbank.NewKnobBank(
		e.K1,
		knobbank.WithDisabledKnob(),
		knobbank.WithLockedKnob("chance", knobbank.InitialPercentageValue(chanceT)),
	)
	if err != nil {
		panic(err)
	}
}

func (m *Settings) Button1(e *europi.EuroPi, p machine.Pin) {
	m.kb.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.updateMenu(e)

	disp := e.Display

	var (
		chanceSelrune = ' '
	)
	switch m.kb.CurrentName() {
	case "chance":
		chanceSelrune = '*'
	default:
	}

	disp.WriteLine(fmt.Sprintf("%cChance:%3.1f%%", chanceSelrune, m.RandomSkips.Chance()*100.0), 0, line1y)
}

func (m *Settings) updateMenu(e *europi.EuroPi) {
	switch m.kb.CurrentName() {
	case "chance":
		m.RandomSkips.SetChance(m.kb.Percent())
	default:
	}
}
