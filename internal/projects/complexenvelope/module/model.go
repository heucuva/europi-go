package module

import (
	"time"

	"github.com/heucuva/europi/units"
)

type model interface {
	Trigger()
	SetCV(cv units.BipolarCV)
	SetAttack(cv units.CV)
	SetDecay(cv units.CV)
	Tick(deltaTime time.Duration)
}
