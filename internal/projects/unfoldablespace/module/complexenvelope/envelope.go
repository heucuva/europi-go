package complexenvelope

import (
	"fmt"
	"strings"
	"time"

	"github.com/heucuva/europi/units"
)

type envelope struct {
	model model
}

func (e *envelope) Init(config EnvelopeConfig) error {
	out := config.Out
	if out == nil {
		out = func(cv units.CV) {}
	}

	switch strings.ToLower(config.Mode) {
	case "ad":
		attackDur := time.Duration(float32(time.Second) * config.Attack.ToFloat32())
		decayDur := time.Duration(float32(time.Second) * config.Decay.ToFloat32())
		amode, err := e.getModeFunc(config.AttackMode)
		if err != nil {
			return err
		}
		dmode, err := e.getModeFunc(config.ReleaseMode)
		if err != nil {
			return err
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

func (e envelope) getModeFunc(mode string) (modeFunc, error) {
	switch strings.ToLower(mode) {
	case "linear":
		return modeFuncLinear, nil
	case "exponential":
		return modeFuncExponential, nil
	case "quartic":
		return modeFuncQuartic, nil
	default:
		return nil, fmt.Errorf("unhandled attack mode: %q", mode)
	}
}
