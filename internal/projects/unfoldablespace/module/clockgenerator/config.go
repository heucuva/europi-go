package clockgenerator

type Config struct {
	BPM      float32
	Enabled  bool
	ClockOut func()
}
