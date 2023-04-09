package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

type Pattern int

const (
	PatternBrownian = Pattern(iota)
	PatternRandom
)

type pattern interface {
	Init(config Config, m *ComplexArp) error
	Next(m *ComplexArp) units.VOct
	UpdateScale(s scale)
}

func (m *ComplexArp) setArpPattern(config Config) error {
	var arpPattern pattern
	switch config.ArpPattern {
	case PatternBrownian:
		arpPattern = &patternBrownian{}
	case PatternRandom:
		arpPattern = &patternRandom{}
	default:
		return fmt.Errorf("unsupported arp pattern: %d", config.ArpPattern)
	}

	if err := arpPattern.Init(config, m); err != nil {
		return err
	}

	m.arpPattern = arpPattern

	return nil
}
