package module

import "github.com/heucuva/europi/units"

type Config struct {
	LFO              [8]func(cv units.CV)
	Rate             units.Hertz
	RateAttenuverter float32
}
