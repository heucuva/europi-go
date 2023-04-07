package threephaselfo

import (
	"time"

	"github.com/heucuva/europi/units"
)

type wave interface {
	Get(t, interval time.Duration) units.CV
}
