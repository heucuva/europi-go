package module

import (
	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/units"
)

type Config struct {
	ArpOut     func(voct units.VOct)
	ArpPattern Pattern
	Scale      Scale
	Quantizer  quantizer.Mode
	ArpRange   units.VOct
	ArpPitch   units.VOct
}
