package clockgenerator

import (
	"fmt"
	"time"
)

type Module struct {
	interval time.Duration
	enabled  bool
	out      func()
	t        time.Duration
}

func noop() {
}

func (m *Module) Init(config Config) error {
	if config.BPM <= 0 {
		return fmt.Errorf("invalid bpm setting: %v", config.BPM)
	}

	m.interval = time.Duration(float32(time.Minute) / config.BPM)
	m.enabled = config.Enabled
	m.out = config.ClockOut
	if m.out == nil {
		m.out = noop
	}
	return nil
}

func (m *Module) Toggle() {
	m.enabled = !m.enabled
	m.t = 0
}

func (m *Module) SetEnabled(enabled bool) {
	m.enabled = enabled
	m.t = 0
}

func (m *Module) Enabled() bool {
	return m.enabled
}

func (m *Module) Tick(deltaTime time.Duration) {
	if !m.enabled {
		return
	}

	t := m.t + deltaTime
	m.t = t % m.interval
	if t >= m.interval {
		m.out()
	}
}
