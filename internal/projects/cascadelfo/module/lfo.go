package module

import (
	"math"

	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

var (
	/* generated by:

	package main

	import (
		"fmt"
		"math"
	)

	func main() {
		p := -math.Pi
		d := 2.0 * math.Pi / 256.0
		a := 2.0 / math.Pi
		for i := 0; i < 256; i++ {
			v := a * math.Asin(math.Sin(p))
			fmt.Printf("%f, ", v)
			if (i+1)%16 == 0 {
				fmt.Println()
			}
			p += d
		}
	}
	*/
	lfoTriangle = [256]units.BipolarCV{
		0.000000, -0.015625, -0.031250, -0.046875, -0.062500, -0.078125, -0.093750, -0.109375, -0.125000, -0.140625, -0.156250, -0.171875, -0.187500, -0.203125, -0.218750, -0.234375,
		-0.250000, -0.265625, -0.281250, -0.296875, -0.312500, -0.328125, -0.343750, -0.359375, -0.375000, -0.390625, -0.406250, -0.421875, -0.437500, -0.453125, -0.468750, -0.484375,
		-0.500000, -0.515625, -0.531250, -0.546875, -0.562500, -0.578125, -0.593750, -0.609375, -0.625000, -0.640625, -0.656250, -0.671875, -0.687500, -0.703125, -0.718750, -0.734375,
		-0.750000, -0.765625, -0.781250, -0.796875, -0.812500, -0.828125, -0.843750, -0.859375, -0.875000, -0.890625, -0.906250, -0.921875, -0.937500, -0.953125, -0.968750, -0.984375,
		-1.000000, -0.984375, -0.968750, -0.953125, -0.937500, -0.921875, -0.906250, -0.890625, -0.875000, -0.859375, -0.843750, -0.828125, -0.812500, -0.796875, -0.781250, -0.765625,
		-0.750000, -0.734375, -0.718750, -0.703125, -0.687500, -0.671875, -0.656250, -0.640625, -0.625000, -0.609375, -0.593750, -0.578125, -0.562500, -0.546875, -0.531250, -0.515625,
		-0.500000, -0.484375, -0.468750, -0.453125, -0.437500, -0.421875, -0.406250, -0.390625, -0.375000, -0.359375, -0.343750, -0.328125, -0.312500, -0.296875, -0.281250, -0.265625,
		-0.250000, -0.234375, -0.218750, -0.203125, -0.187500, -0.171875, -0.156250, -0.140625, -0.125000, -0.109375, -0.093750, -0.078125, -0.062500, -0.046875, -0.031250, -0.015625,
		0.000000, 0.015625, 0.031250, 0.046875, 0.062500, 0.078125, 0.093750, 0.109375, 0.125000, 0.140625, 0.156250, 0.171875, 0.187500, 0.203125, 0.218750, 0.234375,
		0.250000, 0.265625, 0.281250, 0.296875, 0.312500, 0.328125, 0.343750, 0.359375, 0.375000, 0.390625, 0.406250, 0.421875, 0.437500, 0.453125, 0.468750, 0.484375,
		0.500000, 0.515625, 0.531250, 0.546875, 0.562500, 0.578125, 0.593750, 0.609375, 0.625000, 0.640625, 0.656250, 0.671875, 0.687500, 0.703125, 0.718750, 0.734375,
		0.750000, 0.765625, 0.781250, 0.796875, 0.812500, 0.828125, 0.843750, 0.859375, 0.875000, 0.890625, 0.906250, 0.921875, 0.937500, 0.953125, 0.968750, 0.984375,
		1.000000, 0.984375, 0.968750, 0.953125, 0.937500, 0.921875, 0.906250, 0.890625, 0.875000, 0.859375, 0.843750, 0.828125, 0.812500, 0.796875, 0.781250, 0.765625,
		0.750000, 0.734375, 0.718750, 0.703125, 0.687500, 0.671875, 0.656250, 0.640625, 0.625000, 0.609375, 0.593750, 0.578125, 0.562500, 0.546875, 0.531250, 0.515625,
		0.500000, 0.484375, 0.468750, 0.453125, 0.437500, 0.421875, 0.406250, 0.390625, 0.375000, 0.359375, 0.343750, 0.328125, 0.312500, 0.296875, 0.281250, 0.265625,
		0.250000, 0.234375, 0.218750, 0.203125, 0.187500, 0.171875, 0.156250, 0.140625, 0.125000, 0.109375, 0.093750, 0.078125, 0.062500, 0.046875, 0.031250, 0.015625,
	}
)

type lfo struct {
	pos float32
	out func(cv units.BipolarCV)
}

func (o *lfo) Update(delta float32) {
	l := len(lfoTriangle)
	newp := o.pos + delta*float32(l)
	p, f := math.Modf(float64(newp))
	for int(p) >= l {
		newp -= float32(l)
		p -= float64(l)
	}
	if o.pos != newp {
		o.pos = newp
		p0 := int(p)
		x0 := lfoTriangle[p0]
		if f != 0 {
			p1 := (p0 + 1) % len(lfoTriangle)
			x1 := lfoTriangle[p1]
			x0 = lerp.NewLerp32(x0, x1).Lerp(float32(f))
		}

		o.out(x0)
	}
}

func (o *lfo) Reset() {
	o.pos = 0
}