package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func SampleSlewBString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func SampleSlewBToCV(cv units.CV) units.CV {
	return cv
}

func CVToSampleSlewB(cv units.CV) units.CV {
	return cv
}
