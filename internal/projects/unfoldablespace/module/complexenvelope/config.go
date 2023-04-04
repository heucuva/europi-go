package complexenvelope

import "github.com/heucuva/europi/units"

type Config struct {
	Env [2]EnvelopeConfig
}

type EnvelopeConfig struct {
	Out         func(cv units.CV)
	Mode        string
	AttackMode  string
	ReleaseMode string
	Attack      units.CV
	Decay       units.CV
}
