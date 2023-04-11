package screenbank

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
)

type ScreenBank struct {
	screen  europi.UserInterface
	current int
	bank    []screenBankEntry
}

func NewScreenBank(opts ...ScreenBankOption) (*ScreenBank, error) {
	sb := &ScreenBank{}

	for _, opt := range opts {
		if err := opt(sb); err != nil {
			return nil, err
		}
	}

	return sb, nil
}

func (sb *ScreenBank) CurrentName() string {
	if len(sb.bank) == 0 {
		return ""
	}
	return sb.bank[sb.current].name
}

func (sb *ScreenBank) Current() europi.UserInterface {
	if len(sb.bank) == 0 {
		return nil
	}
	return sb.bank[sb.current].screen
}

func (sb *ScreenBank) Next() {
	if len(sb.bank) == 0 {
		sb.current = 0
		return
	}

	cur := &sb.bank[sb.current]
	cur.lock()

	sb.current++
	if sb.current >= len(sb.bank) {
		sb.current = 0
	}
	sb.bank[sb.current].unlock()
}

func (sb *ScreenBank) Start(e *europi.EuroPi) {
	for i := range sb.bank {
		s := &sb.bank[i]

		s.lock()
		s.screen.Start(e)
		s.lastUpdate = time.Now()
		s.unlock()
	}
}

func (sb *ScreenBank) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	cur := &sb.bank[sb.current]
	cur.lock()
	now := time.Now()
	cur.screen.Paint(e, now.Sub(cur.lastUpdate))
	cur.lastUpdate = now
	cur.unlock()
}

func (sb *ScreenBank) Button1(e *europi.EuroPi, p machine.Pin) {
	if cur, ok := sb.Current().(europi.UserInterfaceButton1); ok {
		cur.Button1(e, p)
	}
}

func (sb *ScreenBank) Button1Long(e *europi.EuroPi, p machine.Pin) {
	if cur, ok := sb.Current().(europi.UserInterfaceButton1Long); ok {
		cur.Button1Long(e, p)
	} else {
		// try the short-press
		sb.Button1(e, p)
	}
}

func (sb *ScreenBank) Button2(e *europi.EuroPi, p machine.Pin) {
	if cur, ok := sb.Current().(europi.UserInterfaceButton2); ok {
		cur.Button2(e, p)
	}
}

func (sb *ScreenBank) Button2Long(e *europi.EuroPi, p machine.Pin) {
	sb.Next()
}
