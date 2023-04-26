package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/internal/projects/complexrandom/module"
	"github.com/awonak/EuroPiGo/internal/projects/complexrandom/screen"
)

var (
	rnd        module.ComplexRandom
	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		ComplexRandom: &rnd,
	}
	screenSettings = screen.Settings{
		ComplexRandom: &rnd,
	}
)

func startLoop(e *europi.EuroPi) {
	if err := rnd.Init(module.Config{
		SampleOutA:        e.CV1.SetBipolarCV,
		SampleOutB:        e.CV2.SetBipolarCV,
		SampleAttenuatorA: 0.6,
		IntegrationSlope:  0.0,
		GateDensity:       0.4,
		PulseStageDivider: 1,
		SampleAttenuatorB: 0.2,
		SampleSlewB:       0.3,
		ClockSpeed:        0.4,
		ClockRange:        module.ClockFull,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(value bool, _ time.Duration) {
		rnd.Gate(value)
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	rnd.Tick(deltaTime)
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
