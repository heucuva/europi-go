package module

import (
	"github.com/awonak/EuroPiGo/units"
)

type scaleCDiminished struct{}

func (scaleCDiminished) Keys() []units.VOct {
	return cScaleC_Diminished
}

func (scaleCDiminished) Mode() Scale {
	return ScaleC_Diminished
}

var (
	cScaleC_Diminished = []units.VOct{keyC, keyD, keyEFlat, keyF, keyGFlat, keyAFlat, keyA, keyB}
)
