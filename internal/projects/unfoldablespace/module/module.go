package module

import (
	"time"

	cascadelfo "github.com/awonak/EuroPiGo/internal/projects/cascadelfo/module"
	clockgenerator "github.com/awonak/EuroPiGo/internal/projects/clockgenerator/module"
	complexarp "github.com/awonak/EuroPiGo/internal/projects/complexarp/module"
	complexenvelope "github.com/awonak/EuroPiGo/internal/projects/complexenvelope/module"
	complexrandom "github.com/awonak/EuroPiGo/internal/projects/complexrandom/module"
	randomgates "github.com/awonak/EuroPiGo/internal/projects/randomgates/module"
	randomskips "github.com/awonak/EuroPiGo/internal/projects/randomskips/module"
	threephaselfo "github.com/awonak/EuroPiGo/internal/projects/threephaselfo/module"
	"github.com/awonak/EuroPiGo/quantizer"
	"github.com/awonak/EuroPiGo/units"
)

type UnfoldableSpace struct {
	ModClock clockgenerator.ClockGenerator
	ModArp   complexarp.ComplexArp
	ModLFO   cascadelfo.CascadeLFO
	ModTrig  randomgates.RandomGates
	ModSkip  randomskips.RandomSkips
	ModEnv   complexenvelope.ComplexEnvelope
	ModMod   threephaselfo.ThreePhaseLFO
	ModRnd   complexrandom.ComplexRandom

	voctInputEnabled bool

	onClock           func(high bool)
	onTrigOutputGate1 func(high bool)
	onSkipSetCV1      func(cv units.CV)
	onSkipOutputGate1 func(high bool)
	onLFOOutput5      func(cv units.BipolarCV)
}

func (m *UnfoldableSpace) Init(config Config) error {
	m.onClock = config.OnClock
	m.onTrigOutputGate1 = config.OnTrigOuputGate1
	m.onSkipSetCV1 = config.OnSkipSetCV1
	m.onSkipOutputGate1 = config.OnSkipOutputGate1
	m.onLFOOutput5 = config.OnLFOOutput5

	m.voctInputEnabled = config.VOctInputEnabled

	if err := m.ModClock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  false,
		ClockOut: m.trigClock,
	}); err != nil {
		return err
	}

	if err := m.ModTrig.Init(randomgates.Config{
		Trigger: [3]func(high bool){
			m.trigOuputGate1, // Gate 1
		},
		Mode: randomgates.Mode200msTrig,
	}); err != nil {
		return err
	}

	if err := m.ModMod.Init(threephaselfo.Config{
		WaveMode:  threephaselfo.WaveModeSine,
		Phi3Rate:  threephaselfo.CVToPhi3Rate(0.2),
		SkewRate:  threephaselfo.CVToSkewRate(0.0),
		SkewShape: 0.05,
		Degree0: func(cv units.BipolarCV) {
			v, _ := cv.ToCV()
			m.skipSetCV1(v)
		},
	}); err != nil {
		return err
	}

	if err := m.ModSkip.Init(randomskips.Config{
		Gate:   m.skipOutputGate1,
		Chance: 0.6,
	}); err != nil {
		return err
	}

	if err := m.ModArp.Init(complexarp.Config{
		ArpOut:     config.SetVOct,
		ArpPattern: complexarp.PatternBrownian,
		Scale:      complexarp.ScaleC_Major,
		Quantizer:  quantizer.ModeRound,
		ArpRange:   1.0,
		ArpPitch:   4.0,
	}); err != nil {
		return err
	}

	if err := m.ModLFO.Init(cascadelfo.Config{
		LFO: [8]func(cv units.BipolarCV){
			nil,             // LFO 1
			nil,             // LFO 2
			nil,             // LFO 3
			config.SetMorph, // LFO 4
			m.lfoOutput5,    // LFO 5
			nil,             // LFO 6
			nil,             // LFO 7
			nil,             // LFO 8
		},
		Rate:             cascadelfo.CVToRate(0.8),
		RateAttenuverter: 0.8, // +80%
	}); err != nil {
		panic(err)
	}

	if err := m.ModEnv.Init(complexenvelope.Config{
		Env: [2]complexenvelope.EnvelopeConfig{
			{ // 1
				Out:         config.SetLevel,
				Mode:        complexenvelope.EnvelopeModeAD,
				AttackMode:  complexenvelope.FunctionModeLinear,
				ReleaseMode: complexenvelope.FunctionModeExponential,
				Attack:      0.6666666666666667,
				Decay:       0.6666666666666667,
			},
			{ // 2
				Out: func(cv units.CV) {
					bcv := cv.ToBipolarCV(1)
					config.SetLFOCV(bcv)
					m.ModLFO.SetCV(bcv)
				},
				Mode:        complexenvelope.EnvelopeModeAD,
				AttackMode:  complexenvelope.FunctionModeLinear,
				ReleaseMode: complexenvelope.FunctionModeExponential,
				Attack:      0.5,
				Decay:       0.5,
			},
		},
	}); err != nil {
		return err
	}

	if err := m.ModRnd.Init(complexrandom.Config{
		SampleOutA:        config.SetTimbre,
		SampleOutB:        config.SetHarmo,
		SampleAttenuatorA: 0.6,
		IntegrationSlope:  0.0,
		GateDensity:       0.4,
		PulseStageDivider: 1.0,
		SampleAttenuatorB: 0.2,
		SampleSlewB:       0.3,
		ClockSpeed:        0.4,
		ClockRange:        complexrandom.ClockFull,
	}); err != nil {
		return err
	}

	return nil
}

