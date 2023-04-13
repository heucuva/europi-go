package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	clockScreen "github.com/heucuva/europi/internal/projects/clockgenerator/screen"
	trigScreen "github.com/heucuva/europi/internal/projects/randomgates/screen"
	skipScreen "github.com/heucuva/europi/internal/projects/randomskips/screen"
	modScreen "github.com/heucuva/europi/internal/projects/threephaselfo/screen"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/screen"
	"github.com/heucuva/europi/units"
)

var (
	unfold     module.UnfoldableSpace
	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		Unfold: &unfold,
	}
	screenClock = clockScreen.Settings{
		Clock:           &unfold.ModClock,
		MinBPM:          0.01,
		MaxBPM:          240.0,
		MinGateDuration: time.Millisecond * 1,
		MaxGateDuration: time.Millisecond * 990,
	}
	screenTrig = trigScreen.Settings{
		RandomGates: &unfold.ModTrig,
	}
	screenMod = modScreen.Settings{
		LFO: &unfold.ModMod,
	}
	screenSkip = skipScreen.Settings{
		RandomSkips: &unfold.ModSkip,
	}
)

func startLoop(e *europi.EuroPi) {
	setVOct := func(voct units.VOct) {
		e.CV1.SetVoltage(voct.ToVolts())
	}
	setLevel := func(cv units.CV) {
		e.CV2.SetVoltage(cv.ToVolts())
	}
	setTimbre := func(cv units.CV) {
		e.CV3.SetVoltage(cv.ToVolts())
	}
	setHarmo := func(cv units.CV) {
		e.CV4.SetVoltage(cv.ToVolts())
	}
	setMorph := func(cv units.CV) {
		e.CV5.SetVoltage(cv.ToVolts())
	}
	setLFOCV := func(cv units.CV) {
		e.CV6.SetVoltage(cv.ToVolts())
	}

	if err := unfold.Init(module.Config{
		SetVOct:   setVOct,
		SetLevel:  setLevel,
		SetTimbre: setTimbre,
		SetHarmo:  setHarmo,
		SetMorph:  setMorph,
		SetLFOCV:  setLFOCV,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		unfold.Clock()
	})
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
	unfold.Tick(deltaTime)
}

func main() {
	var err error
	ui, err = screenbank.NewScreenBank(
		screenbank.WithScreen("main", "\u2b50", &screenMain),
		screenbank.WithScreen("clock", "\u23f0", &screenClock),
		screenbank.WithScreen("trig", "\u303d", &screenTrig),
		screenbank.WithScreen("mod", "\u27bf", &screenMod),
		screenbank.WithScreen("skip", "\u2702", &screenSkip),
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
