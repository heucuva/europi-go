package module

import (
	"math"
	"math/rand"

	"github.com/heucuva/europi/units"
)

type patternBrownian struct {
	prevKey  units.VOct
	deltaKey units.VOct

	noise     noiseBrownian
	prevNoise float32
}

func (p *patternBrownian) Init(config Config, m *ComplexArp) error {
	p.noise.beta = 0.025
	p.noise.prev = rand.Float32()

	// impossible value, so can't throw out any item
	p.prevNoise = -1.5

	p.UpdateScale(m.scale)

	// generate a random 'key' in range.
	// really just garbage that will get cleaned up by the
	// quantizer phase within the `next` function
	randKey := m.arpPitch + m.arpRange*units.VOct(rand.Float32()*2.0-1.0)
	p.prevKey = p.next(randKey, m)

	return nil
}

func (p *patternBrownian) Next(m *ComplexArp) units.VOct {
	nextKey := p.next(p.prevKey, m)
	p.prevKey = nextKey
	return nextKey
}

func (p *patternBrownian) Pattern() Pattern {
	return PatternBrownian
}

func (p *patternBrownian) next(prevKey units.VOct, m *ComplexArp) units.VOct {
	nextKey := prevKey
	minPitch := m.arpPitch - m.arpRange
	maxPitch := m.arpPitch + m.arpRange
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
		keys := m.scale.Keys()
		nextKey = m.quantizer.QuantizeToValue(float32(v), keys) + units.VOct(oct)
	}

	return nextKey
}

func (p *patternBrownian) UpdateScale(s scale) {
	keys := s.Keys()
	p.deltaKey = units.VOct(1) / units.VOct(len(keys))
}
