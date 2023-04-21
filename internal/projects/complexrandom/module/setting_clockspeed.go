package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func ClockSpeedString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func ClockSpeedToCV(cv units.CV) units.CV {
	return cv
}

func CVToClockSpeed(cv units.CV) units.CV {
	return cv
}
