package module

import "github.com/heucuva/europi/units"

type Config struct {
	Degree0   func(cv units.BipolarCV)
	Degree120 func(cv units.BipolarCV)
	Degree240 func(cv units.BipolarCV)
	WaveMode  WaveMode
	Phi3Rate  units.Hertz
	SkewRate  units.Hertz
	SkewShape units.CV
}
