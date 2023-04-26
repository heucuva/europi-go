package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func DecayString(cv units.CV) string {
	return fmt.Sprintf("+%d", int(cv*255))
}

func DecayToCV(cv units.CV) units.CV {
	return cv
}

func CVToDecay(cv units.CV) units.CV {
	return cv
}
