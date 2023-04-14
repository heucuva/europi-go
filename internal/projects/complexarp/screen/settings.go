package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/complexarp/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km         *knobmenu.KnobMenu
	ComplexArp *module.ComplexArp
}

func (m *Settings) scaleString() string {
	return module.ScaleToString(m.ComplexArp.Scale())
}

func (m *Settings) scaleValue() units.CV {
	return module.ScaleToCV(m.ComplexArp.Scale())
}

func (m *Settings) setScaleCV(value units.CV) {
	m.ComplexArp.SetScale(module.CVToScale(value))
}

func (m *Settings) pitchString() string {
	return module.PitchToString(m.ComplexArp.ArpPitch())
}

func (m *Settings) pitchValue() units.CV {
	return module.PitchToCV(m.ComplexArp.ArpPitch())
}

func (m *Settings) setPitchCV(value units.CV) {
	m.ComplexArp.SetArpPitch(module.CVToPitch(value))
}

func (m *Settings) rangeString() string {
	return module.RangeToString(m.ComplexArp.ArpRange())
}

func (m *Settings) rangeValue() units.CV {
	return module.RangeToCV(m.ComplexArp.ArpRange())
}

func (m *Settings) setRangeCV(value units.CV) {
	m.ComplexArp.SetArpRange(module.CVToRange(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("scale", "Scale", m.scaleString, m.scaleValue, m.setScaleCV),
		knobmenu.WithItem("pitch", "Pitch", m.pitchString, m.pitchValue, m.setPitchCV),
		knobmenu.WithItem("range", "Range", m.rangeString, m.rangeValue, m.setRangeCV),
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
