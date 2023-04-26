package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleC7 struct{}

func (scaleC7) Keys() []units.VOct {
	return cScaleC_7
}

func (scaleC7) Mode() Scale {
	return ScaleC_7
}

var (
	cScaleC_7 = []units.VOct{keyC, keyE, keyG, keyBFlat}
)
