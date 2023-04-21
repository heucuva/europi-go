package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCAugmented struct{}

func (scaleCAugmented) Keys() []units.VOct {
	return cScaleC_Augmented
}

func (scaleCAugmented) Mode() Scale {
	return ScaleC_Augmented
}

var (
	cScaleC_Augmented = []units.VOct{keyC, keyDSharp, keyE, keyG, keyAFlat, keyB}
)
