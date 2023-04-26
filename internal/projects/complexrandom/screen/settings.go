package screen

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/knobmenu"
	"github.com/awonak/EuroPiGo/internal/projects/complexrandom/module"
	"github.com/awonak/EuroPiGo/units"
)

type Settings struct {
	km            *knobmenu.KnobMenu
	ComplexRandom *module.ComplexRandom
}

func (m *Settings) sampleAttenuatorAString() string {
	return module.SampleAttenuatorAString(m.ComplexRandom.SampleAttenuatorA())
}

func (m *Settings) sampleAttenuatorAValue() units.CV {
	return module.SampleAttenuatorAToCV(m.ComplexRandom.SampleAttenuatorA())
}

func (m *Settings) setSampleAttenuatorAValue(value units.CV) {
	m.ComplexRandom.SetSampleAttenuatorA(module.CVToSampleAttenuatorA(value))
}

func (m *Settings) gateDensityString() string {
	return module.GateDensityString(m.ComplexRandom.GateDensity())
}

func (m *Settings) gateDensityValue() units.CV {
	return module.GateDensityToCV(m.ComplexRandom.GateDensity())
}

func (m *Settings) setGateDensityValue(value units.CV) {
	m.ComplexRandom.SetGateDensity(module.CVToGateDensity(value))
}

func (m *Settings) pulseStageDividerString() string {
	return module.PulseStageDividerString(m.ComplexRandom.PulseStageDivider())
}

func (m *Settings) pulseStageDividerValue() units.CV {
	return module.PulseStageDividerToCV(m.ComplexRandom.PulseStageDivider())
}

func (m *Settings) setPulseStageDividerValue(value units.CV) {
	m.ComplexRandom.SetPulseStageDivider(module.CVToPulseStageDivider(value))
}

func (m *Settings) sampleAttenuatorBString() string {
	return module.SampleAttenuatorBString(m.ComplexRandom.SampleAttenuatorB())
}

func (m *Settings) sampleAttenuatorBValue() units.CV {
	return module.SampleAttenuatorBToCV(m.ComplexRandom.SampleAttenuatorB())
}

func (m *Settings) setSampleAttenuatorBValue(value units.CV) {
	m.ComplexRandom.SetSampleAttenuatorB(module.CVToSampleAttenuatorB(value))
}

func (m *Settings) sampleSlewBString() string {
	return module.SampleSlewBString(m.ComplexRandom.SampleSlewB())
}

func (m *Settings) sampleSlewBValue() units.CV {
	return module.SampleSlewBToCV(m.ComplexRandom.SampleSlewB())
}

func (m *Settings) setSampleSlewBValue(value units.CV) {
	m.ComplexRandom.SetSampleSlewB(module.CVToSampleSlewB(value))
}

func (m *Settings) clockSpeedString() string {
	return module.ClockSpeedString(m.ComplexRandom.ClockSpeed())
}

func (m *Settings) clockSpeedValue() units.CV {
	return module.ClockSpeedToCV(m.ComplexRandom.ClockSpeed())
}

func (m *Settings) setClockSpeedValue(value units.CV) {
	m.ComplexRandom.SetClockSpeed(module.CVToClockSpeed(value))
}

func (m *Settings) clockRangeString() string {
	return module.ClockRangeString(m.ComplexRandom.ClockRange())
}

func (m *Settings) clockRangeValue() units.CV {
	return module.ClockRangeToCV(m.ComplexRandom.ClockRange())
}

func (m *Settings) setClockRangeValue(value units.CV) {
	m.ComplexRandom.SetClockRange(module.CVToClockRange(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("sampleAttenuatorA", "Attn.A", m.sampleAttenuatorAString, m.sampleAttenuatorAValue, m.setSampleAttenuatorAValue),
		knobmenu.WithItem("gateDensity", "GDense", m.gateDensityString, m.gateDensityValue, m.setGateDensityValue),
		knobmenu.WithItem("pulseStageDivider", "PSD", m.pulseStageDividerString, m.pulseStageDividerValue, m.setPulseStageDividerValue),
		knobmenu.WithItem("sampleAttenuatorB", "Attn.B", m.sampleAttenuatorBString, m.sampleAttenuatorBValue, m.setSampleAttenuatorBValue),
		knobmenu.WithItem("sampleSlewB", "SlewB", m.sampleSlewBString, m.sampleSlewBValue, m.setSampleSlewBValue),
		knobmenu.WithItem("clockSpeed", "CSpeed", m.clockSpeedString, m.clockSpeedValue, m.setClockSpeedValue),
		knobmenu.WithItem("clockRange", "CRange", m.clockRangeString, m.clockRangeValue, m.setClockRangeValue),
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
