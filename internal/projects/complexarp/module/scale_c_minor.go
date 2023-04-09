package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCMinor struct{}

func (scaleCMinor) Keys() []units.VOct {
	return cScaleC_Minor
}

func (scaleCMinor) Mode() Scale {
	return ScaleC_Minor
}

func (scaleCMinor) Name() string {
	return "C min"
}

var (
	cScaleC_Minor = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
)
