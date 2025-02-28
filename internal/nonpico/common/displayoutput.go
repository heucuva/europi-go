//go:build !pico
// +build !pico

package common

import (
	"fmt"
	"image/color"

	"github.com/awonak/EuroPiGo/hardware/common"
	"github.com/awonak/EuroPiGo/hardware/hal"
)

const (
	oledWidth  = 128
	oledHeight = 32
)

type nonPicoDisplayOutput struct {
	id     hal.HardwareId
	width  int16
	height int16
}

var (
	// static check
	_ common.DisplayProvider = (*nonPicoDisplayOutput)(nil)
)

func NewNonPicoDisplayOutput(id hal.HardwareId) *nonPicoDisplayOutput {
	dp := &nonPicoDisplayOutput{
		id:     id,
		width:  oledWidth,
		height: oledHeight,
	}

	return dp
}

func (d *nonPicoDisplayOutput) ClearBuffer() {
	bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op: HwDisplayOpClearBuffer,
	})
}

func (d *nonPicoDisplayOutput) Size() (x, y int16) {
	return d.width, d.height
}
func (d *nonPicoDisplayOutput) SetPixel(x, y int16, c color.RGBA) {
	bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op:       HwDisplayOpSetPixel,
		Operands: []int16{x, y, int16(c.R), int16(c.B), int16(c.G), int16(c.A)},
	})
}

func (d *nonPicoDisplayOutput) Display() error {
	bus.Post(fmt.Sprintf("hw_display_%d", d.id), HwMessageDisplay{
		Op: HwDisplayOpDisplay,
	})
	return nil
}
