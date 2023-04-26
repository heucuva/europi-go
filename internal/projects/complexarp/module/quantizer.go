package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/quantizer"
	"github.com/awonak/EuroPiGo/units"
)

func (m *ComplexArp) setArpQuantizer(mode quantizer.Mode) error {
	switch mode {
	case quantizer.ModeRound:
		m.quantizer = &quantizer.Round[units.VOct]{}
	case quantizer.ModeTrunc:
		m.quantizer = &quantizer.Trunc[units.VOct]{}
	default:
		return fmt.Errorf("unsupported quantizer mode: %d", mode)
	}

	return nil
}
