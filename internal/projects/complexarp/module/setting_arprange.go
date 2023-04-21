package module

import (
	"fmt"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ArpRangeString(voct units.VOct) string {
	return fmt.Sprintf("%2.1f", voct)
}

func ArpRangeToCV(voct units.VOct) units.CV {
	return units.CV(europim.InverseLerp(voct, units.MinVOct, units.MaxVOct))
}

func CVToArpRange(cv units.CV) units.VOct {
	return europim.Lerp(cv.ToFloat32(), units.MinVOct, units.MaxVOct)
}
