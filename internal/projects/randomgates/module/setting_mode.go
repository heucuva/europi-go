package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ModeString(mode Mode) string {
	return mode.String()
}

func ModeToCV(mode Mode) units.CV {
	return units.CV(europim.InverseLerp(mode, Mode1msTrig, ModeEqualGateTrig))
}

func CVToMode(cv units.CV) Mode {
	return europim.LerpRound(cv.ToFloat32(), Mode1msTrig, ModeEqualGateTrig)
}
