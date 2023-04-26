package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ArpPitchString(voct units.VOct) string {
	return fmt.Sprintf("%2.1f", voct)
}

var arpPitchLerp = lerp.NewLerp32(units.MinVOct, units.MaxVOct)

func ArpPitchToCV(voct units.VOct) units.CV {
	return units.CV(arpPitchLerp.ClampedInverseLerp(voct))
}

func CVToArpPitch(cv units.CV) units.VOct {
	return arpPitchLerp.ClampedLerp(cv.ToFloat32())
}
