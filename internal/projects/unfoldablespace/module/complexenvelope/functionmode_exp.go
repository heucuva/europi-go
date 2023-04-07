package complexenvelope

import (
	"math"
	"time"

	"github.com/heucuva/europi/units"
)

type functionModeExponential struct{}

func (functionModeExponential) Calc(t, dur time.Duration) units.CV {
	maxT := float32(dur.Seconds())
	if maxT == 0 {
		return 0.0
	}
	curT := float32(t.Seconds())
	return units.CV(math.Exp(-4.0 * float64(curT/maxT)))
}
