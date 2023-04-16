package module

import (
	"math"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func CVToWaveMode(cv units.CV) WaveMode {
	return europim.LerpRound(cv.ToFloat32(), WaveModeSine, WaveModeSine)
}

func WaveModeToCV(wavemode WaveMode) units.CV {
	return units.CV(europim.Clamp(float32(math.Log(float64(wavemode)))/10.0, 0.0, 1.0))
}

func WaveModeToString(wavemode WaveMode) string {
	switch wavemode {
	case WaveModeSine:
		return "sine"
	default:
		return ""
	}
}
