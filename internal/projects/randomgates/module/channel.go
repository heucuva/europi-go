package module

import (
	"math/rand"
	"time"
)

const (
	gateChance float32 = 1.0 / 3.0
)

type channel struct {
	trig    func(high bool)
	trigRem time.Duration
	gate    func(high bool)
	gateRem time.Duration
}

func (c *channel) Start(clockDur time.Duration, mode Mode) {
	if c.gateRem > 0 || rand.Float32() >= gateChance {
		// disallow updates while in an active gate period or we fail our chance
		return
	}

	if c.trigRem > 0 {
		// we got a retrigger input, so force it out
		c.trig(false)
	}

	if c.gateRem > 0 {
		// we got a retrigger input, so force it out
		c.gate(false)
	}

	c.trigRem = mode.GetTriggerDuration(clockDur)
	c.trig(true)

	c.gateRem = clockDur
	c.gate(true)
}

func (c *channel) Tick(deltaTime time.Duration) {
	if c.trigRem > 0 {
		c.trigRem -= deltaTime
		if c.trigRem <= 0 {
			c.trigRem = 0
			c.trig(false)
		}
	}

	if c.gateRem > 0 {
		c.gateRem -= deltaTime
		if c.gateRem <= 0 {
			c.gateRem = 0
			c.gate(false)
		}
	}
}
