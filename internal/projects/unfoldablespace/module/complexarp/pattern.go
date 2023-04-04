package complexarp

import "github.com/heucuva/europi/units"

type pattern interface {
	Init(config Config) error
	Next() units.VOct
}
