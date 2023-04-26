package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ClockRangeString(mode Clock) string {
	return mode.String()
}

var clockRangeLerp = lerp.NewLerp32(ClockFull, ClockLimited)

func ClockRangeToCV(mode Clock) units.CV {
	return units.CV(clockRangeLerp.ClampedInverseLerp(mode))
}

func CVToClockRange(cv units.CV) Clock {
	return clockRangeLerp.ClampedLerpRound(cv.ToFloat32())
}
