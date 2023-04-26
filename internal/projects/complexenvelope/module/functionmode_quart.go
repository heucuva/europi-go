package module

import (
	"math"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

type functionModeQuartic struct{}

func (functionModeQuartic) Calc(t, dur time.Duration) units.BipolarCV {
	maxT := float32(dur.Seconds())
	if maxT == 0 {
		return 0.0
	}
	curT := float32(t.Seconds())
	return units.BipolarCV(math.Pow(float64(curT/maxT), 3.32))
}
