package module

import (
	"time"

	cascadelfo "github.com/heucuva/europi/internal/projects/cascadelfo/module"
	clockgenerator "github.com/heucuva/europi/internal/projects/clockgenerator/module"
	randomgates "github.com/heucuva/europi/internal/projects/randomgates/module"
	randomskips "github.com/heucuva/europi/internal/projects/randomskips/module"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module/complexarp"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module/complexenvelope"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module/complexrandom"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module/threephaselfo"
	"github.com/heucuva/europi/units"
)

type UnfoldableSpace struct {
	clock clockgenerator.ClockGenerator
	harm  complexarp.ComplexArp
	lfo   cascadelfo.CascadeLFO
	trig  randomgates.RandomGates
	skip  randomskips.RandomSkips
	env   complexenvelope.ComplexEnvelope
	mod   threephaselfo.ThreePhaseLFO
	rnd   complexrandom.ComplexRandom

	onClock           func(high bool)
	onTrigOutputGate1 func(high bool)
	onSkipSetCV1      func(cv units.CV)
	onSkipOutputGate1 func(high bool)
	onLFOOutput5      func(cv units.CV)
}

func (m *UnfoldableSpace) Init(config Config) error {
	m.onClock = config.OnClock
	m.onTrigOutputGate1 = config.OnTrigOuputGate1
	m.onSkipSetCV1 = config.OnSkipSetCV1
	m.onSkipOutputGate1 = config.OnSkipOutputGate1
	m.onLFOOutput5 = config.OnLFOOutput5

	if err := m.clock.Init(clockgenerator.Config{
		BPM:      120.0,
		Enabled:  false,
		ClockOut: m.trigClock,
	}); err != nil {
		return err
	}

	if err := m.trig.Init(randomgates.Config{
		Gate: [1]func(high bool){
			m.trigOuputGate1, // Gate 1
		},
		Chance:   0.333333,
		Duration: time.Millisecond * 200,
	}); err != nil {
		return err
	}

	if err := m.mod.Init(threephaselfo.Config{
		WaveMode:  threephaselfo.WaveModeSine,
		Phi3Rate:  0.2,
		SkewRate:  0.0,
		SkewShape: 0.05,
		Degree0:   m.skipSetCV1,
	}); err != nil {
		return err
	}

	if err := m.skip.Init(randomskips.Config{
		Gate: [1]func(high bool){
			m.skipOutputGate1, // Gate1
		},
		Chance: 0.6,
	}); err != nil {
		return err
	}

	if err := m.harm.Init(complexarp.Config{
		ArpOut:     config.SetVOct,
		ArpPattern: complexarp.PatternBrownian,
		Scale:      complexarp.ScaleC_Major,
		Quantizer:  complexarp.QuantizerRound,
		ArpRange:   1.0,
		ArpPitch:   4.0,
	}); err != nil {
		return err
	}

	if err := m.lfo.Init(cascadelfo.Config{
		LFO: [8]func(cv units.CV){
			nil,             // LFO 1
			nil,             // LFO 2
			nil,             // LFO 3
			config.SetMorph, // LFO 4
			m.lfoOutput5,    // LFO 5
			nil,             // LFO 6
			nil,             // LFO 7
			nil,             // LFO 8
		},
		Rate:             0.8,
		RateAttenuverter: 0.9,
	}); err != nil {
		panic(err)
	}

	if err := m.env.Init(complexenvelope.Config{
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
					config.SetLFOCV(cv)
					m.lfo.SetCV(cv)
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

	if err := m.rnd.Init(complexrandom.Config{
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

func (m *UnfoldableSpace) Clock() {
	m.trigClock(true)
}

func (m *UnfoldableSpace) EnableInternalClock(enabled bool) {
	m.clock.SetEnabled(enabled)
}

func (m *UnfoldableSpace) InternalClockEnabled() bool {
	return m.clock.Enabled()
}

func (m *UnfoldableSpace) ToggleInternalClock() {
	m.clock.Toggle()
}

func (m *UnfoldableSpace) Tick(deltaTime time.Duration) {
	m.clock.Tick(deltaTime)
	m.harm.Tick(deltaTime)
	m.lfo.Tick(deltaTime)
	m.trig.Tick(deltaTime)
	m.skip.Tick(deltaTime)
	m.env.Tick(deltaTime)
	m.mod.Tick(deltaTime)
	m.rnd.Tick(deltaTime)
}

func (m *UnfoldableSpace) trigClock(high bool) {
	if m.onClock != nil {
		m.onClock(high)
	}
	m.trig.Clock(high)
}

func (m *UnfoldableSpace) trigOuputGate1(high bool) {
	if m.onTrigOutputGate1 != nil {
		m.onTrigOutputGate1(high)
	}
	m.skip.Gate(0, high)
}

func (m *UnfoldableSpace) skipSetCV1(cv units.CV) {
	if m.onSkipSetCV1 != nil {
		m.onSkipSetCV1(cv)
	}
	m.skip.SetCV(0, cv)
}

func (m *UnfoldableSpace) skipOutputGate1(high bool) {
	if m.onSkipOutputGate1 != nil {
		m.onSkipOutputGate1(high)
	}
	m.env.Gate(0, high)
	m.env.Gate(1, high)
	m.harm.ArpClock(high)
}

func (m *UnfoldableSpace) lfoOutput5(cv units.CV) {
	if m.onLFOOutput5 != nil {
		m.onLFOOutput5(cv)
	}
	m.env.SetCV(0, cv)
}
