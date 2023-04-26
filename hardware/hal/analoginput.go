package hal

import "github.com/awonak/EuroPiGo/units"

type AnalogInput interface {
	Configure(config AnalogInputConfig) error
	Percent() float32
	ReadVoltage() float32
	ReadCV() units.CV
	ReadBipolarCV() units.BipolarCV
	ReadVOct() units.VOct
	MinVoltage() float32
	MaxVoltage() float32
}

type AnalogInputConfig struct {
	Samples         int
	CalibratedMinAI uint16
	CalibratedMaxAI uint16
}
