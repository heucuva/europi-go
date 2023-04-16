package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type ComplexEnvelope struct {
	env [2]envelope
}

func (m *ComplexEnvelope) Init(config Config) error {
	for i := range config.Env {
		e := &m.env[i]
		if err := e.Init(config.Env[i]); err != nil {
			return err
		}
	}
	return nil
}

func (m *ComplexEnvelope) Gate(env int, high bool) {
	if !high {
		return
	}

	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].Trigger()
}

func (m *ComplexEnvelope) SetMode(env int, mode EnvelopeMode) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}

	e := &m.env[env]
	cfg := e.config
	cfg.Mode = mode
	if err := e.Init(cfg); err != nil {
		panic(err)
	}
}

func (m *ComplexEnvelope) Mode(env int) EnvelopeMode {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	return m.env[env].config.Mode
}

func (m *ComplexEnvelope) SetAttackMode(env int, mode FunctionMode) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	e := &m.env[env]
	cfg := e.config
	cfg.AttackMode = mode
	if err := e.Init(cfg); err != nil {
		panic(err)
	}
}

func (m *ComplexEnvelope) AttackMode(env int) FunctionMode {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	return m.env[env].config.AttackMode
}

func (m *ComplexEnvelope) SetReleaseMode(env int, mode FunctionMode) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	e := &m.env[env]
	cfg := e.config
	cfg.ReleaseMode = mode
	if err := e.Init(cfg); err != nil {
		panic(err)
	}
}

func (m *ComplexEnvelope) ReleaseMode(env int) FunctionMode {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	return m.env[env].config.ReleaseMode
}

func (m *ComplexEnvelope) SetAttack(env int, cv units.CV) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].SetAttack(cv)
}

func (m *ComplexEnvelope) Attack(env int) units.CV {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	return m.env[env].config.Attack
}

func (m *ComplexEnvelope) SetDecay(env int, cv units.CV) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].SetDecay(cv)
}

func (m *ComplexEnvelope) Decay(env int) units.CV {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	return m.env[env].config.Decay
}

func (m *ComplexEnvelope) SetCV(env int, cv units.BipolarCV) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].SetCV(cv)
}

func (m *ComplexEnvelope) Tick(deltaTime time.Duration) {
	for i := range m.env {
		e := &m.env[i]
		e.Tick(deltaTime)
	}
}
