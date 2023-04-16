package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/complexrandom/module"
	"github.com/heucuva/europi/internal/projects/complexrandom/screen"
	"github.com/heucuva/europi/units"
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

func bipolarOut(out func(units.CV)) func(cv units.BipolarCV) {
	return func(cv units.BipolarCV) {
		out(cv.ToCV())
	}
}

func startLoop(e *europi.EuroPi) {
	if err := rnd.Init(module.Config{
		SampleOutA:        e.CV1.SetCV,
		SampleOutB:        e.CV2.SetCV,
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

	e.DI.Handler(func(p machine.Pin) {
		high := e.DI.Value()
		rnd.Gate(high)
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
