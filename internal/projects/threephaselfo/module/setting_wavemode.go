package module

import (
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

func WaveModeString(mode WaveMode) string {
	return mode.String()
}

var waveModeLerp = lerp.NewLerp32(WaveModeSine, WaveModeSine)

func WaveModeToCV(mode WaveMode) units.CV {
	return units.CV(waveModeLerp.ClampedInverseLerp(mode))
}

func CVToWaveMode(cv units.CV) WaveMode {
	return waveModeLerp.ClampedLerpRound(cv.ToFloat32())
}
