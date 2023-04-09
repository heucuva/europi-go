package module

import (
	"math"

	"github.com/heucuva/europi/units"
)

type patternRandom struct {
	noise     noiseWhite
	prevNoise float32
}

func (p *patternRandom) Init(config Config, m *ComplexArp) error {
	// impossible value, so can't throw out any item
	p.prevNoise = -1.5

	return nil
}

func (p *patternRandom) Next(m *ComplexArp) units.VOct {
	curNoise := p.prevNoise
	for curNoise == p.prevNoise {
		curNoise = p.noise.Get()
	}
	p.prevNoise = curNoise

	voct := m.patRange*units.VOct(curNoise) + m.patPitch
	oct, v := math.Modf(float64(voct.ToFloat32()))

	keys := m.scale.Keys()

	return m.quantizer.QuantizeToValue(float32(v), keys) + units.VOct(oct)
}

func (p *patternRandom) UpdateScale(s scale) {
}
