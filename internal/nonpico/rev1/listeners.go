//go:build !pico
// +build !pico

package rev1

import (
	"fmt"
	"math"
	"sync"

	"github.com/awonak/EuroPiGo/event"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/hardware/rev1"
	"github.com/awonak/EuroPiGo/lerp"
)

var (
	bus    = event.NewBus()
	voLerp = lerp.NewLerp32[uint16](0, math.MaxUint16)
)

func setupVoltageOutputListeners(cb func(id hal.HardwareId, voltage float32)) {
	for id := hal.HardwareIdVoltage1Output; id <= hal.HardwareIdVoltage6Output; id++ {
		fn := func(hid hal.HardwareId) func(HwMessagePwmValue) {
			return func(msg HwMessagePwmValue) {
				v := voLerp.ClampedInverseLerp(msg.Value) * rev1.MaxOutputVoltage
				cb(hid, v)
			}
		}(id)
		event.Subscribe(bus, fmt.Sprintf("hw_pwm_%d", id), fn)
	}
}

func setupDisplayOutputListener(cb func(id hal.HardwareId, op HwDisplayOp, params []int16)) {
	bus := bus
	id := hal.HardwareIdDisplay1Output
	event.Subscribe(bus, fmt.Sprintf("hw_display_%d", id), func(msg HwMessageDisplay) {
		cb(id, msg.Op, msg.Operands)
	})

}

var (
	states sync.Map
)

func setDigitalInput(id hal.HardwareId, value bool) {
	prevState, _ := states.Load(id)

	states.Store(id, value)
	bus.Post(fmt.Sprintf("hw_value_%d", id), HwMessageDigitalValue{
		Value: value,
	})

	if prevState != value {
		if value {
			// rising
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), HwMessageInterrupt{
				Change: hal.ChangeRising,
			})
		} else {
			// falling
			bus.Post(fmt.Sprintf("hw_interrupt_%d", id), HwMessageInterrupt{
				Change: hal.ChangeFalling,
			})
		}
	}
}

var (
	aiLerp = lerp.NewLerp32[uint16](rev1.DefaultCalibratedMinAI, rev1.DefaultCalibratedMaxAI)
)

func setAnalogInput(id hal.HardwareId, voltage float32) {
	bus.Post(fmt.Sprintf("hw_value_%d", id), HwMessageADCValue{
		Value: aiLerp.Lerp(voltage),
	})
}