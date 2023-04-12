package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/quantizer"
	"github.com/heucuva/europi/internal/projects/complexarp/module"
)

var (
	arp module.ComplexArp
)

func startLoop(e *europi.EuroPi) {
	if err := arp.Init(module.Config{
		ArpOut:     e.CV1.SetVOct,
		ArpPattern: module.PatternBrownian,
		Scale:      module.ScaleC_Major,
		Quantizer:  quantizer.ModeRound,
		ArpRange:   1.0,
		ArpPitch:   4.0,
	}); err != nil {
		panic(err)
	}

	e.DI.HandlerEx(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.DI.Value()
		arp.ArpClock(high)
	})

	e.B1.HandlerExWithDebounce(machine.PinRising|machine.PinFalling, func(p machine.Pin) {
		high := e.B1.Value()
		arp.ArpClock(high)
	}, time.Millisecond*500)
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
	arp.SetArpPitch(e.AI.ReadVOct())
	arp.SetArpRange(e.K2.ReadVOct())
	arp.SetScaleCV(e.K1.ReadCV())
	arp.Tick(deltaTime)

	displayDelay += deltaTime
	if displayDelay > displayRate {
		displayDelay %= displayRate

		disp := e.Display
		disp.ClearBuffer()
		disp.WriteLine(fmt.Sprintf("Scl:%5s Rng:%2.1f", arp.ScaleName(), arp.ArpRange()), 0, line1y)
		disp.WriteLine(fmt.Sprintf("1:%2.1f", e.CV1.Voltage()), 0, line2y)
		disp.Display()
	}
}

func main() {
	// some options shown below are being explicitly set to their defaults
	// only to showcase their existence.
	europi.Bootstrap(
		europi.EnableDisplayLogger(false),
		europi.InitRandom(true),
		europi.StartLoop(startLoop),
		europi.MainLoop(mainLoop),
		europi.MainLoopInterval(time.Millisecond*1),
	)
}
