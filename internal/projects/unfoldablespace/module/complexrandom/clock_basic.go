package complexrandom

import (
	"time"
)

type clockBasic struct {
	interval time.Duration

	dur time.Duration
}

func (c *clockBasic) Tick(deltaTime time.Duration) bool {
	t := c.dur + deltaTime
	triggered := t >= c.interval
	c.dur = t % c.interval

	return triggered
}
