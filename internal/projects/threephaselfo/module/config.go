package module

import (
	"github.com/heucuva/europi/units"
)

type Config struct {
	WaveMode  WaveMode
	Phi3Rate  units.CV
	SkewRate  units.CV
	SkewShape units.CV
	Degree0   func(cv units.CV)
	Degree120 func(cv units.CV)
	Degree240 func(cv units.CV)
}
