package complexenvelope

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type modelAD struct {
	out       func(cv units.CV)
	attack    functionMode
	attackDur time.Duration
	decay     functionMode
	decayDur  time.Duration
	atten     units.CV
	state     state
	stateTime time.Duration
}

func (m *modelAD) Trigger() {
	m.state = stateAttack
	m.stateTime = 0
}

func (m *modelAD) SetCV(cv units.CV) {
	m.atten = cv
}

func (m *modelAD) Tick(deltaTime time.Duration) {
	switch m.state {
	case stateAttack:
		t := m.stateTime + deltaTime
		maxTime := time.Duration(float32(m.attackDur) * float32(m.atten))
		cv := europim.Clamp(m.attack.Calc(t, maxTime), 0.0, 1.0)
		if t >= maxTime {
			m.state = stateDecay
			t = 0
		}
		m.stateTime = t
		m.out(cv)

	case stateDecay:
		t := m.stateTime + deltaTime
		if t >= m.decayDur {
			m.state = stateIdle
			m.stateTime = 0
			m.out(0)
			return
		}
		cv := europim.Clamp(m.decay.Calc(t, m.decayDur), 0.0, 1.0)
		m.stateTime = t
		m.out(cv)

	default:
	}
}
