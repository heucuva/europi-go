package complexrandom

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type Module struct {
	attenA        units.CV
	outA          func(cv units.CV)
	attenB        units.CV
	outB          func(cv units.CV)
	holdDuration  time.Duration
	clockInterval time.Duration
	slewLength    time.Duration
	slewT         time.Duration
	slewStart     units.CV
	slewEnd       units.CV
	gd            units.CV
	pd            float32
	pc            float32
}

const (
	fullSpectrum    = time.Second / 22050
	limitedSpectrum = time.Second * 15 / 22050
)

func (m *Module) Init(config Config) error {
	switch strings.ToLower(config.ClockRange) {
	case "full":
		m.clockInterval = fullSpectrum
	case "limited":
		m.clockInterval = limitedSpectrum
	default:
		return fmt.Errorf("unsupported clock range: %q", config.ClockRange)
	}

	m.outA = config.SampleOutA
	if m.outA == nil {
		m.outA = func(cv units.CV) {}
	}

	m.outB = config.SampleOutB
	if m.outB == nil {
		m.outB = func(cv units.CV) {}
	}

	m.attenA = config.SampleAttenuatorA
	m.attenB = config.SampleAttenuatorB

	m.slewLength = europim.Lerp(config.SampleSlewB.ToFloat32(), 0, time.Second)
	m.slewT = 0

	m.gd = config.GateDensity
	m.pd = config.PulseStageDivider

	return nil
}

func (m *Module) Tick(deltaTime time.Duration) {
	t := m.holdDuration + deltaTime
	triggered := t > m.clockInterval
	m.holdDuration = t % m.clockInterval

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
			if x >= 1 {
				t = 0
			}
			m.slewT = t
			b := europim.Clamp(europim.Lerp(0.00144151*(x*x)+0.209508*x-0.20705, m.slewStart, m.slewEnd), m.slewStart, m.slewEnd)
			m.outB(b)
			if m.slewT >= m.slewLength {
				m.slewStart = m.slewEnd
			}
		}
	}
}

func (m *Module) processTrigger() {
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
