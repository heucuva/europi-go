package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

const (
	MinPulseStageDivider = 1
	MaxPulseStageDivider = 16
)

func PulseStageDividerString(psd int) string {
	return fmt.Sprint(psd)
}

var pulseStageDividerLerp = lerp.NewLerp32(MinPulseStageDivider, MaxPulseStageDivider)

func PulseStageDividerToCV(psd int) units.CV {
	return units.CV(pulseStageDividerLerp.ClampedInverseLerp(psd))
}

func CVToPulseStageDivider(cv units.CV) int {
	return pulseStageDividerLerp.ClampedLerpRound(cv.ToFloat32())
}
