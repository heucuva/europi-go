package complexarp

import (
	"strings"

	"github.com/heucuva/europi/units"
)

type keyboard struct {
	keys []units.VOct
}

const (
	keyC      = units.VOct(0)       // C
	keyCSharp = units.VOct(1) / 12  // C#
	keyDFlat  = keyCSharp           // Db
	keyD      = units.VOct(2) / 12  // D
	keyDSharp = units.VOct(3) / 12  // D#
	keyEFlat  = keyDSharp           // Eb
	keyE      = units.VOct(4) / 12  // E
	keyF      = units.VOct(5) / 12  // F
	keyFSharp = units.VOct(6) / 12  // F#
	keyGFlat  = keyFSharp           // Gb
	keyG      = units.VOct(7) / 12  // G
	keyGSharp = units.VOct(8) / 12  // G#
	keyAFlat  = keyGSharp           // Ab
	keyA      = units.VOct(9) / 12  // A
	keyASharp = units.VOct(10) / 12 // A#
	keyBFlat  = keyASharp           // Bb
	keyB      = units.VOct(11) / 12 // B
)

func getArpKeyboard(config Config) (*keyboard, error) {
	kb := &keyboard{}

	switch strings.ToLower(config.ChordMode) {
	case "c_lyd":
		kb.keys = []units.VOct{keyC, keyD, keyE, keyFSharp, keyG, keyA, keyB}
	case "c_maj":
		kb.keys = []units.VOct{keyC, keyD, keyE, keyF, keyG, keyA, keyB}
	case "c_7":
		kb.keys = []units.VOct{keyC, keyE, keyG, keyBFlat}
	case "c_sus":
		kb.keys = []units.VOct{keyC, keyD, keyF, keyG, keyBFlat}
	case "c_hm5":
		kb.keys = []units.VOct{keyC, keyDFlat, keyE, keyF, keyG, keyAFlat, keyBFlat}
	case "c_dor":
		kb.keys = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyA, keyBFlat}
	case "c_min":
		kb.keys = []units.VOct{keyC, keyD, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
	case "c_phr":
		kb.keys = []units.VOct{keyC, keyDFlat, keyEFlat, keyF, keyG, keyAFlat, keyBFlat}
	case "c_dim":
		kb.keys = []units.VOct{keyC, keyD, keyEFlat, keyF, keyGFlat, keyAFlat, keyA, keyB}
	case "c_aug":
		kb.keys = []units.VOct{keyC, keyDSharp, keyE, keyG, keyAFlat, keyB}
	}

	return kb, nil
}
