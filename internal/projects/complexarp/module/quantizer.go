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

func getArpQuantizer(mode Quantizer) (quantizer[units.VOct], error) {
	switch mode {
	case QuantizerRound:
		return &quantizerRound[units.VOct]{}, nil
	case QuantizerTrunc:
		return &quantizerTrunc[units.VOct]{}, nil
	default:
		return nil, fmt.Errorf("unsupported quantizer mode: %d", mode)
	}
}
