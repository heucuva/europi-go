package module

type Config struct {
	Gate   [1]func(high bool)
	Chance float32
}
