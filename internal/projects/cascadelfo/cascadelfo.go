package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/cascadelfo/module"
	"github.com/heucuva/europi/units"
)

var (
	lfo module.CascadeLFO
)

func startLoop(e *europi.EuroPi) {
	if err := lfo.Init(module.Config{
		LFO: [8]func(cv units.CV){
			e.CV1.SetCV, // LFO 1
			e.CV2.SetCV, // LFO 2
			e.CV3.SetCV, // LFO 3
			e.CV4.SetCV, // LFO 4
			e.CV5.SetCV, // LFO 5
			e.CV6.SetCV, // LFO 6
			nil,         // LFO 7
			nil,         // LFO 8
		},
		Rate:             0.8,
		RateAttenuverter: 0.9,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		lfo.Reset()
	})
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
	lfo.SetAttenuverter(e.K2.ReadCV())
	lfo.SetRate(e.K1.ReadCV())
	lfo.SetCV(e.AI.ReadCV())
	lfo.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
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
