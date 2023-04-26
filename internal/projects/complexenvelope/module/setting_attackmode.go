package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func AttackModeString(mode FunctionMode) string {
	return mode.String()
}

var attackModeLerp = lerp.NewLerp32(FunctionModeLinear, FunctionModeQuartic)

func AttackModeToCV(mode FunctionMode) units.CV {
	return units.CV(attackModeLerp.ClampedInverseLerp(mode))
}

func CVToAttackMode(cv units.CV) FunctionMode {
	return attackModeLerp.ClampedLerpRound(cv.ToFloat32())
}
