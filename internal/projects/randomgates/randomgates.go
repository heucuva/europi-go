package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	clockgenerator "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	clockScreen "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/screen"
	"github.com/awonak/EuroPiGo/internal/projects/randomgates/module"
	"github.com/awonak/EuroPiGo/internal/projects/randomgates/screen"
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

func makeGate(out hal.VoltageOutput) func(high bool) {
	return func(high bool) {
		if high {
			out.SetCV(1.0)
		} else {
			out.SetCV(0.0)
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

	e.DI.HandlerEx(hal.ChangeAny, func(value bool, _ time.Duration) {
		trig.Clock(value)
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
