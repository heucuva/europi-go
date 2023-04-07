package complexarp

import (
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
)

type scale []units.VOct

func getScale(mode Scale) (scale, error) {
	var s scale

	switch mode {
	case ScaleC_Lydian:
		s = []units.VOct{keyC, keyD, keyE, keyFSharp, keyG, keyA, keyB}
	case ScaleC_Major:
		s = []units.VOct{keyC, keyD, keyE, keyF, keyG, keyA, keyB}
	case ScaleC_7:
		s = []units.VOct{keyC, keyE, keyG, keyBFlat}
	case ScaleC_Suspended:
		s = []units.VOct{keyC, keyD, keyF, keyG, keyBFlat}
	case ScaleC_Harmonic_5:
		s = []units.VOct{keyC, keyDFlat, keyE, keyF, keyG, keyAFlat, keyBFlat}
	case ScaleC_Dorian:
		s = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyA, keyBFlat}
	case ScaleC_Minor:
		s = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
	case ScaleC_Phrygian:
		s = []units.VOct{keyC, keyDFlat, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
	case ScaleC_Diminished:
		s = []units.VOct{keyC, keyD, keyEFlat, keyF, keyGFlat, keyAFlat, keyA, keyB}
	case ScaleC_Augmented:
		s = []units.VOct{keyC, keyDSharp, keyE, keyG, keyAFlat, keyB}
	}

	return s, nil
}
