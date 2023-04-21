package module

import (
	"time"

	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/units"
)

type ComplexArp struct {
	arpOut     func(voct units.VOct)
	arpPattern pattern
	scale      scale
	quantizer  quantizer.Quantizer[units.VOct]
	quantMode  quantizer.Mode
	arpPitch   units.VOct
	arpRange   units.VOct
}

func (m *ComplexArp) Init(config Config) error {
	fnArpOut := config.ArpOut
	if fnArpOut == nil {
		fnArpOut = noopArpOut
	}
	m.arpOut = fnArpOut

	m.arpPitch = config.ArpPitch
	m.arpRange = config.ArpRange

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

func noopArpOut(voct units.VOct) {
}

func (m *ComplexArp) ArpClock(high bool) {
	if high {
		voct := m.arpPattern.Next(m)
		m.arpOut(voct)
	}
}

func (m *ComplexArp) SetArpPattern(pat Pattern) {
	if err := m.setArpPattern(Config{
		ArpPattern: pat,
	}); err != nil {
		panic(err)
	}
}

func (m *ComplexArp) ArpPattern() Pattern {
	return m.arpPattern.Pattern()
}

func (m *ComplexArp) SetScale(s Scale) {
	if s == m.scale.Mode() {
		// no change
		return
	}

	if err := m.setScale(s); err != nil {
		panic(err)
	}

	m.arpPattern.UpdateScale(m.scale)
}

func (m *ComplexArp) Scale() Scale {
	return m.scale.Mode()
}

func (m *ComplexArp) SetQuantizer(q quantizer.Mode) {
	if q == m.quantMode {
		// no change
		return
	}

	if err := m.setArpQuantizer(q); err != nil {
		panic(err)
	}
	m.quantMode = q
}

func (m *ComplexArp) Quantizer() quantizer.Mode {
	return m.quantMode
}

func (m *ComplexArp) SetArpPitch(voct units.VOct) {
	m.arpPitch = voct
}

func (m *ComplexArp) ArpPitch() units.VOct {
	return m.arpPitch
}

func (m *ComplexArp) SetArpRange(voct units.VOct) {
	m.arpRange = voct
}

func (m *ComplexArp) ArpRange() units.VOct {
	return m.arpRange
}

func (m *ComplexArp) Tick(deltaTime time.Duration) {
}
