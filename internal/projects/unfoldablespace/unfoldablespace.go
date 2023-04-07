package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module"
	"github.com/heucuva/europi/output"
	"github.com/heucuva/europi/units"
)

var (
	unfold module.UnfoldableSpace
)

func startLoop(e *europi.EuroPi) {
	setVOct := func(voct units.VOct) {
		e.CV1.SetVoltage(voct.ToVolts())
	}
	setLevel := func(cv units.CV) {
		e.CV2.SetVoltage(cv.ToVolts())
	}
	setTimbre := func(cv units.CV) {
		e.CV3.SetVoltage(cv.ToVolts())
	}
	setHarmo := func(cv units.CV) {
		e.CV4.SetVoltage(cv.ToVolts())
	}
	setMorph := func(cv units.CV) {
		e.CV5.SetVoltage(cv.ToVolts())
	}
	setLFOCV := func(cv units.CV) {
		e.CV6.SetVoltage(cv.ToVolts())
	}

	if err := unfold.Init(module.Config{
		SetVOct:   setVOct,
		SetLevel:  setLevel,
		SetTimbre: setTimbre,
		SetHarmo:  setHarmo,
		SetMorph:  setMorph,
		SetLFOCV:  setLFOCV,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		unfold.Clock()
	})

	e.B1.HandlerWithDebounce(func(p machine.Pin) {
		unfold.ToggleInternalClock()
	}, time.Millisecond*500)
}

var (
	displayDelay time.Duration
)

const (
	displayRate       = time.Millisecond * 150
	line1y      int16 = 11
	line2y      int16 = 23
)

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	unfold.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		if unfold.InternalClockEnabled() {
			disp.DrawHLine(0, 0, 7, output.White)
		}
		disp.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y)
		disp.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y)
		disp.Display()
	}
}

func main() {
	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
	)
}
