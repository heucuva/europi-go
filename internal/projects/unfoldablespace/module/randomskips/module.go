package randomskips

import (
	"math/rand"
	"time"

	"github.com/heucuva/europi/units"
)

type RandomSkips struct {
	gate   [1]gate
	chance float32
}

func noop(_ bool) {
}

func (m *RandomSkips) Init(config Config) error {
	m.chance = config.Chance
	for i := range m.gate {
		f := config.Gate[i]
		if f == nil {
			f = noop
		}
		m.gate[i].out = f
		m.SetCV(i, 1)
	}
	return nil
}

func (m *RandomSkips) Gate(gate int, high bool) {
	if gate < 0 || gate > len(m.gate) {
		panic("gate: out of range")
	}

	g := &m.gate[gate]
	prev := g.enabled
	next := prev

	switch high {
	case true:
		if rand.Float32() < g.chance {
			next = true
		}
	case false:
		next = false
	}

	if prev != next {
		g.enabled = next
		g.out(next)
	}
}

func (m *RandomSkips) SetCV(gate int, cv units.CV) {
	if gate < 0 || gate > len(m.gate) {
		panic("gate: out of range")
	}

	g := &m.gate[gate]
	g.chance = m.chance * float32(cv)
}

func (m *RandomSkips) Tick(deltaTime time.Duration) {

}
