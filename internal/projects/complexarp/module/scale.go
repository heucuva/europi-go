package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

type Scale int

const (
	ScaleC_Lydian = Scale(iota)
	ScaleC_Major
	ScaleC_7
	ScaleC_Suspended
	ScaleC_Harmonic_5
	ScaleC_Dorian
	ScaleC_Minor
	ScaleC_Phrygian
	ScaleC_Diminished
	ScaleC_Augmented

	//====
	cScaleCount
)

type scale interface {
	Keys() []units.VOct
	Mode() Scale
	Name() string
}

func getScale(mode Scale) (scale, error) {
	switch mode {
	case ScaleC_Lydian:
		return &scaleCLydian{}, nil
	case ScaleC_Major:
		return &scaleCMajor{}, nil
	case ScaleC_7:
		return &scaleC7{}, nil
	case ScaleC_Suspended:
		return &scaleCSuspended{}, nil
	case ScaleC_Harmonic_5:
		return &scaleCHarmonic5{}, nil
	case ScaleC_Dorian:
		return &scaleCDorian{}, nil
	case ScaleC_Minor:
		return &scaleCMinor{}, nil
	case ScaleC_Phrygian:
		return &scaleCPhrygian{}, nil
	case ScaleC_Diminished:
		return &scaleCDiminished{}, nil
	case ScaleC_Augmented:
		return &scaleCAugmented{}, nil
	default:
		return nil, fmt.Errorf("unsupported scale: %d", mode)
	}
}
