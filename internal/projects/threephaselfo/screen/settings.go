package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/internal/projects/threephaselfo/module"
)

type Settings struct {
	LFO *module.ThreePhaseLFO
	kb  *knobbank.KnobBank
}

func (m *Settings) Start(e *europi.EuroPi) {
	modeT := module.GetWaveModeCV(m.LFO.Wave())

	var err error
	m.kb, err = knobbank.NewKnobBank(
		e.K1,
		knobbank.WithDisabledKnob(),
		knobbank.WithLockedKnob("wave", knobbank.InitialPercentageValue(modeT.ToFloat32())),
		knobbank.WithLockedKnob("rate", knobbank.InitialPercentageValue(m.LFO.Rate().ToFloat32())),
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
		waveSelrune = ' '
		rateSelrune = ' '
	)
	switch m.kb.CurrentName() {
	case "wave":
		waveSelrune = '*'
	case "rate":
		rateSelrune = '*'
	default:
	}

	var (
		waveName string
	)
	switch m.LFO.Wave() {
	case module.WaveModeSine:
		waveName = "sine"
	default:
		waveName = "unk"
	}
	disp.WriteLine(fmt.Sprintf("%cWave:%5s", waveSelrune, waveName), 0, line1y)
	disp.WriteLine(fmt.Sprintf("%cRate:%5.2fHz", rateSelrune, m.LFO.RateHz()), 0, line2y)
}

func (m *Settings) updateMenu(e *europi.EuroPi) {
	switch m.kb.CurrentName() {
	case "wave":
		m.LFO.SetWaveCV(m.kb.ReadCV())
	case "rate":
		m.LFO.SetRate(m.kb.ReadCV())
	default:
	}
}
