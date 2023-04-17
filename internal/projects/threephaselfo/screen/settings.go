package screen

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/threephaselfo/module"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km  *knobmenu.KnobMenu
	LFO *module.ThreePhaseLFO
}

func (m *Settings) waveModeString() string {
	return module.WaveModeString(m.LFO.WaveMode())
}

func (m *Settings) waveModeValue() units.CV {
	return module.WaveModeToCV(m.LFO.WaveMode())
}

func (m *Settings) setWaveModeValue(value units.CV) {
	m.LFO.SetWaveMode(module.CVToWaveMode(value))
}

func (m *Settings) phi3RateString() string {
	return module.Phi3RateString(m.LFO.Phi3Rate())
}

func (m *Settings) phi3RateValue() units.CV {
	return module.Phi3RateToCV(m.LFO.Phi3Rate())
}

func (m *Settings) setPhi3RateValue(value units.CV) {
	m.LFO.SetPhi3Rate(module.CVToPhi3Rate(value))
}

func (m *Settings) skewRateString() string {
	return module.SkewRateString(m.LFO.SkewRate())
}

func (m *Settings) skewRateValue() units.CV {
	return module.SkewRateToCV(m.LFO.SkewRate())
}

func (m *Settings) setSkewRateValue(value units.CV) {
	m.LFO.SetSkewRate(module.CVToSkewRate(value))
}

func (m *Settings) skewShapeString() string {
	return module.SkewShapeString(m.LFO.SkewShape())
}

func (m *Settings) skewShapeValue() units.CV {
	return module.SkewShapeToCV(m.LFO.SkewShape())
}

func (m *Settings) setSkewShapeValue(value units.CV) {
	m.LFO.SetSkewShape(module.CVToSkewShape(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("waveMode", "Wave", m.waveModeString, m.waveModeValue, m.setWaveModeValue),
		knobmenu.WithItem("phi3Rate", "Rate", m.phi3RateString, m.phi3RateValue, m.setPhi3RateValue),
		knobmenu.WithItem("skewRate", "SRate", m.skewRateString, m.skewRateValue, m.setSkewRateValue),
		knobmenu.WithItem("skewShape", "SShape", m.skewShapeString, m.skewShapeValue, m.setSkewShapeValue),
	)
	if err != nil {
		panic(err)
	}

	m.km = km
}

func (m *Settings) Button1Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Settings) Button1(e *europi.EuroPi, p machine.Pin) {
	m.km.Next()
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
}
