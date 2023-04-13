package module

import (
	"math/rand"
	"time"

	"github.com/heucuva/europi/units"
)

type RandomSkips struct {
	out       func(high bool)
	active    bool
	lastInput bool
	chance    float32
	cv        float32
	ac        float32 // attenuated chance (cv * chance)
}

func noop(_ bool) {
}

func (m *RandomSkips) Init(config Config) error {
	m.chance = config.Chance
	f := config.Gate
	if f == nil {
		f = noop
	}
	m.out = f
	m.SetCV(1)
	return nil
}

func (m *RandomSkips) Gate(high bool) {
	prev := m.active
	lastInput := m.lastInput
	next := prev
	m.lastInput = high

	if high != lastInput && rand.Float32() < m.ac {
		next = !prev
	}

	if prev != next {
		m.active = next
		m.out(next)
	}
}

func (m *RandomSkips) SetCV(cv units.CV) {
	m.cv = cv.ToFloat32()
	m.ac = m.chance * m.cv
}

func (m *RandomSkips) SetChance(chance float32) {
	m.chance = chance
}

func (m *RandomSkips) Chance() float32 {
	return m.chance
}

func (m *RandomSkips) Tick(deltaTime time.Duration) {

}
