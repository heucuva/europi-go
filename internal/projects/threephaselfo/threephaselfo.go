package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/threephaselfo/module"
	"github.com/heucuva/europi/internal/projects/threephaselfo/screen"
	"github.com/heucuva/europi/units"
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
		out(cv.ToCV())
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

	e.DI.Handler(func(p machine.Pin) {
		lfo.Reset()
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
