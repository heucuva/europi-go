package module

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type ComplexArp struct {
	out        func(voct units.VOct)
	arpPattern pattern
}

func (m *ComplexArp) Init(config Config) error {
	var err error
	m.arpPattern, err = getArpPattern(config)
	if err != nil {
		return err
	}

	m.out = config.ArpOut

	return nil
}

func (m *ComplexArp) ArpClock(high bool) {
	if high {
		voct := m.arpPattern.Next()
		m.out(voct)
	}
}

func (m *ComplexArp) SetArpPitch(voct units.VOct) {
	m.arpPattern.SetArpPitch(voct)
}

func (m *ComplexArp) SetArpRange(voct units.VOct) {
	m.arpPattern.SetArpRange(voct)
}

func (m *ComplexArp) SetScaleCV(cv units.CV) {
	mode := europim.Lerp(cv.ToFloat32(), ScaleC_Lydian, cScaleCount-1)
	if m.arpPattern.Scale() == mode {
		// no change
		return
	}

	scale, err := getScale(mode)
	if err != nil {
		panic(err)
	}

	m.arpPattern.SetScale(scale)
}

func (m *ComplexArp) Scale() Scale {
	return m.arpPattern.Scale()
}

func (m *ComplexArp) ScaleName() string {
	return m.arpPattern.ScaleName()
}

func (m *ComplexArp) ArpRange() units.VOct {
	return m.arpPattern.ArpRange()
}

func (m *ComplexArp) Tick(deltaTime time.Duration) {
}
