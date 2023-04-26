package module

import "github.com/awonak/EuroPiGo/units"

type Config struct {
	Env [2]EnvelopeConfig
}

type EnvelopeConfig struct {
	Out         func(cv units.CV)
	Mode        EnvelopeMode
	AttackMode  FunctionMode
	ReleaseMode FunctionMode
	Attack      units.CV
	Decay       units.CV
}
