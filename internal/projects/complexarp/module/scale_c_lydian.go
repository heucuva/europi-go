package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCLydian struct{}

func (scaleCLydian) Keys() []units.VOct {
	return cScaleC_Lydian
}

func (scaleCLydian) Mode() Scale {
	return ScaleC_Lydian
}

var (
	cScaleC_Lydian = []units.VOct{keyC, keyD, keyE, keyFSharp, keyG, keyA, keyB}
)
