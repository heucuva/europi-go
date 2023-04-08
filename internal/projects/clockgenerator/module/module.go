package module

import (
	"fmt"
	"time"
)

type ClockGenerator struct {
	interval     time.Duration
	gateDuration time.Duration
	enabled      bool
	out          func(high bool)
	t            time.Duration
	gateT        time.Duration
	gateLevel    bool

	bpm float32 // informational
}

func noop(_ bool) {
}

func (m *ClockGenerator) Init(config Config) error {
	if config.BPM <= 0 {
		return fmt.Errorf("invalid bpm setting: %v", config.BPM)
	}

	m.enabled = config.Enabled
	m.gateDuration = config.GateDuration
	if m.gateDuration == 0 {
		m.gateDuration = DefaultGateDuration
	}

	m.out = config.ClockOut
	if m.out == nil {
		m.out = noop
	}

	m.SetBPM(config.BPM)
	return nil
}

func (m *ClockGenerator) Toggle() {
	m.enabled = !m.enabled
	m.t = 0
}

func (m *ClockGenerator) SetEnabled(enabled bool) {
	m.enabled = enabled
	m.t = 0
}

func (m *ClockGenerator) Enabled() bool {
	return m.enabled
}

func (m *ClockGenerator) BPM() float32 {
	return m.bpm
}

func (m *ClockGenerator) SetBPM(bpm float32) {
	if bpm == 0 {
		bpm = 120.0
	}
	m.bpm = bpm
	m.interval = time.Duration(float32(time.Minute) / bpm)
}

func (m *ClockGenerator) Tick(deltaTime time.Duration) {
	if !m.enabled {
		return
	}

	prevGateLevel := m.gateLevel

	var reset bool
	deltaTime, reset = m.processClockInterval(deltaTime)

	if reset {
		m.gateT = 0
		m.gateLevel = true
	}

	gateT := m.gateT + deltaTime
	m.gateT = gateT % m.gateDuration
	if gateT >= m.gateDuration {
		m.gateLevel = false
	}

	if m.gateLevel != prevGateLevel {
		m.out(m.gateLevel)
	}
}

func (m *ClockGenerator) processClockInterval(deltaTime time.Duration) (time.Duration, bool) {
	t := m.t + deltaTime
	m.t = t % m.interval

	if t >= m.interval {
		return m.t, true
	}

	return deltaTime, false
}
