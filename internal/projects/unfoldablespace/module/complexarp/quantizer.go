package complexarp

import (
	"fmt"
	"strings"

	"github.com/heucuva/europi/units"
)

type quantizer[T any] interface {
	QuantizeToIndex(in float32, length int) int
	QuantizeToValue(in float32, list []T) T
}

func getArpQuantizer(config Config) (quantizer[units.VOct], error) {
	switch strings.ToLower(config.QuantizerMode) {
	case "round":
		return &quantizerRound[units.VOct]{}, nil
	case "trunc":
		return &quantizerTrunc[units.VOct]{}, nil
	default:
		return nil, fmt.Errorf("unsupported quantizer mode: %q", config.QuantizerMode)
	}
}
