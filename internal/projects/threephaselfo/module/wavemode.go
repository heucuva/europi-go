package module

import (
	"fmt"

	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/units"
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

var (
	waveModeQuant quantizer.Round[WaveMode]
	cvWaveModes   = []WaveMode{
		WaveModeSine,
	}
)

func GetWaveModeCV(mode WaveMode) units.CV {
	for i, v := range cvWaveModes {
		if v == mode {
			return units.CV(i) / units.CV(len(cvWaveModes))
		}
	}
	panic("unknown mode")
}
