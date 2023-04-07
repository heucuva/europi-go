package complexarp

import "github.com/heucuva/europi/units"

type Config struct {
	ArpOut     func(voct units.VOct)
	ArpPattern Pattern
	Scale      Scale
	Quantizer  Quantizer
	ArpRange   units.VOct
	ArpPitch   units.VOct
}
