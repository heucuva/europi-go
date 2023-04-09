package module

import (
	"fmt"
	"time"

	"github.com/heucuva/europi/units"
)

type FunctionMode int

const (
	FunctionModeLinear = FunctionMode(iota)
	FunctionModeExponential
	FunctionModeQuartic
)

type functionMode interface {
	Calc(t, dur time.Duration) units.CV
}

func (e envelope) getFunctionMode(mode FunctionMode) (functionMode, error) {
	switch mode {
	case FunctionModeLinear:
		return &functionModeLinear{}, nil
	case FunctionModeExponential:
		return &functionModeExponential{}, nil
	case FunctionModeQuartic:
		return &functionModeQuartic{}, nil
	default:
		return nil, fmt.Errorf("unhandled function mode: %q", mode)
	}
}
