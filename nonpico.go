//go:build !pico
// +build !pico

package europi

import (
	_ "github.com/awonak/EuroPiGo/internal/nonpico"
)

// This file exists to import the non-pico code into the active build
// do not remove this file or remove the init() function below

func init() {
}
