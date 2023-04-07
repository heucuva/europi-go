package complexenvelope

import (
	"time"

	"github.com/heucuva/europi/units"
)

type functionModeLinear struct{}

func (functionModeLinear) Calc(t, dur time.Duration) units.CV {
	maxT := float32(dur.Seconds())
	if maxT == 0 {
		return 0.0
	}
	curT := float32(t.Seconds())
	return units.CV(curT / maxT)
}
