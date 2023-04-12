package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/internal/projects/randomgates/module"
)

type Settings struct {
	RandomGates *module.RandomGates
	kb          *knobbank.KnobBank
	modeQuant   quantizer.Round[module.Mode]
}

const (
	modeCount = module.ModeEqualGateTrig + 1
)

var (
	modeList = []module.Mode{
		module.Mode1msTrig,
		module.Mode200msTrig,
		module.ModeQuarterGateTrig,
		module.ModeHalfGateTrig,
		module.ModeEqualGateTrig,
	}
)

func (m *Settings) Start(e *europi.EuroPi) {
	modeT := float32(m.RandomGates.Mode()) / float32(modeCount)

	var err error
	m.kb, err = knobbank.NewKnobBank(
		e.K1,
		knobbank.WithDisabledKnob(),
		knobbank.WithLockedKnob("mode", knobbank.InitialPercentageValue(modeT)),
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
		modeSelrune = ' '
	)
	switch m.kb.CurrentName() {
	case "mode":
		modeSelrune = '*'
	default:
	}

	var (
		modeName string
	)
	switch m.RandomGates.Mode() {
	case module.Mode1msTrig:
		modeName = "1ms"
	case module.Mode200msTrig:
		modeName = "200ms"
	case module.ModeQuarterGateTrig:
		modeName = "1/4"
	case module.ModeHalfGateTrig:
		modeName = "1/2"
	case module.ModeEqualGateTrig:
		modeName = "1:1"
	default:
		modeName = "unk"
	}
	disp.WriteLine(fmt.Sprintf("%cMode:%5s", modeSelrune, modeName), 0, line1y)
}

func (m *Settings) updateMenu(e *europi.EuroPi) {
	switch m.kb.CurrentName() {
	case "mode":
		mode := m.modeQuant.QuantizeToValue(m.kb.Percent(), modeList)
		m.RandomGates.SetMode(mode)
	default:
	}
}
