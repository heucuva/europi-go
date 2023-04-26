package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ModeString(mode EnvelopeMode) string {
	return mode.String()
}

var modeLerp = lerp.NewLerp32(EnvelopeModeAD, EnvelopeModeAD)

func ModeToCV(mode EnvelopeMode) units.CV {
	return units.CV(modeLerp.ClampedInverseLerp(mode))
}

func CVToMode(cv units.CV) EnvelopeMode {
	return modeLerp.LerpRound(cv.ToFloat32())
}
