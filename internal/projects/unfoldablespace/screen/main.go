package screen

import (
	"fmt"
	"time"

	europi "github.com/awonak/EuroPiGo"
	"github.com/awonak/EuroPiGo/experimental/draw"
	"github.com/awonak/EuroPiGo/experimental/fontwriter"
	"github.com/awonak/EuroPiGo/internal/projects/unfoldablespace/module"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont/proggy"
)

type Main struct {
	Unfold *module.UnfoldableSpace
	w      fontwriter.Writer
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

func (m *Main) Start(e *europi.EuroPi) {
	m.w.Display = e.Display
	m.w.Font = &proggy.TinySZ8pt7b
}

func (m *Main) Button1(e *europi.EuroPi, _ time.Duration) {
	m.Unfold.ToggleInternalClock()
}

func (m *Main) Button2Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Main) Button2(e *europi.EuroPi, _ time.Duration) {
	m.Unfold.ToggleVOctInputEnabled()
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	disp := e.Display
	if m.Unfold.InternalClockEnabled() {
		tinydraw.Line(disp, 0, 0, 7, 0, draw.White)
	}
	if m.Unfold.VOctInputEnabled() {
		width, _ := disp.Size()
		tinydraw.Line(disp, width-7, 0, width, 0, draw.White)
	}
	m.w.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y, draw.White)
	m.w.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y, draw.White)
}
