package module

import (
	"github.com/heucuva/europi/units"
)

type scaleC7 struct{}

func (scaleC7) Keys() []units.VOct {
	return cScaleC_7
}

func (scaleC7) Mode() Scale {
	return ScaleC_7
}

func (scaleC7) Name() string {
	return "C 7"
}

var (
	cScaleC_7 = []units.VOct{keyC, keyE, keyG, keyBFlat}
)
