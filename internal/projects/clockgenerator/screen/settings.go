package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	europim "github.com/heucuva/europi/math"
)

type Settings struct {
	Clock           *module.ClockGenerator
	MinBPM          float32
	MaxBPM          float32
	MinGateDuration time.Duration
	MaxGateDuration time.Duration
	kb              *knobbank.KnobBank
}

func (m *Settings) Start(e *europi.EuroPi) {
	bpmT := europim.InverseLerp(m.Clock.BPM(), m.MinBPM, m.MaxBPM)
	gateT := europim.InverseLerp(m.Clock.GateDuration(), m.MinGateDuration, m.MaxGateDuration)

	var err error
	m.kb, err = knobbank.NewKnobBank(
		e.K1,
		knobbank.WithDisabledKnob(),
		knobbank.WithLockedKnob("bpm", knobbank.InitialPercentageValue(bpmT)),
		knobbank.WithLockedKnob("gate", knobbank.InitialPercentageValue(gateT)),
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
		bpmSel  rune = ' '
		gateSel rune = ' '
	)
	switch m.kb.CurrentName() {
	case "bpm":
		bpmSel = '*'
	case "gate":
		gateSel = '*'
	default:
	}

	disp.WriteLine(fmt.Sprintf("%cBPM:%3.1f", bpmSel, m.Clock.BPM()), 0, line1y)

	gate := m.Clock.GateDuration()
	var gateStr string
	switch {
	case gate < time.Millisecond:
		gate -= (gate % time.Microsecond)
		gateStr = fmt.Sprintf("%3.1fus", gate.Seconds()*1000000.0)
	case gate < time.Second:
		gate -= (gate % time.Millisecond)
		gateStr = fmt.Sprintf("%3.1fms", gate.Seconds()*1000.0)
	default:
		gateStr = fmt.Sprint(gate)
	}
	disp.WriteLine(fmt.Sprintf("%cGate:%s", gateSel, gateStr), 0, line2y)
}

func (m *Settings) updateMenu(e *europi.EuroPi) {
	switch m.kb.CurrentName() {
	case "bpm":
		m.Clock.SetBPM(europim.Lerp[float32](m.kb.Percent(), m.MinBPM, m.MaxBPM))
	case "gate":
		m.Clock.SetGateDuration(europim.Lerp(m.kb.Percent(), m.MinGateDuration, m.MaxGateDuration))
	default:
	}
}
