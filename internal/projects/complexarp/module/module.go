package module

import (
	"time"

	"github.com/heucuva/europi/experimental/quantizer"
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type ComplexArp struct {
	out      func(voct units.VOct)
	patRange units.VOct
	patPitch units.VOct

	scale      scale
	quantizer  quantizer.Quantizer[units.VOct]
	arpPattern pattern
}

func (m *ComplexArp) Init(config Config) error {
	m.out = config.ArpOut
	m.patRange = config.ArpRange
	m.patPitch = config.ArpPitch

	if err := m.setScale(config.Scale); err != nil {
		return err
	}

	if err := m.setArpQuantizer(config.Quantizer); err != nil {
		return err
	}

	if err := m.setArpPattern(config); err != nil {
		return err
	}

	return nil
}

func (m *ComplexArp) ArpClock(high bool) {
	if high {
		voct := m.arpPattern.Next(m)
		m.out(voct)
	}
}

func (m *ComplexArp) SetArpPitch(voct units.VOct) {
	m.patPitch = voct
}

func (m *ComplexArp) SetArpRange(voct units.VOct) {
	m.patRange = voct
}

func (m *ComplexArp) ArpRange() units.VOct {
	return m.patRange
}

func (m *ComplexArp) SetScaleCV(cv units.CV) {
	mode := europim.Lerp(cv.ToFloat32(), ScaleC_Lydian, cScaleCount-1)
	if m.scale.Mode() == mode {
		// no change
		return
	}

	if err := m.setScale(mode); err != nil {
		panic(err)
	}

	m.arpPattern.UpdateScale(m.scale)
}

func (m *ComplexArp) Scale() Scale {
	return m.scale.Mode()
}

func (m *ComplexArp) ScaleName() string {
	return m.scale.Name()
}

func (m *ComplexArp) Tick(deltaTime time.Duration) {
}
