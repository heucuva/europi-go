package complexenvelope

import (
	"time"

	"github.com/heucuva/europi/units"
)

type Module struct {
	env [2]envelope
}

func (m *Module) Init(config Config) error {
	for i := range m.env {
		e := &m.env[i]
		if err := e.Init(config.Env[i]); err != nil {
			return err
		}
	}
	return nil
}

func (m *Module) SetCV(env int, cv units.CV) {
	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].SetCV(cv)
}

func (m *Module) Gate(env int, high bool) {
	if !high {
		return
	}

	if env < 0 || env >= len(m.env) {
		panic("env: out of range")
	}
	m.env[env].Trigger()
}

func (m *Module) Tick(deltaTime time.Duration) {
	for i := range m.env {
		e := &m.env[i]
		e.Tick(deltaTime)
	}
}
