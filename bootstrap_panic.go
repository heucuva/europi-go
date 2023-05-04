package europi

import (
	"fmt"
	"log"
	"os"

	"github.com/awonak/EuroPiGo/experimental/draw"
	"tinygo.org/x/tinydraw"
)

// DefaultPanicHandler is the default handler for panics
// This will be set by the build flag `onscreenpanic` to `handlePanicOnScreenLog`
// Not setting the build flag will set it to `handlePanicDisplayCrash`
var DefaultPanicHandler func(e Hardware, reason any)

var (
	// silence linter
	_ = handlePanicOnScreenLog
)

func handlePanicOnScreenLog(e Hardware, reason any) {
	if e == nil {
		// can't do anything if it's not enabled
		return
	}

	// force display-logging to enabled
	enableDisplayLogger(e)

	// show the panic on the screen
	log.Println(fmt.Sprint(reason))

	flushDisplayLogger(e)

	os.Exit(1)
}

func handlePanicLogger(e Hardware, reason any) {
	log.Panic(reason)
}

func handlePanicDisplayCrash(e Hardware, reason any) {
	display := Display(e)
	if display == nil {
		// can't do anything if we don't have a display
		return
	}

	// display a diagonal line pattern through the screen to show that the EuroPi is crashed
	width, height := display.Size()
	ymax := height - 1
	for x := -ymax; x < width; x += 4 {
		lx, ly := x, int16(0)
		if x < 0 {
			lx = 0
			ly = -x
		}
		tinydraw.Line(display, lx, ly, x+ymax, ymax, draw.White)
	}
}
