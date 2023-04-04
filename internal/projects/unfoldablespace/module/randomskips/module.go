package randomskips

import (
	"math/rand"
	"time"

	"github.com/heucuva/europi/units"
)

type Module struct {
	gate   [1]gate
	chance float32
}

func (m *Module) Init(config Config) error {
	m.chance = config.Chance
	for i := range m.gate {
		f := config.Gate[i]
		if f == nil {
			f = func(high bool) {}
		}
		m.gate[i].out = f
		m.SetCV(i, 1)
	}
	return nil
}

func (m *Module) Gate(gate int, high bool) {
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

func (m *Module) SetCV(gate int, cv units.CV) {
	if gate < 0 || gate > len(m.gate) {
		panic("gate: out of range")
	}

	g := &m.gate[gate]
	g.chance = m.chance * float32(cv)
}

func (m *Module) Tick(deltaTime time.Duration) {

}
