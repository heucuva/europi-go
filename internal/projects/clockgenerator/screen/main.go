package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/clockgenerator/module"
	"github.com/heucuva/europi/output"
)

type Main struct {
	Clock *module.ClockGenerator
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

func (m *Main) Start(e *europi.EuroPi) {
}

func (m *Main) Button1(e *europi.EuroPi, p machine.Pin) {
	m.Clock.Toggle()
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	disp := e.Display
	disp.ClearBuffer()
	if m.Clock.Enabled() {
		disp.DrawHLine(0, 0, 7, output.White)
	}
	disp.WriteLine(fmt.Sprintf("1:%2.1f", e.CV1.Voltage()), 0, line1y)
	disp.Display()
}
