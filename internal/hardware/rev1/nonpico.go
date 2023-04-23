//go:build !pico && test
// +build !pico,test

package rev1

import (
	"fmt"
	"image/color"
	"math"

	"github.com/heucuva/europi/internal/event"
	"github.com/heucuva/europi/internal/hardware/hal"
)

var (
	DefaultEventBus = event.NewBus()
)

//============= ADC =============//

type nonPicoAdc struct {
	bus   event.Bus
	id    hal.HardwareId
	value uint16
}

func newNonPicoAdc(bus event.Bus, id hal.HardwareId) adcProvider {
	adc := &nonPicoAdc{
		bus: bus,
		id:  id,
	}
	event.Subscribe(bus, fmt.Sprintf("hw_value_%d", id), func(msg HwMessageADCValue) {
		adc.value = msg.Value
	})
	return adc
}

func (a *nonPicoAdc) Get(samples int) uint16 {
	var sum int
	for i := 0; i < samples; i++ {
		sum += int(a.value)
	}
	return uint16(sum / samples)
}

//============= DigitalReader =============//

type nonPicoDigitalReader struct {
	bus     event.Bus
	id      hal.HardwareId
	value   bool
	handler func()
}

func newNonPicoDigitalReader(bus event.Bus, id hal.HardwareId) digitalReaderProvider {
	dr := &nonPicoDigitalReader{
		bus: bus,
		id:  id,
	}
	event.Subscribe(bus, fmt.Sprintf("hw_value_%d", id), func(msg HwMessageDigitalValue) {
		dr.value = msg.Value
	})
	return dr
}

func (d *nonPicoDigitalReader) Get() bool {
	// Invert signal to match expected behavior.
	return !d.value
}

func (d *nonPicoDigitalReader) SetHandler(changes hal.ChangeFlags, handler func()) {
	event.Subscribe(d.bus, fmt.Sprintf("hw_interrupt_%d", d.id), func(msg HwMessageInterrupt) {
		if (msg.Change & changes) != 0 {
			handler()
		}
	})
}

//============= PWM =============//

type nonPicoPwm struct {
	bus event.Bus
	id  hal.HardwareId
	v   float32
}

func newNonPicoPwm(bus event.Bus, id hal.HardwareId) pwmProvider {
	p := &nonPicoPwm{
		bus: bus,
		id:  id,
	}
	return p
}

func (p *nonPicoPwm) Configure(config hal.VoltageOutputConfig) error {
	return nil
}

func (p *nonPicoPwm) Set(v float32, ofs uint16) {
	invertedV := v * math.MaxUint16
	// volts := (float32(o.pwm.Top()) - invertedCv) - o.ofs
	volts := invertedV - float32(ofs)
	p.v = v
	p.bus.Post(fmt.Sprintf("hw_pwm_%d", p.id), HwMessagePwmValue{
		Value: uint16(volts),
	})
}

func (p *nonPicoPwm) Get() float32 {
	return p.v
}

//============= Display =============//

const (
	oledWidth  = 128
	oledHeight = 32
)

type nonPicoDisplayOutput struct {
	bus    event.Bus
	id     hal.HardwareId
	width  int16
	height int16
}

func newNonPicoDisplayOutput(bus event.Bus, id hal.HardwareId) displayProvider {
	dp := &nonPicoDisplayOutput{
		bus:    bus,
		id:     id,
		width:  oledWidth,
		height: oledHeight,
	}

	return dp
}

func (d *nonPicoDisplayOutput) ClearBuffer() {
	d.bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op: HwDisplayOpClearBuffer,
	})
}

func (d *nonPicoDisplayOutput) Size() (x, y int16) {
	return d.width, d.height
}
func (d *nonPicoDisplayOutput) SetPixel(x, y int16, c color.RGBA) {
	d.bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op:       HwDisplayOpSetPixel,
		Operands: []int16{x, y, int16(c.R), int16(c.B), int16(c.G), int16(c.A)},
	})
}

func (d *nonPicoDisplayOutput) Display() error {
	d.bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op: HwDisplayOpDisplay,
	})
	return nil
}

//============= Init =============//

func init() {
	hwDigital1Input = newDigitalInput(newNonPicoDigitalReader(DefaultEventBus, hal.HardwareIdDigital1Input))
	hwAnalog1Input = newAnalogInput(newNonPicoAdc(DefaultEventBus, hal.HardwareIdAnalog1Input))
	hwDisplay1Output = newDisplayOutput(newNonPicoDisplayOutput(DefaultEventBus, hal.HardwareIdDisplay1Output))
	hwButton1Input = newDigitalInput(newNonPicoDigitalReader(DefaultEventBus, hal.HardwareIdButton1Input))
	hwButton2Input = newDigitalInput(newNonPicoDigitalReader(DefaultEventBus, hal.HardwareIdButton2Input))
	hwKnob1Input = newAnalogInput(newNonPicoAdc(DefaultEventBus, hal.HardwareIdKnob1Input))
	hwKnob2Input = newAnalogInput(newNonPicoAdc(DefaultEventBus, hal.HardwareIdKnob2Input))
	hwCV1Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage1Output))
	hwCV2Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage2Output))
	hwCV3Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage3Output))
	hwCV4Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage4Output))
	hwCV5Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage5Output))
	hwCV6Output = newVoltageOuput(newNonPicoPwm(DefaultEventBus, hal.HardwareIdVoltage6Output))
	hwRandom1Generator = newRandomGenerator(nil)
}