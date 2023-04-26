package rev1

import (
	"errors"

	"github.com/awonak/EuroPiGo/clamp"
	"github.com/awonak/EuroPiGo/hardware/hal"
	"github.com/awonak/EuroPiGo/lerp"
	"github.com/awonak/EuroPiGo/units"
)

const (
	// DefaultCalibrated[Min|Max]AI was calculated using the EuroPi calibration program:
	// https://github.com/Allen-Synthesis/EuroPi/blob/main/software/programming_instructions.md#calibrate-the-module
	DefaultCalibratedMinAI = 300
	DefaultCalibratedMaxAI = 44009

	DefaultSamples = 1000

	MaxInputVoltage = 10.0
	MinInputVoltage = 0.0
)

// Analog is a struct for handling the reading of analogue control voltage.
// The analogue input allows you to 'read' CV from anywhere between 0 and 12V.
type analoginput struct {
	adc     ADCProvider
	samples int
	cal     lerp.Lerper32[uint16]
}

var (
	// static check
	_ hal.AnalogInput = &analoginput{}
	// silence linter
	_ = newAnalogInput
)

type ADCProvider interface {
	Get(samples int) uint16
}

// newAnalogInput creates a new Analog Input
func newAnalogInput(adc ADCProvider) *analoginput {
	return &analoginput{
		adc:     adc,
		samples: DefaultSamples,
		cal:     lerp.NewLerp32[uint16](DefaultCalibratedMinAI, DefaultCalibratedMaxAI),
	}
}

// Configure updates the device with various configuration parameters
func (a *analoginput) Configure(config hal.AnalogInputConfig) error {
	if config.Samples == 0 {
		return errors.New("samples must be non-zero")
	}

	if config.CalibratedMinAI != 0 || config.CalibratedMaxAI != 0 {
		if config.CalibratedMinAI == config.CalibratedMaxAI {
			return errors.New("calibratedminai and calibratedmaxai must be different")
		} else if config.CalibratedMinAI > config.CalibratedMaxAI {
			return errors.New("calibtatedminai must be less than calibratedmaxai")
		}
		a.cal = lerp.NewLerp32(config.CalibratedMinAI, config.CalibratedMaxAI)
	}

	a.samples = config.Samples
	return nil
}

// ReadVoltage returns the current percentage read between 0.0 and 1.0.
func (a *analoginput) Percent() float32 {
	return a.cal.InverseLerp(a.adc.Get(a.samples))
}

func (a *analoginput) BipolarPercent() float32 {
	// TODO: make this actually work on Bipolar data
	return a.Percent()*2.0 - 1.0
}

// ReadVoltage returns the current read voltage between 0.0 and 10.0 volts.
func (a *analoginput) ReadVoltage() float32 {
	// NOTE: if MinInputVoltage ever becomes non-zero, then we need to use a lerp instead
	return a.Percent() * MaxInputVoltage
}

func (a *analoginput) ReadBipolarVoltage() float32 {
	// NOTE: if MinInputVoltage ever becomes non-zero, then we need to use a lerp instead
	return a.BipolarPercent() * MaxInputVoltage
}

// ReadCV returns the current read voltage as a CV value.
func (a *analoginput) ReadCV() units.CV {
	// we can't use a.Percent() here, because we might get over 5.0 volts input
	// just clamp it
	v := a.ReadVoltage()
	return clamp.Clamp(units.CV(v/5.0), 0.0, 1.0)
}

// ReadBipolarCV returns the current read voltage as a BipolarCV value.
func (a *analoginput) ReadBipolarCV() units.BipolarCV {
	// we can't use a.BipolarPercent() here, because we might get over 5.0 volts or under -5.0 volts input
	// just clamp it
	v := a.ReadBipolarVoltage()
	return clamp.Clamp(units.BipolarCV(v/5.0), -1.0, 1.0)
}

// ReadCV returns the current read voltage as a V/Octave value.
func (a *analoginput) ReadVOct() units.VOct {
	return units.VOct(a.ReadVoltage())
}

// MinVoltage returns the minimum voltage that that input can ever read by this device
func (a *analoginput) MinVoltage() float32 {
	return MinInputVoltage
}

// MaxVoltage returns the maximum voltage that the input can ever read by this device
func (a *analoginput) MaxVoltage() float32 {
	return MaxInputVoltage
}
