package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ReleaseModeString(mode FunctionMode) string {
	return mode.String()
}

var releaseModeLerp = lerp.NewLerp32(FunctionModeLinear, FunctionModeQuartic)

func ReleaseModeToCV(mode FunctionMode) units.CV {
	return units.CV(releaseModeLerp.ClampedInverseLerp(mode))
}

func CVToReleaseMode(cv units.CV) FunctionMode {
	return releaseModeLerp.ClampedLerpRound(cv.ToFloat32())
}
