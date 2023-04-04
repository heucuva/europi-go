package complexarp

import (
	"fmt"
	"strings"
	"time"

	"github.com/heucuva/europi/units"
)

type Module struct {
	out        func(voct units.VOct)
	arpPattern pattern
}

func (m *Module) Init(config Config) error {
	arpKeyboard, err := getArpKeyboard(config)
	if err != nil {
		return err
	}

	arpQuant, err := getArpQuantizer(config)
	if err != nil {
		return err
	}

	switch strings.ToLower(config.ArpPattern) {
	case "brownian":
		m.arpPattern = &patternBrownian{
			keyboard:  arpKeyboard,
			quantizer: arpQuant,
		}
	case "random":
		m.arpPattern = &patternRandom{
			keyboard:  arpKeyboard,
			quantizer: arpQuant,
		}
	default:
		return fmt.Errorf("unsupported arp pattern: %q", config.ArpPattern)
	}
	if err := m.arpPattern.Init(config); err != nil {
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
