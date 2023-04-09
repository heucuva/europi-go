package module

import (
	"math"
	"math/rand"
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type ComplexRandom struct {
	attenA     units.CV
	outA       func(cv units.CV)
	attenB     units.CV
	outB       func(cv units.CV)
	clock      clock
	slewB      units.CV
	slewLength time.Duration
	slewT      time.Duration
	slewStart  units.CV
	slewEnd    units.CV
	gd         units.CV
	pd         float32
	pc         float32
}

const (
	fullSpectrum    = time.Second / 22050
	limitedSpectrum = time.Second * 15 / 22050
)

func noop(_ units.CV) {
}

func (m *ComplexRandom) Init(config Config) error {
	var err error
	m.clock, err = getClock(config.ClockRange)
	if err != nil {
		return err
	}

	m.SetClockRate(config.ClockSpeed)

	m.outA = config.SampleOutA
	if m.outA == nil {
		m.outA = noop
	}

	m.outB = config.SampleOutB
	if m.outB == nil {
		m.outB = noop
	}

	m.attenA = config.SampleAttenuatorA
	m.attenB = config.SampleAttenuatorB

	m.SetSlewB(config.SampleSlewB)

	m.gd = config.GateDensity
	m.pd = config.PulseStageDivider

	return nil
}

func (m *ComplexRandom) SetClockRate(cv units.CV) {
	m.clock.SetRate(cv)
}

func (m *ComplexRandom) ClockRate() units.CV {
	return m.clock.Rate()
}

func (m *ComplexRandom) SetSlewB(cv units.CV) {
	if m.slewB == cv {
		// no change
		return
	}

	m.slewB = cv
	m.slewLength = europim.Lerp(m.slewB.ToFloat32(), 0, time.Second)
	if m.slewLength < m.slewT {
		m.slewT = 0
	}
}

func (m *ComplexRandom) SlewB() units.CV {
	return m.slewB
}

func (m *ComplexRandom) Gate(high bool) {
	// TODO
}

func (m *ComplexRandom) SetSample(cv units.CV) {
	// TODO
}

func (m *ComplexRandom) Tick(deltaTime time.Duration) {
	triggered := m.clock.Tick(deltaTime)

	if triggered {
		m.processTrigger()
	}

	if m.slewStart != m.slewEnd {
		if m.slewLength == 0 {
			m.slewStart = m.slewEnd
			m.outB(m.slewStart)
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
			m.outB(b)
		}
	}
}

func (m *ComplexRandom) processTrigger() {
	if rand.Float32() < m.gd.ToFloat32() {
		return
	}

	pcd := m.pc + 1
	i, f := math.Modf(float64(pcd) / float64(m.pd))
	if i == 0 {
		m.pc = pcd
		return
	}

	m.pc = float32(f)

	ra := units.CV(rand.Float32())
	cva := ra * m.attenA

	m.outA(cva)

	rb := units.CV(rand.Float32())
	cvb := rb * m.attenB

	m.slewEnd = cvb
}
