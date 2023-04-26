package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleCMinor struct{}

func (scaleCMinor) Keys() []units.VOct {
	return cScaleC_Minor
}

func (scaleCMinor) Mode() Scale {
	return ScaleC_Minor
}

var (
	cScaleC_Minor = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
)
