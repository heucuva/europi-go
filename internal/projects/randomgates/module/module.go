package module

import (
	"math/rand"
	"time"
)

type RandomGates struct {
	rem    time.Duration
	gate   [1]func(high bool)
	chance float32
	dur    time.Duration
}

func noop(_ bool) {
}

func (m *RandomGates) Init(config Config) error {
	for i := range m.gate {
		f := config.Gate[i]
		if f == nil {
			f = noop
		}
		m.gate[i] = f
	}
	m.chance = config.Chance
	m.dur = config.Duration
	return nil
}

func (m *RandomGates) Clock(high bool) {
	if m.rem > 0 {
		// disallow updates while in an active gate period
		return
	}

	if !high {
		return
	}

	if rand.Float32() < m.chance {
		m.rem = m.dur
		for _, gate := range m.gate {
			gate(true)
		}
	}
}

func (m *RandomGates) Tick(deltaTime time.Duration) {
	if m.rem > 0 {
		m.rem -= deltaTime
		if m.rem <= 0 {
			for _, gate := range m.gate {
				gate(false)
			}
		}
	}
}
