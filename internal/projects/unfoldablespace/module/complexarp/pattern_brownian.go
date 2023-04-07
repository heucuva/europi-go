package complexarp

import (
	"math"
	"math/rand"

	"github.com/heucuva/europi/units"
)

type patternBrownian struct {
	patRange units.VOct
	patPitch units.VOct

	scale    scale
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

	p.deltaKey = units.VOct(1) / units.VOct(len(p.scale))

	// generate a random 'key' in range.
	// really just garbage that will get cleaned up by the
	// quantizer phase within the `next` function
	randKey := p.patPitch + p.patRange*units.VOct(rand.Float32()-0.5)
	p.prevKey = p.next(randKey)

	return nil
}

func (p *patternBrownian) Next() units.VOct {
	nextKey := p.next(p.prevKey)
	p.prevKey = nextKey
	return nextKey
}

func (p *patternBrownian) next(prevKey units.VOct) units.VOct {
	nextKey := prevKey
	halfRange := p.patRange / 2.0
	minPitch := p.patPitch - halfRange
	maxPitch := p.patPitch + halfRange
	for nextKey == prevKey {
		curNoise := p.noise.Get()
		up := curNoise >= p.prevNoise
		p.prevNoise = curNoise

		if up {
			nextKey += p.deltaKey
		} else {
			nextKey -= p.deltaKey
		}

		voct := nextKey
		// loop the pitch around a ring of the scale
		if voct >= maxPitch {
			voct = minPitch
		} else if voct <= minPitch {
			voct = maxPitch
		}

		oct, v := math.Modf(float64(voct.ToFloat32()))
		nextKey = p.quantizer.QuantizeToValue(float32(v), p.scale) + units.VOct(oct)
	}

	return nextKey
}
