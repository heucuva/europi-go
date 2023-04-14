package module

import (
	"fmt"

	"github.com/heucuva/europi/input"
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToPitch(cv units.CV) units.VOct {
	return europim.Lerp[units.VOct](cv.ToFloat32(), input.MinVoltage, input.MaxVoltage)
}

func PitchToCV(pitch units.VOct) units.CV {
	return units.CV(europim.InverseLerp(pitch, input.MinVoltage, input.MaxVoltage))
}

func PitchToString(pitch units.VOct) string {
	return fmt.Sprintf("%2.1fV/Oct", pitch)
}
