package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ClockRangeString(mode Clock) string {
	return mode.String()
}

func ClockRangeToCV(mode Clock) units.CV {
	return units.CV(europim.InverseLerp(mode, ClockFull, ClockLimited))
}

func CVToClockRange(cv units.CV) Clock {
	return europim.LerpRound(cv.ToFloat32(), ClockFull, ClockLimited)
}
