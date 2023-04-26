package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func GateDensityString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func GateDensityToCV(cv units.CV) units.CV {
	return cv
}

func CVToGateDensity(cv units.CV) units.CV {
	return cv
}
