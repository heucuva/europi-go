package module

import (
	"time"

	europim "github.com/heucuva/europi/math"
	"github.com/heucuva/europi/units"
)

type clockBasic struct {
	spectrum time.Duration
	interval time.Duration
	rate     units.CV

	dur time.Duration
}

func (c *clockBasic) SetRate(cv units.CV) {
	c.rate = cv
	c.interval = europim.Lerp(c.rate.ToFloat32(), time.Second, time.Second/c.spectrum)
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
