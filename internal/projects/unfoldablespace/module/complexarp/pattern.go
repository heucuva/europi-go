package complexarp

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
	Init(config Config) error
	Next() units.VOct
}

func getArpPattern(config Config) (pattern, error) {
	arpScale, err := getScale(config.Scale)
	if err != nil {
		return nil, err
	}

	arpQuant, err := getArpQuantizer(config.Quantizer)
	if err != nil {
		return nil, err
	}

	var arpPattern pattern
	switch config.ArpPattern {
	case PatternBrownian:
		arpPattern = &patternBrownian{
			scale:     arpScale,
			quantizer: arpQuant,
		}
	case PatternRandom:
		arpPattern = &patternRandom{
			scale:     arpScale,
			quantizer: arpQuant,
		}
	default:
		return nil, fmt.Errorf("unsupported arp pattern: %d", config.ArpPattern)
	}

	if err := arpPattern.Init(config); err != nil {
		return nil, err
	}

	return arpPattern, nil
}
