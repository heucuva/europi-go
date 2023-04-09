package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCSuspended struct{}

func (scaleCSuspended) Keys() []units.VOct {
	return cScaleC_Suspended
}

func (scaleCSuspended) Mode() Scale {
	return ScaleC_Suspended
}

func (scaleCSuspended) Name() string {
	return "C sus"
}

var (
	cScaleC_Suspended = []units.VOct{keyC, keyD, keyF, keyG, keyBFlat}
)
