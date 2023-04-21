package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func ScaleString(s Scale) string {
	switch s {
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

func ScaleToCV(s Scale) units.CV {
	return units.CV(europim.InverseLerp(s, ScaleC_Lydian, ScaleC_Augmented))
}

func CVToScale(cv units.CV) Scale {
	return europim.LerpRound(cv.ToFloat32(), ScaleC_Lydian, ScaleC_Augmented)
}
