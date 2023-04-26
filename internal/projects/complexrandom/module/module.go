package module

import (
	"math"
	"math/rand"
	"time"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

type ComplexRandom struct {
	sampleOutA        func(cv units.BipolarCV)
	sampleOutB        func(cv units.BipolarCV)
	sampleAttenuatorA units.BipolarCV
	integrationSlope  units.CV
	gateDensity       units.CV
	pulseStageDivider float32
	sampleAttenuatorB units.BipolarCV
	sampleSlewB       units.CV
	clockRange        Clock

	slewLength time.Duration
	slewT      time.Duration
	slewStart  units.BipolarCV
	slewEnd    units.BipolarCV
	pc         float32
	clock      clock
}

var slewLerp = lerp.NewLerp32(0, time.Second)

func (m *ComplexRandom) Init(config Config) error {
	fnSampleOutA := config.SampleOutA
	if fnSampleOutA == nil {
		fnSampleOutA = noopSampleOut
	}
	m.sampleOutA = fnSampleOutA

	fnSampleOutB := config.SampleOutB
	if fnSampleOutB == nil {
		fnSampleOutB = noopSampleOut
	}
	m.sampleOutB = fnSampleOutB

	m.sampleAttenuatorA = config.SampleAttenuatorA
	m.integrationSlope = config.IntegrationSlope
	m.gateDensity = config.GateDensity
	m.sampleAttenuatorB = config.SampleAttenuatorB

	var err error
	m.clock, err = getClock(config.ClockRange)
	if err != nil {
		return err
	}
	m.clockRange = config.ClockRange

	m.SetClockSpeed(config.ClockSpeed)

	m.sampleSlewB = config.SampleSlewB
	m.slewLength = slewLerp.Lerp(m.sampleSlewB.ToFloat32())
	if m.slewLength < m.slewT {
		m.slewT = 0
	}

	m.SetPulseStageDivider(config.PulseStageDivider)

	return nil
}

func noopSampleOut(cv units.BipolarCV) {
}

func (m *ComplexRandom) Gate(high bool) {
	// TODO
}

func (m *ComplexRandom) SetSampleAttenuatorA(cv units.BipolarCV) {
	m.sampleAttenuatorA = cv
}

func (m *ComplexRandom) SampleAttenuatorA() units.BipolarCV {
	return m.sampleAttenuatorA
}

func (m *ComplexRandom) SetIntegrationSlope(cv units.CV) {
	m.integrationSlope = cv
}

func (m *ComplexRandom) IntegrationSlope() units.CV {
	return m.integrationSlope
}

func (m *ComplexRandom) SetGateDensity(cv units.CV) {
	m.gateDensity = cv
}

func (m *ComplexRandom) GateDensity() units.CV {
	return m.gateDensity
}

func (m *ComplexRandom) SetPulseStageDivider(psd int) {
	m.pulseStageDivider = float32(psd)
}

func (m *ComplexRandom) PulseStageDivider() int {
	return int(m.pulseStageDivider)
}

func (m *ComplexRandom) SetSampleAttenuatorB(cv units.BipolarCV) {
	m.sampleAttenuatorB = cv
}

func (m *ComplexRandom) SampleAttenuatorB() units.BipolarCV {
	return m.sampleAttenuatorB
}

func (m *ComplexRandom) SetSampleSlewB(cv units.CV) {
	if m.sampleSlewB == cv {
		// no change
		return
	}

	m.sampleSlewB = cv
	m.slewLength = slewLerp.ClampedLerp(m.sampleSlewB.ToFloat32())
	if m.slewLength < m.slewT {
		m.slewT = 0
	}
}

func (m *ComplexRandom) SampleSlewB() units.CV {
	return m.sampleSlewB
}

func (m *ComplexRandom) SetClockSpeed(cv units.CV) {
	m.clock.SetRate(cv)
}

func (m *ComplexRandom) ClockSpeed() units.CV {
	return m.clock.Rate()
}

func (m *ComplexRandom) SetClockRange(mode Clock) {
	if m.clockRange == mode {
		// no change
		return
	}

	speed := m.clock.Rate()

	var err error
	m.clock, err = getClock(mode)
	if err != nil {
		panic(err)
	}

	m.clockRange = mode
	m.clock.SetRate(speed)
}

func (m *ComplexRandom) ClockRange() Clock {
	return m.clockRange
}

func (m *ComplexRandom) Tick(deltaTime time.Duration) {
	triggered := m.clock.Tick(deltaTime)

	if triggered {
		m.processTrigger()
	}

	if m.slewStart != m.slewEnd {
		if m.slewLength == 0 {
			m.slewStart = m.slewEnd
			m.sampleOutB(m.slewStart)
		} else {
			t := clamp.Clamp(m.slewT+deltaTime, 0, m.slewLength)
			x := float32(t.Seconds() / m.slewLength.Seconds())

			var b units.BipolarCV
			if x >= 1 {
				t = 0
				b = m.slewEnd
				m.slewStart = m.slewEnd
			} else {
				b = lerp.NewLerp32(m.slewStart, m.slewEnd).ClampedLerp(x)
			}

			m.slewT = t
			m.sampleOutB(b)
		}
	}
}

func (m *ComplexRandom) processTrigger() {
	if rand.Float32() < m.gateDensity.ToFloat32() {
		return
	}

	pcd := m.pc + 1
	i, f := math.Modf(float64(pcd) / float64(m.pulseStageDivider))
	if i == 0 {
		m.pc = pcd
		return
	}

	m.pc = float32(f)

	ra := units.BipolarCV(rand.Float32()*2.0 - 1.0)
	cva := ra * m.sampleAttenuatorA

	m.sampleOutA(cva)

	rb := units.BipolarCV(rand.Float32()*2.0 - 1.0)
	cvb := rb * m.sampleAttenuatorB

	m.slewEnd = cvb
}
