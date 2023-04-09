package module

type gate struct {
	out       func(high bool)
	active    bool
	lastInput bool
	chance    float32
}
