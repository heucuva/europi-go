package complexenvelope

import (
	"time"

	"github.com/heucuva/europi/units"
)

type modeFunc func(t, dur time.Duration) units.CV
