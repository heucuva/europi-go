package complexarp

import (
	"time"

	"github.com/heucuva/europi/units"
)

type Module struct {
	out        func(voct units.VOct)
	arpPattern pattern
}

func (m *Module) Init(config Config) error {
	var err error
	m.arpPattern, err = getArpPattern(config)
	if err != nil {
		return err
	}

	m.out = config.ArpOut

	return nil
}

func (m *Module) ArpClock(high bool) {
	if high {
		voct := m.arpPattern.Next()
		m.out(voct)
	}
}

func (m *Module) Tick(deltaTime time.Duration) {
}
