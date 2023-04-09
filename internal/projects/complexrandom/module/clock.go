package module

import (
	"fmt"
	"time"

	"github.com/heucuva/europi/units"
)

type Clock int

const (
	ClockFull = Clock(iota)
	ClockLimited
)

type clock interface {
	Tick(timeDelta time.Duration) bool
	SetRate(cv units.CV)
	Rate() units.CV
}

func getClock(mode Clock) (clock, error) {
	switch mode {
	case ClockFull:
		return &clockBasic{
			spectrum: fullSpectrum,
		}, nil
	case ClockLimited:
		return &clockBasic{
			spectrum: limitedSpectrum,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported clock mode: %d", mode)
	}
}
