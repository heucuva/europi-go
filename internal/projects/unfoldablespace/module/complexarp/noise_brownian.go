package complexarp

import "math/rand"

type noiseBrownian struct {
	beta float32
	prev float32
}

func (n *noiseBrownian) Get() float32 {
	white := rand.Float32()
	n.prev = (n.prev + (n.beta * white)) / (1 + n.beta)
	return n.prev
}
