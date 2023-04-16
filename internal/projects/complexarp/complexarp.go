package main

import (
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/experimental/screenbank"
	"github.com/heucuva/europi/internal/projects/complexarp/module"
	"github.com/heucuva/europi/internal/projects/complexarp/screen"
)

var (
	arp        module.ComplexArp
	ui         *screenbank.ScreenBank
	screenMain = screen.Main{
		ComplexArp: &arp,
	}
	screenSettings = screen.Settings{
		ComplexArp: &arp,
	}
)

func startLoop(e *europi.EuroPi) {
	if err := arp.Init(module.Config{
		ArpOut:     e.CV1.SetVOct,
		ArpPattern: module.PatternBrownian,
		Scale:      module.ScaleC_Major,
		Quantizer:  quantizer.ModeRound,
		ArpPitch:   4.0,
		ArpRange:   1.0,
	}); err != nil {
		panic(err)
	}

	e.DI.Handler(func(p machine.Pin) {
		high := e.DI.Value()
		arp.ArpClock(high)
	})
}

func mainLoop(e *europi.EuroPi, deltaTime time.Duration) {
	arp.Tick(deltaTime)
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
