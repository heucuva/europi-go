package module

import (
	"math/rand"
)

type noiseWhite struct{}

func (n noiseWhite) Get() float32 {
	return rand.Float32()
}
