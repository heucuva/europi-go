package main

import (
	"fmt"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	europim "github.com/heucuva/europi/math"
)

var (
	clock module.ClockGenerator
)

func startLoop(e *europi.EuroPi) {
	if err := clock.Init(module.Config{
		BPM:     120.0,
		Enabled: true,
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
	cv := e.K1.ReadCV()
	clock.SetBPM(europim.Lerp[float32](cv.ToFloat32(), 0.01, 240.0))
	clock.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		disp.WriteLine(fmt.Sprintf("BPM:%3.1f", clock.BPM()), 0, line1y)
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
