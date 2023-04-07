package europi

import (
	"log"
	"machine"
	"math/rand"
	"os"

	"github.com/heucuva/europi/experimental/displaylogger"
)

func enableDisplayLogger(e *EuroPi) {
	log.SetFlags(0)
	log.SetOutput(&displaylogger.Logger{
		Display: e.Display,
	})
}

func disableDisplayLogger(e *EuroPi) {
	log.SetOutput(os.Stdout)
}

func initRandom(e *EuroPi) {
	xl, _ := machine.GetRNG()
	xh, _ := machine.GetRNG()
	x := int64(xh)<<32 | int64(xl)
	rand.Seed(x)
}

func uninitRandom(e *EuroPi) {
}
