package module

import (
	"fmt"

	"github.com/heucuva/europi/input"
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToRange(cv units.CV) units.VOct {
	return europim.Lerp[units.VOct](cv.ToFloat32(), input.MinVoltage, input.MaxVoltage/2.0)
}

func RangeToCV(rng units.VOct) units.CV {
	return units.CV(europim.InverseLerp(rng, input.MinVoltage, input.MaxVoltage/2.0))
}

func RangeToString(rng units.VOct) string {
	return fmt.Sprintf("%2.1f Oct", rng)
}
