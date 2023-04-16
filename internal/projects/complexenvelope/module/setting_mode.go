package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ModeString(mode EnvelopeMode) string {
	return mode.String()
}

func ModeToCV(mode EnvelopeMode) units.CV {
	return units.CV(europim.InverseLerp(mode, EnvelopeModeAD, EnvelopeModeAD))
}

func CVToMode(cv units.CV) EnvelopeMode {
	return europim.LerpRound(cv.ToFloat32(), EnvelopeModeAD, EnvelopeModeAD)
}
