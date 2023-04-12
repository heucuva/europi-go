package module

import "time"

type Mode int

const (
	// Mode1msTrig = fixed 1ms trigger output
	Mode1msTrig = Mode(iota)
	// Mode200msTrig = fixed 200ms trigger output
	Mode200msTrig
	// ModeQuarterGateTrig = variable trigger output equal to a quarter of the input gate length
	ModeQuarterGateTrig
	// ModeHalfGateTrig = variable trigger output equal to half of the input gate length
	ModeHalfGateTrig
	// ModeEqualGateTrig = variable trigger output equal to the input gate length
	ModeEqualGateTrig
)

func (m Mode) GetTriggerDuration(input time.Duration) time.Duration {
	switch m {
	case Mode1msTrig:
		return time.Millisecond * 1
	case Mode200msTrig:
		return time.Millisecond * 200
	case ModeQuarterGateTrig:
		return input / 4
	case ModeHalfGateTrig:
		return input / 2
	case ModeEqualGateTrig:
		return input
	default:
		panic("unsupported mode")
	}
}
