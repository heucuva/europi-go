package complexarp

import "github.com/heucuva/europi/units"

type Config struct {
	ArpOut        func(voct units.VOct)
	ArpPattern    string
	ChordMode     string
	QuantizerMode string
	ArpRange      units.VOct
	ArpPitch      units.VOct
}
