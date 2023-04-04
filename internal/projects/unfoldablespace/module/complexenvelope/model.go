package complexenvelope

import (
	"time"

	"github.com/heucuva/europi/units"
)

type model interface {
	Trigger()
	SetCV(cv units.CV)
	Tick(deltaTime time.Duration)
}
