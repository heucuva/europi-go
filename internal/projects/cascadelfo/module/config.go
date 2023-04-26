package module

import "github.com/awonak/EuroPiGo/units"

type Config struct {
	LFO              [8]func(cv units.BipolarCV)
	Rate             units.Hertz
	RateAttenuverter units.BipolarCV
}
