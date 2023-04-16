package module

import (
	"time"
)

type RandomGates struct {
	channel   [3]channel
	mode      Mode
	startHigh time.Time
}

func (m *RandomGates) Init(config Config) error {
	for i, fn := range config.Trigger {
		if fn == nil {
			fn = noop
		}
		m.channel[i].trig = fn
	}

	for i, fn := range config.Gate {
		if fn == nil {
			fn = noop
		}
		m.channel[i].gate = fn
	}
	m.mode = config.Mode

	return nil
}

func noop(high bool) {
}

func (m *RandomGates) Clock(high bool) {
	if high {
		m.startHigh = time.Now()
		return
	}

	clockDur := time.Since(m.startHigh)
	for i := range m.channel {
		g := &m.channel[i]
		g.Start(clockDur, m.mode)
	}
}

func (m *RandomGates) SetMode(mode Mode) {
	m.mode = mode
}

func (m *RandomGates) Mode() Mode {
	return m.mode
}

func (m *RandomGates) Tick(deltaTime time.Duration) {
	for i := range m.channel {
		g := &m.channel[i]
		g.Tick(deltaTime)
	}
}
