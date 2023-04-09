package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCMajor struct{}

func (scaleCMajor) Keys() []units.VOct {
	return cScaleC_Major
}

func (scaleCMajor) Mode() Scale {
	return ScaleC_Major
}

func (scaleCMajor) Name() string {
	return "C maj"
}

var (
	cScaleC_Major = []units.VOct{keyC, keyD, keyE, keyF, keyG, keyA, keyB}
)
