package screen



import (
	"fmt"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/internal/projects/complexrandom/module"
)

type Main struct {
	ComplexRandom *module.ComplexRandom
}

const (
	line1y int16 = 11
	line2y int16 = 23
)

func (m *Main) Start(e *europi.EuroPi) {
}

func (m *Main) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	disp := e.Display
	disp.WriteLine(fmt.Sprintf("1:%2.1f 2:%2.1f 3:%2.1f", e.CV1.Voltage(), e.CV2.Voltage(), e.CV3.Voltage()), 0, line1y)
	disp.WriteLine(fmt.Sprintf("4:%2.1f 5:%2.1f 6:%2.1f", e.CV4.Voltage(), e.CV5.Voltage(), e.CV6.Voltage()), 0, line2y)
}
