package module

import (
	"github.com/awonak/EuroPiGo/quantizer"
	"github.com/awonak/EuroPiGo/units"
)

type Config struct {
	ArpOut     func(voct units.VOct)
	ArpPattern Pattern
	Scale      Scale
	Quantizer  quantizer.Mode
	ArpPitch   units.VOct
	ArpRange   units.VOct
}
