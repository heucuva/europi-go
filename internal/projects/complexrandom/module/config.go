package module

import "github.com/awonak/EuroPiGo/units"

type Config struct {
	SampleOutA        func(cv units.CV)
	SampleOutB        func(cv units.CV)
	SampleAttenuatorA units.CV
	IntegrationSlope  units.CV
	GateDensity       units.CV
	PulseStageDivider int
	SampleAttenuatorB units.CV
	SampleSlewB       units.CV
	ClockSpeed        units.CV
	ClockRange        Clock
}