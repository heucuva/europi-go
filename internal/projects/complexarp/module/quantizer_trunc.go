package module

import (
	"math"

	europim "github.com/heucuva/europi/math"
)

type quantizerTrunc[T any] struct{}

func (quantizerTrunc[T]) QuantizeToIndex(in float32, length int) int {
	if length == 0 {
		return -1
	}

	idx := int(math.Trunc(float64(length-1) * float64(in)))
	idx = europim.Clamp(idx, 0, length-1)
	return idx
}

func (q quantizerTrunc[T]) QuantizeToValue(in float32, list []T) T {
	idx := q.QuantizeToIndex(in, len(list))
	if idx == -1 {
		var empty T
		return empty
	}

	return list[idx]
}
