package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func ArpRangeString(voct units.VOct) string {
	return fmt.Sprintf("%2.1f", voct)
}

var arpRangeLerp = lerp.NewLerp32(units.MinVOct, units.MaxVOct)

func ArpRangeToCV(voct units.VOct) units.CV {
	return units.CV(arpRangeLerp.ClampedInverseLerp(voct))
}

func CVToArpRange(cv units.CV) units.VOct {
	return arpRangeLerp.ClampedLerp(cv.ToFloat32())
}