func (m *UnfoldableSpace) Clock(high bool) {
	m.trigClock(high)
}

func (m *UnfoldableSpace) EnableInternalClock(enabled bool) {
	m.ModClock.SetEnabled(enabled)
}

func (m *UnfoldableSpace) ToggleInternalClock() {
	m.ModClock.Toggle()
}

func (m *UnfoldableSpace) InternalClockEnabled() bool {
	return m.ModClock.Enabled()
}

func (m *UnfoldableSpace) ToggleVOctInputEnabled() {
	m.voctInputEnabled = !m.voctInputEnabled
}

func (m *UnfoldableSpace) VOctInputEnabled() bool {
	return m.voctInputEnabled
}

func (m *UnfoldableSpace) SetVOct(voct units.VOct) {
	if m.voctInputEnabled {
		m.ModArp.SetArpPitch(voct)
	}
}

func (m *UnfoldableSpace) Tick(deltaTime time.Duration) {
	m.ModClock.Tick(deltaTime)
	m.ModArp.Tick(deltaTime)
	m.ModLFO.Tick(deltaTime)
	m.ModTrig.Tick(deltaTime)
	m.ModSkip.Tick(deltaTime)
	m.ModEnv.Tick(deltaTime)
	m.ModMod.Tick(deltaTime)
	m.ModRnd.Tick(deltaTime)
}

func (m *UnfoldableSpace) trigClock(high bool) {
	if m.onClock != nil {
		m.onClock(high)
	}
	m.ModTrig.Clock(high)
}

func (m *UnfoldableSpace) trigOuputGate1(high bool) {
	if m.onTrigOutputGate1 != nil {
		m.onTrigOutputGate1(high)
	}
	m.ModSkip.Gate(high)
}

func (m *UnfoldableSpace) skipSetCV1(cv units.CV) {
	if m.onSkipSetCV1 != nil {
		m.onSkipSetCV1(cv)
	}
	m.ModSkip.SetCV(cv)
}

func (m *UnfoldableSpace) skipOutputGate1(high bool) {
	if m.onSkipOutputGate1 != nil {
		m.onSkipOutputGate1(high)
	}
	m.ModEnv.Gate(0, high)
	m.ModEnv.Gate(1, high)
	m.ModArp.ArpClock(high)
}

func (m *UnfoldableSpace) lfoOutput5(cv units.BipolarCV) {
	if m.onLFOOutput5 != nil {
		m.onLFOOutput5(cv)
	}
	m.ModEnv.SetCV(0, cv)
}
