package screen

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/knobmenu"
	"github.com/awonak/EuroPiGo/internal/projects/complexarp/module"
	"github.com/awonak/EuroPiGo/units"
)

type Settings struct {
	km         *knobmenu.KnobMenu
	ComplexArp *module.ComplexArp
}

func (m *Settings) scaleString() string {
	return module.ScaleString(m.ComplexArp.Scale())
}

func (m *Settings) scaleValue() units.CV {
	return module.ScaleToCV(m.ComplexArp.Scale())
}

func (m *Settings) setScaleValue(value units.CV) {
	m.ComplexArp.SetScale(module.CVToScale(value))
}

func (m *Settings) arpPitchString() string {
	return module.ArpPitchString(m.ComplexArp.ArpPitch())
}

func (m *Settings) arpPitchValue() units.CV {
	return module.ArpPitchToCV(m.ComplexArp.ArpPitch())
}

func (m *Settings) setArpPitchValue(value units.CV) {
	m.ComplexArp.SetArpPitch(module.CVToArpPitch(value))
}

func (m *Settings) arpRangeString() string {
	return module.ArpRangeString(m.ComplexArp.ArpRange())
}

func (m *Settings) arpRangeValue() units.CV {
	return module.ArpRangeToCV(m.ComplexArp.ArpRange())
}

func (m *Settings) setArpRangeValue(value units.CV) {
	m.ComplexArp.SetArpRange(module.CVToArpRange(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("scale", "Scale", m.scaleString, m.scaleValue, m.setScaleValue),
		knobmenu.WithItem("arpPitch", "Pitch", m.arpPitchString, m.arpPitchValue, m.setArpPitchValue),
		knobmenu.WithItem("arpRange", "Range", m.arpRangeString, m.arpRangeValue, m.setArpRangeValue),
	)
	if err != nil {
		panic(err)
	}

	m.km = km
}

func (m *Settings) Button1Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Settings) Button1(e *europi.EuroPi, _ time.Duration) {
	m.km.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
}
