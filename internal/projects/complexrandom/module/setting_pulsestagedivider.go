package module

import (
	"fmt"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

const (
	MinPulseStageDivider = 1
	MaxPulseStageDivider = 16
)

func PulseStageDividerString(psd int) string {
	return fmt.Sprint(psd)
}

func PulseStageDividerToCV(psd int) units.CV {
	return units.CV(europim.InverseLerp(psd, MinPulseStageDivider, MaxPulseStageDivider))
}

func CVToPulseStageDivider(cv units.CV) int {
	return europim.LerpRound(cv.ToFloat32(), MinPulseStageDivider, MaxPulseStageDivider)
}
