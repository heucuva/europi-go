package input

import (
	"machine"
	"runtime/interrupt"
	"time"
)

const DefaultDebounceDelay = time.Duration(50 * time.Millisecond)

// Digital is a struct for handling reading of the digital input.
type Digital struct {
	Pin           machine.Pin
	debounceDelay time.Duration
	lastInput     time.Time
	callback      func(p machine.Pin)
}

// NewDigital creates a new Digital struct.
func NewDigital(pin machine.Pin) *Digital {
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return &Digital{
		Pin:           pin,
		lastInput:     time.Now(),
		debounceDelay: DefaultDebounceDelay,
	}
}

// LastInput return the time of the last high input (triggered at 0.8v).
func (d *Digital) LastInput() time.Time {
	return d.lastInput
}

// Value returns true if the input is high (above 0.8v), else false.
func (d *Digital) Value() bool {
	state := interrupt.Disable()
	// Invert signal to match expected behavior.
	v := !d.Pin.Get()
	interrupt.Restore(state)
	return v
}

// Handler sets the callback function to be call when a rising edge is detected.
func (d *Digital) Handler(handler func(p machine.Pin)) {
	d.HandlerExWithDebounce(machine.PinRising, handler, 0)
}

// HandlerEx sets the callback function to be call when a specified edge is (or edges are) detected.
func (d *Digital) HandlerEx(pinChange machine.PinChange, handler func(p machine.Pin)) {
	d.HandlerExWithDebounce(pinChange, handler, 0)
}

// Handler sets the callback function to be call when a rising edge is detected and debounce delay time has elapsed.
func (d *Digital) HandlerWithDebounce(handler func(p machine.Pin), delay time.Duration) {
	d.HandlerExWithDebounce(machine.PinRising, handler, delay)
}

func (d *Digital) HandlerExWithDebounce(pinChange machine.PinChange, handler func(p machine.Pin), delay time.Duration) {
	d.callback = handler
	d.debounceDelay = delay
	state := interrupt.Disable()
	d.Pin.SetInterrupt(pinChange, d.debounceWrapper)
	interrupt.Restore(state)
}

func (d *Digital) debounceWrapper(p machine.Pin) {
	t := time.Now()
	if t.Before(d.lastInput.Add(d.debounceDelay)) {
		return
	}
	d.callback(p)
	d.lastInput = t
}
