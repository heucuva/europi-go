package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func CVToRateAV(cv units.CV) units.BipolarCV {
	return cv.ToBipolarCV()
}

func RateAVToCV(rateav units.BipolarCV) units.CV {
	return rateav.ToCV()
}

func RateAVToString(rateav units.BipolarCV) string {
	return fmt.Sprintf("%+3.1f%%", rateav*100.0)
}
