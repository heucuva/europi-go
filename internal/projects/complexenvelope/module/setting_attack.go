package module

import (
	"fmt"

	"github.com/heucuva/europi/units"
)

func AttackString(cv units.CV) string {
	return fmt.Sprintf("+%d", int(cv*255))
}

func AttackToCV(cv units.CV) units.CV {
	return cv
}

func CVToAttack(cv units.CV) units.CV {
	return cv
}
