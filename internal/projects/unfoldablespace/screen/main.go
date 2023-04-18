package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/unfoldablespace/module"
	"github.com/heucuva/europi/output"
)

type Main struct {
	Unfold *module.UnfoldableSpace
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

func (m *Main) Start(e *europi.EuroPi) {
}

func (m *Main) Button1(e *europi.EuroPi, p machine.Pin) {
	m.Unfold.ToggleInternalClock()
}

func (m *Main) Button2Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Main) Button2(e *europi.EuroPi, p machine.Pin) {
	m.Unfold.ToggleVOctInputEnabled()
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	disp := e.Display
	if m.Unfold.InternalClockEnabled() {
		disp.DrawHLine(0, 0, 7, output.White)
	}
	if m.Unfold.VOctInputEnabled() {
		disp.DrawHLine(output.OLEDWidth-7, 0, 7, output.White)
	}
	disp.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y)
	disp.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y)
}
