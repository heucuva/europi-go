package threephaselfo

import "fmt"

type WaveMode int

const (
	WaveModeSine = WaveMode(iota)
)

func (m *Module) getWaveMode(mode WaveMode) (wave, error) {
	switch mode {
	case WaveModeSine:
		return &waveSine{}, nil
	default:
		return nil, fmt.Errorf("unsupported mode: %d", mode)
	}
}
