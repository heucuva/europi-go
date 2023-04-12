package module

import (
	"fmt"

	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/units"
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
