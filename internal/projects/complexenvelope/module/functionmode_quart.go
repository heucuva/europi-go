package module

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type functionModeQuartic struct{}

func (functionModeQuartic) Calc(t, dur time.Duration) units.CV {
	maxT := float32(dur.Seconds())
	if maxT == 0 {
		return 0.0
	}
	curT := float32(t.Seconds())
	return units.CV(math.Pow(float64(curT/maxT), 3.32))
}
