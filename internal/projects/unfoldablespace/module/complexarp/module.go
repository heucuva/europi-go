package complexarp

import (
	"time"

	"github.com/heucuva/europi/units"
)

type ComplexArp struct {
	out        func(voct units.VOct)
	arpPattern pattern
}

func (m *ComplexArp) Init(config Config) error {
	var err error
	m.arpPattern, err = getArpPattern(config)
	if err != nil {
		return err
	}

	m.out = config.ArpOut

	return nil
}

func (m *ComplexArp) ArpClock(high bool) {
	if high {
		voct := m.arpPattern.Next()
		m.out(voct)
	}
}

func (m *ComplexArp) Tick(deltaTime time.Duration) {
}
