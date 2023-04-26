package module

import (
	"fmt"

	"github.com/awonak/EuroPiGo/units"
)

func SkewShapeString(cv units.CV) string {
	return fmt.Sprintf("%3.1f%%", cv*100.0)
}

func SkewShapeToCV(cv units.CV) units.CV {
	return cv
}

func CVToSkewShape(cv units.CV) units.CV {
	return cv
}
