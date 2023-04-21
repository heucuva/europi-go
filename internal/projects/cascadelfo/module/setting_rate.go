package module

import (
	"math"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func RateString(freq units.Hertz) string {
	return freq.String()
}

func RateToCV(freq units.Hertz) units.CV {
	exp := float32(math.Log2(float64(freq)))
	return units.CV(europim.InverseLerp(exp, -3.5, 7.5))
}

func CVToRate(cv units.CV) units.Hertz {
	pow := europim.Lerp[float32](cv.ToFloat32(), -3.5, 7.5)
	return units.Hertz(math.Pow(2.0, float64(pow)))
}

func AdjustRate(rate units.Hertz, cv units.BipolarCV, atten units.BipolarCV) units.Hertz {
	modifiedRate := RateToCV(rate).ToBipolarCV() + cv*atten
	return CVToRate(modifiedRate.ToCV())
}
