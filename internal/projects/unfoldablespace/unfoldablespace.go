package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/screenbank"
	lfoScreen "github.com/heucuva/europi/internal/projects/cascadelfo/screen"
	clockScreen "github.com/heucuva/europi/internal/projects/clockgenerator/screen"
	arpScreen "github.com/heucuva/europi/internal/projects/complexarp/screen"
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
		Clock: &unfold.ModClock,
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
	screenArp = arpScreen.Settings{
		ComplexArp: &unfold.ModArp,
	}
	screenLFO = lfoScreen.Settings{
		LFO: &unfold.ModLFO,
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

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.DI.Value()
		unfold.Clock(high)
	})
}

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
		screenbank.WithScreen("arp", "\u26bd", &screenArp),
		screenbank.WithScreen("lfo", "\u2797", &screenLFO),
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
