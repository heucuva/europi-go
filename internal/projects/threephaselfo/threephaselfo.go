package main

import (
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/threephaselfo/module"
	"github.com/heucuva/europi/internal/projects/threephaselfo/screen"
)

var (
	lfo module.ThreePhaseLFO

	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		LFO: &lfo,
	}
	screenSettings = screen.Settings{
		LFO: &lfo,
	}
)

func startLoop(e *europi.EuroPi) {
	if err := lfo.Init(module.Config{
		WaveMode:  module.WaveModeSine,
		Phi3Rate:  0.2,
		SkewRate:  0.0,
		SkewShape: 0.05,
		Degree0:   e.CV1.SetCV,
		Degree120: e.CV2.SetCV,
		Degree240: e.CV3.SetCV,
	}); err != nil {
		panic(err)
	}
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	lfo.Tick(deltaTime)
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
