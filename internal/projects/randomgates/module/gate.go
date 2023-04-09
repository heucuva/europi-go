package module

import "time"

type gate struct {
	out   func(high bool)
	level bool
	rem   time.Duration
}
