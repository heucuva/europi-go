package module

import (
	"time"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/units"
)

type modelAD struct {
	out       func(cv units.CV)
	attack    functionMode
	attackDur time.Duration
	decay     functionMode
	decayDur  time.Duration
	atten     units.BipolarCV
	state     state
	stateTime time.Duration
}

func (m *modelAD) Trigger() {
	m.state = stateAttack
	m.stateTime = 0
}

func (m *modelAD) SetCV(cv units.BipolarCV) {
	m.atten = cv
}

func (m *modelAD) SetAttack(cv units.CV) {
	m.attackDur = time.Duration(float32(time.Second) * cv.ToFloat32())
}

func (m *modelAD) SetDecay(cv units.CV) {
	m.decayDur = time.Duration(float32(time.Second) * cv.ToFloat32())
}

func (m *modelAD) Tick(deltaTime time.Duration) {
	switch m.state {
	case stateAttack:
		t := m.stateTime + deltaTime
		maxTime := time.Duration(float32(m.attackDur) * float32(m.atten))
		bcv := m.attack.Calc(t, maxTime)
		cv := clamp.Clamp(units.CV(bcv.ToFloat32()), 0.0, 1.0)
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
		bcv := m.attack.Calc(t, m.decayDur)
		cv := clamp.Clamp(units.CV(bcv.ToFloat32()), 0.0, 1.0)
		m.stateTime = t
		m.out(cv)

	default:
	}
}