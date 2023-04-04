package randomgates

import (
	"time"
)

type Config struct {
	Gate     [1]func(high bool)
	Duration time.Duration
}
