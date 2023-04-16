package module

type Config struct {
	Trigger [3]func(high bool)
	Gate    [3]func(high bool)
	Mode    Mode
}
