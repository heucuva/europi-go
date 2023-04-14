package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/cascadelfo/module"
	"github.com/heucuva/europi/internal/projects/cascadelfo/screen"
	"github.com/heucuva/europi/units"
)

var (
	lfo        module.CascadeLFO
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
		Rate:             16.0,
		RateAttenuverter: module.CVToRateAV(0.9),
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		lfo.Reset()
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	lfo.SetCV(e.AI.ReadCV())
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
