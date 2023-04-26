package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func SampleAttenuatorBString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func SampleAttenuatorBToCV(cv units.CV) units.CV {
	return cv
}

func CVToSampleAttenuatorB(cv units.CV) units.CV {
	return cv
}
