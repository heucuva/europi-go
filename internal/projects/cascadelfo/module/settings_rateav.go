package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func CVToRateAV(cv units.CV) float32 {
	return cv.ToFloat32()*2.0 - 1.0
}

func RateAVToCV(rateav float32) units.CV {
	return units.CV((rateav + 1.0) / 2.0)
}

func RateAVToString(rateav float32) string {
	return fmt.Sprintf("%+3.1f%%", rateav*100.0)
}
