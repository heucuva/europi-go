package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleCMajor struct{}

func (scaleCMajor) Keys() []units.VOct {
	return cScaleC_Major
}

func (scaleCMajor) Mode() Scale {
	return ScaleC_Major
}

var (
	cScaleC_Major = []units.VOct{keyC, keyD, keyE, keyF, keyG, keyA, keyB}
)
