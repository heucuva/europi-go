package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/internal/projects/randomgates/module"
	"github.com/heucuva/europi/output"
)

var (
	trig  module.RandomGates
	clock clockgenerator.ClockGenerator
)

func startLoop(e *europi.EuroPi) {
	if err := trig.Init(module.Config{
		Gate: [1]func(high bool){
			func(high bool) { // Gate 1
				if high {
					e.CV1.On()
				} else {
					e.CV1.Off()
				}
			},
		},
		Chance:   0.333333,
		Duration: time.Millisecond * 200,
	}); err != nil {
		panic(err)
	}

	if err := clock.Init(clockgenerator.Config{
		BPM:     120.0,
		Enabled: false,
		ClockOut: func(high bool) {
			if high {
				e.CV1.On()
			} else {
				e.CV1.Off()
			}
		},
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		trig.Clock(true)
	})

	e.B1.HandlerWithDebounce(func(p machine.Pin) {
		clock.Toggle()
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
	clock.Tick(deltaTime)
	trig.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		if clock.Enabled() {
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
