package module

import (
	"math"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func RateString(freq units.Hertz) string {
	return freq.String()
}

var rateLerp = lerp.NewLerp32[float32](-3.5, 7.5)

func RateToCV(freq units.Hertz) units.CV {
	exp := float32(math.Log2(float64(freq)))
	return units.CV(rateLerp.ClampedInverseLerp(exp))
}

func CVToRate(cv units.CV) units.Hertz {
	pow := rateLerp.ClampedLerp(cv.ToFloat32())
	return units.Hertz(math.Pow(2.0, float64(pow)))
}

func AdjustRate(rate units.Hertz, cv units.BipolarCV, atten units.BipolarCV) units.Hertz {
	modifiedRate := RateToCV(rate).ToBipolarCV(1) + cv*atten
	v, _ := modifiedRate.ToCV()
	return CVToRate(v)
}
