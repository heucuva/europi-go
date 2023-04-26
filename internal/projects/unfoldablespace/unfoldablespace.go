package main

import (
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/screenbank"
	"github.com/awonak/EuroPiGo/hardware/hal"
	lfoScreen "github.com/awonak/EuroPiGo/internal/projects/cascadelfo/screen"
	clockScreen "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/screen"
	arpScreen "github.com/awonak/EuroPiGo/internal/projects/complexarp/screen"
	envScreen "github.com/awonak/EuroPiGo/internal/projects/complexenvelope/screen"
	rndScreen "github.com/awonak/EuroPiGo/internal/projects/complexrandom/screen"
	trigScreen "github.com/awonak/EuroPiGo/internal/projects/randomgates/screen"
	skipScreen "github.com/awonak/EuroPiGo/internal/projects/randomskips/screen"
	modScreen "github.com/awonak/EuroPiGo/internal/projects/threephaselfo/screen"
	"github.com/awonak/EuroPiGo/internal/projects/unfoldablespace/module"
	"github.com/awonak/EuroPiGo/internal/projects/unfoldablespace/screen"
	"github.com/awonak/EuroPiGo/units"
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
	screenEnv = envScreen.Settings{
		Env: &unfold.ModEnv,
	}
	screenRnd = rndScreen.Settings{
		ComplexRandom: &unfold.ModRnd,
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
	setMorph := func(cv units.BipolarCV) {
		e.CV5.SetVoltage(cv.ToVolts())
	}
	setLFOCV := func(cv units.BipolarCV) {
		e.CV6.SetVoltage(cv.ToVolts())
	}

	if err := unfold.Init(module.Config{
		SetVOct:          setVOct,
		SetLevel:         setLevel,
		SetTimbre:        setTimbre,
		SetHarmo:         setHarmo,
		SetMorph:         setMorph,
		SetLFOCV:         setLFOCV,
		VOctInputEnabled: false,
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(hal.ChangeAny, func(value bool, _ time.Duration) {
		unfold.Clock(value)
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	unfold.SetVOct(e.AI.ReadVOct())
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
		screenbank.WithScreen("env", "\u2709", &screenEnv),
		screenbank.WithScreen("rnd", "\u2744", &screenRnd),
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
