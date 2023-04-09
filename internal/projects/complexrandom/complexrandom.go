package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/complexrandom/module"
)

var (
	rnd module.ComplexRandom
)

func startLoop(e *europi.EuroPi) {
	if err := rnd.Init(module.Config{
		SampleOutA:        e.CV1.SetCV,
		SampleOutB:        e.CV2.SetCV,
		SampleAttenuatorA: 0.6,
		IntegrationSlope:  0.0,
		GateDensity:       0.4,
		PulseStageDivider: 1.0,
		SampleAttenuatorB: 0.2,
		SampleSlewB:       0.3,
		ClockSpeed:        0.4,
		ClockRange:        module.ClockFull,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		high := e.DI.Value()
		rnd.Gate(high)
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
	rnd.SetClockRate(e.K1.ReadCV())
	rnd.SetSlewB(e.K2.ReadCV())
	rnd.SetSample(e.AI.ReadCV())
	rnd.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		disp.WriteLine(fmt.Sprintf("Clk:%2.1f Slw:%2.1f", rnd.ClockRate(), rnd.SlewB()), 0, line1y)
		disp.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f", e.CV1.Voltage(), e.CV2.Voltage()), 0, line2y)
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
