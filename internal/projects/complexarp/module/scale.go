package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
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
)

type scale interface {
	Keys() []units.VOct
	Mode() Scale
}

func (m *ComplexArp) setScale(mode Scale) error {
	switch mode {
	case ScaleC_Lydian:
		m.scale = &scaleCLydian{}
	case ScaleC_Major:
		m.scale = &scaleCMajor{}
	case ScaleC_7:
		m.scale = &scaleC7{}
	case ScaleC_Suspended:
		m.scale = &scaleCSuspended{}
	case ScaleC_Harmonic_5:
		m.scale = &scaleCHarmonic5{}
	case ScaleC_Dorian:
		m.scale = &scaleCDorian{}
	case ScaleC_Minor:
		m.scale = &scaleCMinor{}
	case ScaleC_Phrygian:
		m.scale = &scaleCPhrygian{}
	case ScaleC_Diminished:
		m.scale = &scaleCDiminished{}
	case ScaleC_Augmented:
		m.scale = &scaleCAugmented{}
	default:
		return fmt.Errorf("unsupported scale: %d", mode)
	}

	return nil
}
