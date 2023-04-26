package screen

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"github.com/awonak/EuroPiGo/internal/projects/cascadelfo/module"
	"tinygo.org/x/tinyfont/proggy"
)

type Main struct {
	LFO *module.CascadeLFO
	w   fontwriter.Writer
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

func (m *Main) Start(e *europi.EuroPi) {
	m.w.Display = e.Display
	m.w.Font = &proggy.TinySZ8pt7b
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.w.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y, draw.White)
	m.w.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y, draw.White)
}
