package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/complexenvelope/module"
)

var (
	env module.ComplexEnvelope
)

func startLoop(e *europi.EuroPi) {
	if err := env.Init(module.Config{
		Env: [2]module.EnvelopeConfig{
			{ // 1
				Out:         e.CV1.SetCV,
				Mode:        module.EnvelopeModeAD,
				AttackMode:  module.FunctionModeLinear,
				ReleaseMode: module.FunctionModeExponential,
				Attack:      0.6666666666666667,
				Decay:       0.6666666666666667,
			},
			{ // 2
				Out:         e.CV2.SetCV,
				Mode:        module.EnvelopeModeAD,
				AttackMode:  module.FunctionModeLinear,
				ReleaseMode: module.FunctionModeExponential,
				Attack:      0.5,
				Decay:       0.5,
			},
		},
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.DI.Value()
		env.Gate(0, high)
		env.Gate(1, high)
	})

	e.B1.HandlerExWithDebounce(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.B1.Value()
		env.Gate(0, high)
	}, time.Millisecond*500)

	e.B2.HandlerExWithDebounce(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.B2.Value()
		env.Gate(1, high)
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
	env.SetCV(0, e.K1.ReadCV())
	env.SetCV(1, e.K2.ReadCV())
	env.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		disp.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f", e.CV1.Voltage(), e.CV2.Voltage()), 0, line1y)
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
