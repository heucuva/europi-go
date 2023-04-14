package module

import (
	"fmt"
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	MinGateDuration time.Duration = time.Microsecond
	MaxGateDuration time.Duration = time.Millisecond * 990
)

func CVToGateDuration(cv units.CV) time.Duration {
	return europim.Lerp(cv.ToFloat32(), MinGateDuration, MaxGateDuration)
}

func GateDurationToCV(gate time.Duration) units.CV {
	return units.CV(europim.InverseLerp(gate, MinGateDuration, MaxGateDuration))
}

func GateDurationToString(gate time.Duration) string {
	switch {
	case gate < time.Millisecond:
		return fmt.Sprintf("%3.1fus", gate.Seconds()*1000000.0)
	case gate < time.Second:
		return fmt.Sprintf("%3.1fms", gate.Seconds()*1000.0)
	default:
		return fmt.Sprint(gate)
	}
}
