package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/complexenvelope/module"
	"github.com/heucuva/europi/internal/projects/complexenvelope/screen"
)

var (
	env        module.ComplexEnvelope
	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		Env: &env,
	}
	screenSettings = screen.Settings{
		Env: &env,
	}
)

func startLoop(e *europi.EuroPi) {
	if err := env.Init(module.Config{
		Env: [2]module.EnvelopeConfig{
			{ // 1
				Out:         e.CV1.SetCV,
				Mode:        module.EnvelopeModeAD,
				AttackMode:  module.FunctionModeLinear,
				ReleaseMode: module.FunctionModeExponential,
				Attack:      0.5,
				Decay:       0.5,
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

}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	env.Tick(deltaTime)
}

func main() {
	var err error
	ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &screenMain),
		screenbank.WithScreen("settings", "\u2611", &screenSettings),
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
