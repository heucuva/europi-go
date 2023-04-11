package screenbank

import (
	"time"

	"github.com/heucuva/europi"
)

type ScreenBankOption func(sb *ScreenBank) error

func WithScreen(name string, screen europi.UserInterface) ScreenBankOption {
	return func(sb *ScreenBank) error {
		e := screenBankEntry{
			name:       name,
			screen:     screen,
			enabled:    true,
			locked:     true,
			lastUpdate: time.Now(),
		}

		sb.bank = append(sb.bank, e)
		return nil
	}
}
