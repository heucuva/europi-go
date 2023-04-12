package module

import (
	"time"
)

type RandomGates struct {
	channel   [3]channel
	mode      Mode
	startHigh time.Time
}

func noop(_ bool) {
}

func (m *RandomGates) Init(config Config) error {
	for i := range m.channel {
		t := config.Trigger[i]
		if t == nil {
			t = noop
		}
		g := config.Gate[i]
		if g == nil {
			g = noop
		}
		m.channel[i] = channel{
			trig: t,
			gate: g,
		}
	}
	m.mode = config.Mode

	return nil
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

func (m *RandomGates) Tick(deltaTime time.Duration) {
	for i := range m.channel {
		g := &m.channel[i]
		g.Tick(deltaTime)
	}
}

func (m *RandomGates) SetMode(mode Mode) {
	m.mode = mode
}

func (m *RandomGates) Mode() Mode {
	return m.mode
}
