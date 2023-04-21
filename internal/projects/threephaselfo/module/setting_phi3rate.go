package module

import (
	"math"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	MinPhi3Rate units.Hertz = 1.0 / 65536.0 // 2**-16
	MaxPhi3Rate units.Hertz = 1.0 * 65536.0 // 2**16
)

func Phi3RateString(freq units.Hertz) string {
	return freq.String()
}

func Phi3RateToCV(freq units.Hertz) units.CV {
	rr := europim.Clamp(freq, MinPhi3Rate, MaxPhi3Rate)
	exp := float32(math.Log2(float64(rr)))
	return units.CV(europim.Clamp((exp+16.0)/32.0, 0.0, 1.0))
}

func CVToPhi3Rate(cv units.CV) units.Hertz {
	pow := cv.ToFloat32()*32.0 - 16.0
	return units.Hertz(math.Pow(2.0, float64(pow)))
}
