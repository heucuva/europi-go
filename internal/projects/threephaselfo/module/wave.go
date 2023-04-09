package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type wave interface {
	Get(t, interval time.Duration) (units.CV, units.CV, units.CV)
	Mode() WaveMode
}