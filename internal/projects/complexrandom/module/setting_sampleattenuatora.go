package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func SampleAttenuatorAString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func SampleAttenuatorAToCV(cv units.CV) units.CV {
	return cv
}

func CVToSampleAttenuatorA(cv units.CV) units.CV {
	return cv
}
