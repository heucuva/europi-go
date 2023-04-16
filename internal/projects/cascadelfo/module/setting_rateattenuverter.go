package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func RateAttenuverterString(cv units.BipolarCV) string {
	return fmt.Sprintf("%+3.1f%%", cv*100.0)
}

func RateAttenuverterToCV(cv units.BipolarCV) units.CV {
	return cv.ToCV()
}

func CVToRateAttenuverter(cv units.CV) units.BipolarCV {
	return cv.ToBipolarCV()
}
