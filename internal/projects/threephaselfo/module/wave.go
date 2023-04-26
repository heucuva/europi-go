package module

import (
	"time"

	"github.com/awonak/EuroPiGo/units"
)

type wave interface {
	Get(t, interval time.Duration) (units.BipolarCV, units.BipolarCV, units.BipolarCV)
	Mode() WaveMode
}
