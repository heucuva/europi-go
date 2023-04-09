package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

type Quantizer int

const (
	QuantizerRound = Quantizer(iota)
	QuantizerTrunc
)

type quantizer[T any] interface {
	QuantizeToIndex(in float32, length int) int
	QuantizeToValue(in float32, list []T) T
}

func (m *ComplexArp) setArpQuantizer(mode Quantizer) error {
	switch mode {
	case QuantizerRound:
		m.quantizer = &quantizerRound[units.VOct]{}
	case QuantizerTrunc:
		m.quantizer = &quantizerTrunc[units.VOct]{}
	default:
		return fmt.Errorf("unsupported quantizer mode: %d", mode)
	}

	return nil
}
