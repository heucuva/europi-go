package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ReleaseModeString(mode FunctionMode) string {
	return mode.String()
}

func ReleaseModeToCV(mode FunctionMode) units.CV {
	return units.CV(europim.InverseLerp(mode, FunctionModeLinear, FunctionModeQuartic))
}

func CVToReleaseMode(cv units.CV) FunctionMode {
	return europim.LerpRound(cv.ToFloat32(), FunctionModeLinear, FunctionModeQuartic)
}
