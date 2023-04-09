package module

import (
	"math/rand"
	"time"
)

type RandomGates struct {
	gate   [1]gate
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
		m.gate[i] = gate{
			out: f,
		}
	}
	m.chance = config.Chance
	m.dur = config.Duration
	return nil
}

func (m *RandomGates) SetChance(chance float32) {
	m.chance = chance
}

func (m *RandomGates) Chance() float32 {
	return m.chance
}

func (m *RandomGates) Clock(high bool) {
	for i := range m.gate {
		g := &m.gate[i]

		if g.rem > 0 || rand.Float32() >= m.chance {
			// disallow updates while in an active gate period or we fail our chance
			continue
		}

		g.level = !g.level
		if g.level {
			g.rem = m.dur
		}
		g.out(g.level)
	}
}

func (m *RandomGates) Tick(deltaTime time.Duration) {
	for i := range m.gate {
		g := &m.gate[i]

		if g.rem > 0 && g.level {
			g.rem -= deltaTime
			if g.rem <= 0 {
				g.level = false
				g.rem = 0
				g.out(g.level)
			}
		}
	}
}
