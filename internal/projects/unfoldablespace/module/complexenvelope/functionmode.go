package complexenvelope

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

type functionModeCalc func(t, dur time.Duration) units.CV

func (e envelope) getFunctionModeCalc(mode FunctionMode) (functionModeCalc, error) {
	switch mode {
	case FunctionModeLinear:
		return modeFuncLinear, nil
	case FunctionModeExponential:
		return modeFuncExponential, nil
	case FunctionModeQuartic:
		return modeFuncQuartic, nil
	default:
		return nil, fmt.Errorf("unhandled attack mode: %q", mode)
	}
}
