package module

import (
	"fmt"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

type envelope struct {
	model  model
	config EnvelopeConfig
}

func noopOut(cv units.CV) {
}

func (e *envelope) Init(config EnvelopeConfig) error {
	out := config.Out
	if out == nil {
		out = noopOut
	}

	e.config = config

	switch config.Mode {
	case EnvelopeModeAD:
		amode, err := e.getFunctionMode(config.AttackMode)
		if err != nil {
			return fmt.Errorf("attack: %w", err)
		}
		dmode, err := e.getFunctionMode(config.ReleaseMode)
		if err != nil {
			return fmt.Errorf("decay: %w", err)
		}
		e.model = &modelAD{
			out:    out,
			atten:  1.0,
			attack: amode,
			decay:  dmode,
		}
		e.model.SetAttack(config.Attack)
		e.model.SetDecay(config.Decay)

	default:
		return fmt.Errorf("unhandled mode: %q", config.Mode)
	}
	return nil
}

func (e *envelope) SetCV(cv units.BipolarCV) {
	e.model.SetCV(cv)
}

func (e *envelope) SetAttack(cv units.CV) {
	e.config.Attack = cv
	e.model.SetAttack(cv)
}

func (e *envelope) SetDecay(cv units.CV) {
	e.config.Decay = cv
	e.model.SetDecay(cv)
}

func (e *envelope) Trigger() {
	e.model.Trigger()
}

func (e *envelope) Tick(deltaTime time.Duration) {
	e.model.Tick(deltaTime)
}
