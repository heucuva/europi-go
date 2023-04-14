package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToMode(cv units.CV) Mode {
	return europim.Lerp(cv.ToFloat32(), Mode1msTrig, ModeEqualGateTrig)
}

func ModeToCV(mode Mode) units.CV {
	return units.CV(europim.InverseLerp(mode, Mode1msTrig, ModeEqualGateTrig))
}

func ModeToString(mode Mode) string {
	switch mode {
	case Mode1msTrig:
		return "1ms"
	case Mode200msTrig:
		return "200ms"
	case ModeQuarterGateTrig:
		return "1/4"
	case ModeHalfGateTrig:
		return "1/2"
	case ModeEqualGateTrig:
		return "1:1"
	default:
		panic("unsupported mode")
	}
}
