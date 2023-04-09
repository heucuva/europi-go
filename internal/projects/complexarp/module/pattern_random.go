package module

import (
	"math"

	"github.com/heucuva/europi/units"
)

type patternRandom struct {
	patRange units.VOct
	patPitch units.VOct

	scale scale

	noise     noiseWhite
	prevNoise float32

	quantizer quantizer[units.VOct]
}

func (p *patternRandom) Init(config Config) error {
	p.patRange = config.ArpRange
	p.patPitch = config.ArpPitch

	// impossible value, so can't throw out any item
	p.prevNoise = -1.5

	return nil
}

func (p *patternRandom) Next() units.VOct {
	curNoise := p.prevNoise
	for curNoise == p.prevNoise {
		curNoise = p.noise.Get()
	}
	p.prevNoise = curNoise

	voct := p.patRange*units.VOct(curNoise) + p.patPitch
	oct, v := math.Modf(float64(voct.ToFloat32()))

	keys := p.scale.Keys()

	return p.quantizer.QuantizeToValue(float32(v), keys) + units.VOct(oct)
}

func (p *patternRandom) SetArpPitch(voct units.VOct) {
	p.patPitch = voct
}

func (p *patternRandom) SetArpRange(voct units.VOct) {
	p.patRange = voct
}

func (p *patternRandom) SetScale(s scale) {
	p.scale = s
}

func (p *patternRandom) Scale() Scale {
	return p.scale.Mode()
}

func (p *patternRandom) ScaleName() string {
	return p.scale.Name()
}

func (p *patternRandom) ArpRange() units.VOct {
	return p.patRange
}
