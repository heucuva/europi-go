package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleCDorian struct{}

func (scaleCDorian) Keys() []units.VOct {
	return cScaleC_Dorian
}

func (scaleCDorian) Mode() Scale {
	return ScaleC_Dorian
}

var (
	cScaleC_Dorian = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
)
