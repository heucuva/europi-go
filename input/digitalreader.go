package input

import (
	"machine"
	"time"
)

// DigitalReader is an interface for common digital inputs methods.
type DigitalReader interface {
	Handler(func(machine.Pin))
	HandlerEx(machine.PinChange, func(machine.Pin))
	HandlerWithDebounce(func(machine.Pin), time.Duration)
	HandlerExWithDebounce(machine.PinChange, func(machine.Pin), time.Duration)
	LastInput() time.Time
	Value() bool
}
