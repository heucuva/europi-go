package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/internal/projects/complexenvelope/module"
	"github.com/awonak/EuroPiGo/internal/projects/complexenvelope/screen"
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

	e.DI.HandlerEx(hal.ChangeAny, func(value bool, _ time.Duration) {
		env.Gate(0, value)
		env.Gate(1, value)
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
