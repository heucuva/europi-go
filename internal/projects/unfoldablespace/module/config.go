package module

import "github.com/awonak/EuroPiGo/units"

type Config struct {
	SetVOct   func(voct units.VOct)
	SetLevel  func(cv units.CV)
	SetTimbre func(cv units.CV)
	SetHarmo  func(cv units.CV)
	SetMorph  func(cv units.BipolarCV)
	SetLFOCV  func(cv units.BipolarCV)

	VOctInputEnabled bool

	OnClock           func(high bool)
	OnTrigOuputGate1  func(high bool)
	OnSkipSetCV1      func(cv units.CV)
	OnSkipOutputGate1 func(high bool)
	OnLFOOutput5      func(cv units.BipolarCV)
}