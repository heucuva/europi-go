package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToScale(cv units.CV) Scale {
	return europim.Lerp(cv.ToFloat32(), ScaleC_Lydian, ScaleC_Augmented)
}

func ScaleToCV(mode Scale) units.CV {
	return units.CV(europim.InverseLerp(mode, ScaleC_Lydian, ScaleC_Augmented))
}

func ScaleToString(mode Scale) string {
	switch mode {
	case ScaleC_Lydian:
		return "C lyd"
	case ScaleC_Major:
		return "C maj"
	case ScaleC_7:
		return "C 7"
	case ScaleC_Suspended:
		return "C sus"
	case ScaleC_Harmonic_5:
		return "C hm5"
	case ScaleC_Dorian:
		return "C dor"
	case ScaleC_Minor:
		return "C min"
	case ScaleC_Phrygian:
		return "C phr"
	case ScaleC_Diminished:
		return "C dim"
	case ScaleC_Augmented:
		return "C aug"
	default:
		return ""
	}
}
