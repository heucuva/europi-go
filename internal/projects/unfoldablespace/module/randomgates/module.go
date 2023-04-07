package randomgates

import (
	"math/rand"
	"time"
)

type Module struct {
	rem  time.Duration
	gate [1]func(high bool)
	dur  time.Duration
}

func noop(_ bool) {
}

func (m *Module) Init(config Config) error {
	for i := range m.gate {
		f := config.Gate[i]
		if f == nil {
			f = noop
		}
		m.gate[i] = f
	}
	m.dur = config.Duration
	return nil
}

const oneThird = 1.0 / 3.0

func (m *Module) Clock() {
	if m.rem > 0 {
		// disallow updates while in an active gate period
		return
	}

	if rand.Float32() < oneThird {
		m.rem = m.dur
		for _, gate := range m.gate {
			gate(true)
		}
	}
}

func (m *Module) Tick(deltaTime time.Duration) {
	if m.rem > 0 {
		m.rem -= deltaTime
		if m.rem <= 0 {
			for _, gate := range m.gate {
				gate(false)
			}
		}
	}
}
