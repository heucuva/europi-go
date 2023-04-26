package module

import (
	"time"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

type clockBasic struct {
	spectrum time.Duration
	interval time.Duration
	rate     units.CV

	dur time.Duration
}

func (c *clockBasic) SetRate(cv units.CV) {
	c.rate = cv
	l := lerp.NewLerp32(time.Second, time.Second/c.spectrum)
	c.interval = l.ClampedLerp(c.rate.ToFloat32())
}

func (c *clockBasic) Rate() units.CV {
	return c.rate
}

func (c *clockBasic) Tick(deltaTime time.Duration) bool {
	t := c.dur + deltaTime
	triggered := t >= c.interval
	c.dur = t % c.interval

	return triggered
}
