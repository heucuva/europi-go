package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func SampleAttenuatorAString(cv units.BipolarCV) string {
	return fmt.Sprintf("%+3.1f%%", cv*100.0)
}

func SampleAttenuatorAToCV(cv units.BipolarCV) units.CV {
	return units.CV((cv.ToFloat32() + 1.0) * 0.5)
}

func CVToSampleAttenuatorA(cv units.CV) units.BipolarCV {
	return units.BipolarCV(cv.ToFloat32()*2.0 - 1.0)
}
