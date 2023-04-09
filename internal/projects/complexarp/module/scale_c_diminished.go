package module

import (
	"github.com/heucuva/europi/units"
)

type scaleCDiminished struct{}

func (scaleCDiminished) Keys() []units.VOct {
	return cScaleC_Diminished
}

func (scaleCDiminished) Mode() Scale {
	return ScaleC_Diminished
}

func (scaleCDiminished) Name() string {
	return "C dim"
}

var (
	cScaleC_Diminished = []units.VOct{keyC, keyD, keyEFlat, keyF, keyGFlat, keyAFlat, keyA, keyB}
)
