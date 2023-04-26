package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/internal/projects/threephaselfo/module"
	"github.com/awonak/EuroPiGo/internal/projects/threephaselfo/screen"
	"github.com/awonak/EuroPiGo/units"
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

func bipolarOut(out func(units.CV)) func(cv units.BipolarCV) {
	return func(cv units.BipolarCV) {
		v, _ := cv.ToCV()
		out(v)
	}
}

func startLoop(e *europi.EuroPi) {
	if err := lfo.Init(module.Config{
		Degree0:   bipolarOut(e.CV1.SetCV),
		Degree120: bipolarOut(e.CV2.SetCV),
		Degree240: bipolarOut(e.CV3.SetCV),
		WaveMode:  module.WaveModeSine,
		Phi3Rate:  1.0,
		SkewRate:  20.0,
		SkewShape: 0.05,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(value bool, _ time.Duration) {
		if value {
			lfo.Reset()
		}
	})
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
