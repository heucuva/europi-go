package complexarp

import (
	"math"
	"math/rand"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type patternBrownian struct {
	patRange units.VOct
	patPitch units.VOct

	keyboard *keyboard
	prevKey  units.VOct
	deltaKey units.VOct

	noise     noiseBrownian
	prevNoise float32

	quantizer quantizer[units.VOct]
}

func (p *patternBrownian) Init(config Config) error {
	p.patRange = config.ArpRange
	p.patPitch = config.ArpPitch

	p.noise.beta = 0.025
	p.noise.prev = rand.Float32()

	// impossible value, so can't throw out any item
	p.prevNoise = -1.5

	p.deltaKey = units.VOct(1) / units.VOct(len(p.keyboard.keys))

	return nil
}

func (p *patternBrownian) Next() units.VOct {
	nextKey := p.prevKey
	for nextKey == p.prevKey {
		curNoise := p.noise.Get()
		up := curNoise >= p.prevNoise
		p.prevNoise = curNoise

		if up {
			nextKey += p.deltaKey
		} else {
			nextKey -= p.deltaKey
		}

		halfRange := p.patRange / 2.0
		nextKey = europim.Clamp(nextKey, -halfRange, halfRange)
	}
	p.prevKey = nextKey

	voct := nextKey + p.patPitch
	oct, v := math.Modf(float64(voct.ToFloat32()))

	return p.quantizer.QuantizeToValue(float32(v), p.keyboard.keys) + units.VOct(oct)
}
