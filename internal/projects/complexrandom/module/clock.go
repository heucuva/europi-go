package module

import (
	"fmt"
	"time"

	"github.com/awonak/EuroPiGo/units"
)

type Clock int

const (
	ClockFull = Clock(iota)
	ClockLimited
)

const (
	fullSpectrum    = time.Second / 22050
	limitedSpectrum = time.Second * 15 / 22050
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

func (c Clock) String() string {
	switch c {
	case ClockFull:
		return "Full"
	case ClockLimited:
		return "Limited"
	default:
		return ""
	}
}
