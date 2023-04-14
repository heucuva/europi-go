package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCPhrygian struct{}

func (scaleCPhrygian) Keys() []units.VOct {
	return cScaleC_Phrygian
}

func (scaleCPhrygian) Mode() Scale {
	return ScaleC_Phrygian
}

var (
	cScaleC_Phrygian = []units.VOct{keyC, keyDFlat, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
)
