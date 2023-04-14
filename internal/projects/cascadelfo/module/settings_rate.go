package module

import (
	"fmt"
	"math"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToRate(cv units.CV) float32 {
	pow := europim.Lerp[float32](cv.ToFloat32(), -3.5, 7.5)
	return float32(math.Pow(2.0, float64(pow)))
}

func RateToCV(rate float32) units.CV {
	exp := float32(math.Log2(float64(rate)))
	return units.CV(europim.InverseLerp(exp, -3.5, 7.5))
}

func RateToString(rate float32) string {
	switch {
	case rate < 0.001:
		return fmt.Sprintf("%3.1fuHz", rate*1000000.0)
	case rate < 1:
		return fmt.Sprintf("%3.1fmHz", rate*1000.0)
	case rate >= 1000:
		return fmt.Sprintf("%3.1fkHz", rate/1000.0)
	default:
		return fmt.Sprintf("%5.1fHz", rate)
	}
}
