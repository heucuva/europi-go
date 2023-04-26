package module

import "github.com/awonak/EuroPiGo/units"

type Config struct {
	SampleOutA        func(cv units.BipolarCV)
	SampleOutB        func(cv units.BipolarCV)
	SampleAttenuatorA units.BipolarCV
	IntegrationSlope  units.CV
	GateDensity       units.CV
	PulseStageDivider int
	SampleAttenuatorB units.BipolarCV
	SampleSlewB       units.CV
	ClockSpeed        units.CV
	ClockRange        Clock
}
