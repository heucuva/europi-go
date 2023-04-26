package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/internal/projects/cascadelfo/module"
	"github.com/awonak/EuroPiGo/internal/projects/cascadelfo/screen"
	"github.com/awonak/EuroPiGo/units"
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

func bipolarOut(out func(units.CV)) func(cv units.BipolarCV) {
	return func(cv units.BipolarCV) {
		v, _ := cv.ToCV()
		out(v)
	}
}

func startLoop(e *europi.EuroPi) {
	if err := lfo.Init(module.Config{
		LFO: [8]func(cv units.BipolarCV){
			bipolarOut(e.CV1.SetCV), // LFO 1
			bipolarOut(e.CV2.SetCV), // LFO 2
			bipolarOut(e.CV3.SetCV), // LFO 3
			bipolarOut(e.CV4.SetCV), // LFO 4
			bipolarOut(e.CV5.SetCV), // LFO 5
			bipolarOut(e.CV6.SetCV), // LFO 6
			nil,                     // LFO 7
			nil,                     // LFO 8
		},
		Rate:             16.0, // Hz
		RateAttenuverter: 0.8,  // +80%
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
	lfo.SetCV(e.AI.ReadBipolarCV())
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