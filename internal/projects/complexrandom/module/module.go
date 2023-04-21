package module

import (
	"math"
	"math/rand"
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type ComplexRandom struct {
	sampleOutA        func(cv units.CV)
	sampleOutB        func(cv units.CV)
	sampleAttenuatorA units.CV
	integrationSlope  units.CV
	gateDensity       units.CV
	pulseStageDivider float32
	sampleAttenuatorB units.CV
	sampleSlewB       units.CV
	clockRange        Clock

	slewLength time.Duration
	slewT      time.Duration
	slewStart  units.CV
	slewEnd    units.CV
	pc         float32
	clock      clock
}

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
	m.slewLength = europim.Lerp(m.sampleSlewB.ToFloat32(), 0, time.Second)
	if m.slewLength < m.slewT {
		m.slewT = 0
	}

	m.SetPulseStageDivider(config.PulseStageDivider)

	return nil
}

func noopSampleOut(cv units.CV) {
}

func (m *ComplexRandom) Gate(high bool) {
	// TODO
}

func (m *ComplexRandom) SetSampleAttenuatorA(cv units.CV) {
	m.sampleAttenuatorA = cv
}

func (m *ComplexRandom) SampleAttenuatorA() units.CV {
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

func (m *ComplexRandom) SetSampleAttenuatorB(cv units.CV) {
	m.sampleAttenuatorB = cv
}

func (m *ComplexRandom) SampleAttenuatorB() units.CV {
	return m.sampleAttenuatorB
}

func (m *ComplexRandom) SetSampleSlewB(cv units.CV) {
	if m.sampleSlewB == cv {
		// no change
		return
	}

	m.sampleSlewB = cv
	m.slewLength = europim.Lerp(m.sampleSlewB.ToFloat32(), 0, time.Second)
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
			t := europim.Clamp(m.slewT+deltaTime, 0, m.slewLength)
			x := float32(t.Seconds() / m.slewLength.Seconds())

			var b units.CV
			if x >= 1 {
				t = 0
				b = m.slewEnd
				m.slewStart = m.slewEnd
			} else {
				b = europim.Clamp(europim.Lerp(x, m.slewStart, m.slewEnd), m.slewStart, m.slewEnd)
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

	ra := units.CV(rand.Float32())
	cva := ra * m.sampleAttenuatorA

	m.sampleOutA(cva)

	rb := units.CV(rand.Float32())
	cvb := rb * m.sampleAttenuatorB

	m.slewEnd = cvb
}
