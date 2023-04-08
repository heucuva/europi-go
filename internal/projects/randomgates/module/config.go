package module

import (
	"time"
)

type Config struct {
	Gate     [1]func(high bool)
	Chance   float32
	Duration time.Duration
}
