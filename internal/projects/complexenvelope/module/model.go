package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type model interface {
	Trigger()
	SetCV(cv units.BipolarCV)
	Tick(deltaTime time.Duration)
}
