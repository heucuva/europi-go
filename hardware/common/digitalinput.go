package common

import (
	"time"

	"github.com/awonak/EuroPiGo/debounce"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

// Digitalinput is a struct for handling reading of the digital input.
type Digitalinput struct {
	dr         DigitalReaderProvider
	lastChange time.Time
}

var (
	// static check
	_ hal.DigitalInput = (*Digitalinput)(nil)
	// silence linter
	_ = NewDigitalInput
)

type DigitalReaderProvider interface {
	Get() bool
	SetHandler(changes hal.ChangeFlags, handler func())
}

// NewDigitalInput creates a new digital input struct.
func NewDigitalInput(dr DigitalReaderProvider) *Digitalinput {
	if dr == nil {
		return nil
	}
	return &Digitalinput{
		dr:         dr,
		lastChange: time.Now(),
	}
}

// Configure updates the device with various configuration parameters
func (d *Digitalinput) Configure(config hal.DigitalInputConfig) error {
	return nil
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *Digitalinput) Value() bool {
	return d.dr.Get()
}

// Handler sets the callback function to be call when the incoming signal going high event is detected.
func (d *Digitalinput) Handler(handler func(value bool, deltaTime time.Duration)) {
	d.HandlerEx(hal.ChangeRising, handler)
}

// HandlerEx sets the callback function to be call when the input changes in a specified way.
func (d *Digitalinput) HandlerEx(changes hal.ChangeFlags, handler func(value bool, deltaTime time.Duration)) {
	d.dr.SetHandler(changes, func() {
		now := time.Now()
		timeDelta := now.Sub(d.lastChange)
		handler(d.Value(), timeDelta)
		d.lastChange = now
	})
}

// HandlerWithDebounce sets the callback function to be call when the incoming signal going high event is detected and debounce delay time has elapsed.
func (d *Digitalinput) HandlerWithDebounce(handler func(value bool, deltaTime time.Duration), delay time.Duration) {
	db := debounce.NewDebouncer(handler).Debounce(delay)
	d.Handler(func(value bool, _ time.Duration) {
		// throw away the deltaTime coming in on the handler
		// we want to use what's on the debouncer, instead
		db(value)
	})
}
