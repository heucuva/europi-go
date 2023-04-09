package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/internal/projects/randomskips/module"
	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/output"
)

var (
	skip  module.RandomSkips
	clock clockgenerator.ClockGenerator
)

func startLoop(e *europi.EuroPi) {
	if err := skip.Init(module.Config{
		Gate: [1]func(high bool){
			func(high bool) { // Gate 1
				if high {
					e.CV1.On()
				} else {
					e.CV1.Off()
				}
			},
		},
		Chance: 0.333333,
	}); err != nil {
		panic(err)
	}

	if err := clock.Init(clockgenerator.Config{
		BPM:     120.0,
		Enabled: false,
		ClockOut: func(high bool) {
			skip.Gate(0, high)
		},
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.DI.Value()
		skip.Gate(0, high)
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
	skip.SetChance(e.K1.ReadCV().ToFloat32())
	cv := e.K2.ReadCV()
	clock.SetBPM(europim.Lerp[float32](cv.ToFloat32(), 0.01, 240.0))
	clock.Tick(deltaTime)
	skip.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		if clock.Enabled() {
			disp.DrawHLine(0, 0, 7, output.White)
			disp.WriteLine(fmt.Sprintf("BPM:%3.1f", clock.BPM()), 64, line1y)
		}
		disp.WriteLine(fmt.Sprintf("Chn:%3.1f%%", skip.Chance()*100.0), 0, line1y)
		disp.WriteLine(fmt.Sprintf("1:%2.1f", e.CV1.Voltage()), 0, line2y)
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
