package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleCHarmonic5 struct{}

func (scaleCHarmonic5) Keys() []units.VOct {
	return cScaleC_Harmonic_5
}

func (scaleCHarmonic5) Mode() Scale {
	return ScaleC_Harmonic_5
}

var (
	cScaleC_Harmonic_5 = []units.VOct{keyC, keyDFlat, keyE, keyF, keyG, keyAFlat, keyBFlat}
)
