package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCHarmonic5 struct{}

func (scaleCHarmonic5) Keys() []units.VOct {
	return cScaleC_Harmonic_5
}

func (scaleCHarmonic5) Mode() Scale {
	return ScaleC_Harmonic_5
}

func (scaleCHarmonic5) Name() string {
	return "C hm5"
}

var (
	cScaleC_Harmonic_5 = []units.VOct{keyC, keyDFlat, keyE, keyF, keyG, keyAFlat, keyBFlat}
)
