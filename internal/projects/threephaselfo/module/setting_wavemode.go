package module

import (
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

func WaveModeString(mode WaveMode) string {
	return mode.String()
}

func WaveModeToCV(mode WaveMode) units.CV {
	return units.CV(europim.InverseLerp(mode, WaveModeSine, WaveModeSine))
}

func CVToWaveMode(cv units.CV) WaveMode {
	return europim.LerpRound(cv.ToFloat32(), WaveModeSine, WaveModeSine)
}
