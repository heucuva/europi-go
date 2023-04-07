package complexrandom

import (
	"fmt"
	"time"
)

type Clock int

const (
	ClockFull = Clock(iota)
	ClockLimited
)

type clock interface {
	Tick(timeDelta time.Duration) bool
}

func getClock(mode Clock) (clock, error) {
	switch mode {
	case ClockFull:
		return &clockBasic{
			interval: fullSpectrum,
		}, nil
	case ClockLimited:
		return &clockBasic{
			interval: limitedSpectrum,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported clock mode: %d", mode)
	}
}
