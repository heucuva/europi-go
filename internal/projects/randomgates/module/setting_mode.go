package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ModeString(mode Mode) string {
	return mode.String()
}

var modeLerp = lerp.NewLerp32(Mode1msTrig, ModeEqualGateTrig)

func ModeToCV(mode Mode) units.CV {
	return units.CV(modeLerp.ClampedInverseLerp(mode))
}

func CVToMode(cv units.CV) Mode {
	return modeLerp.ClampedLerpRound(cv.ToFloat32())
}
