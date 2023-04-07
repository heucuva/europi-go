package complexenvelope

import (
	"fmt"
	"time"

	"github.com/heucuva/europi/units"
)

type envelope struct {
	model model
}

func noop(_ units.CV) {
}

func (e *envelope) Init(config EnvelopeConfig) error {
	out := config.Out
	if out == nil {
		out = noop
	}

	switch config.Mode {
	case EnvelopeModeAD:
		attackDur := time.Duration(float32(time.Second) * config.Attack.ToFloat32())
		decayDur := time.Duration(float32(time.Second) * config.Decay.ToFloat32())
		amode, err := e.getFunctionMode(config.AttackMode)
		if err != nil {
			return fmt.Errorf("attack: %w", err)
		}
		dmode, err := e.getFunctionMode(config.ReleaseMode)
		if err != nil {
			return fmt.Errorf("decay: %w", err)
		}
		e.model = &modelAD{
			out:       out,
			atten:     1.0,
			attack:    amode,
			attackDur: attackDur,
			decay:     dmode,
			decayDur:  decayDur,
		}

	default:
		return fmt.Errorf("unhandled mode: %q", config.Mode)
	}
	return nil
}

func (e *envelope) SetCV(cv units.CV) {
	e.model.SetCV(cv)
}

func (e *envelope) Trigger() {
	e.model.Trigger()
}

func (e *envelope) Tick(deltaTime time.Duration) {
	e.model.Tick(deltaTime)
}
