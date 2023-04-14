package module

import (
	"fmt"
)

type WaveMode int

const (
	WaveModeSine = WaveMode(iota)
)

func getWaveMode(mode WaveMode) (wave, error) {
	switch mode {
	case WaveModeSine:
		return &waveSine{}, nil
	default:
		return nil, fmt.Errorf("unsupported mode: %d", mode)
	}
}
