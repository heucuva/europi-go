package complexarp

import (
	"github.com/heucuva/europi/units"
)

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
