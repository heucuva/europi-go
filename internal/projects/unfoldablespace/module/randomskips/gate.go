package randomskips

type gate struct {
	out     func(high bool)
	enabled bool
	chance  float32
}
