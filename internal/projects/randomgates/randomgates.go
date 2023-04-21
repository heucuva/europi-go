package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	clockScreen "github.com/heucuva/europi/internal/projects/clockgenerator/screen"
	"github.com/heucuva/europi/internal/projects/randomgates/module"
	"github.com/heucuva/europi/internal/projects/randomgates/screen"
	"github.com/heucuva/europi/output"
)

var (
	trig  module.RandomGates
	clock clockgenerator.ClockGenerator

	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		RandomGates: &trig,
		Clock:       &clock,
	}
	screenClock = clockScreen.Settings{
		Clock: &clock,
	}
	screenSettings = screen.Settings{
		RandomGates: &trig,
	}
)

func makeGate(out output.Output) func(high bool) {
	return func(high bool) {
		if high {
			out.On()
		} else {
			out.Off()
		}
	}
}

func startLoop(e *europi.EuroPi) {
	if err := trig.Init(module.Config{
		Trigger: [3]func(high bool){
			makeGate(e.CV1), // Trigger 1
			makeGate(e.CV2), // Trigger 2
			makeGate(e.CV3), // Trigger 3
		},
		Gate: [3]func(high bool){
			makeGate(e.CV4), // Gate 1
			makeGate(e.CV5), // Gate 2
			makeGate(e.CV6), // Gate 3
		},
		Mode: module.Mode1msTrig,
	}); err != nil {
		panic(err)
	}

	if err := clock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  false,
		ClockOut: trig.Clock,
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		trig.Clock(e.DI.Value())
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	clock.Tick(deltaTime)
	trig.Tick(deltaTime)
}

func main() {
	var err error
	ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &screenMain),
		screenbank.WithScreen("settings", "\u2611", &screenSettings),
		screenbank.WithScreen("clock", "\u23f0", &screenClock),
	)
	if err != nil {
		panic(err)
	}

	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
		europi.UI(ui),
		europi.UIRefreshRate(time.Millisecond*50),
	)
}
