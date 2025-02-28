package knobbank

import (
	"errors"

	"github.com/awonak/EuroPiGo/hardware/hal"
)

type KnobBank struct {
	hal.KnobInput
	current   int
	lastValue float32
	bank      []knobBankEntry
}

func NewKnobBank(knob hal.KnobInput, opts ...KnobBankOption) (*KnobBank, error) {
	if knob == nil {
		return nil, errors.New("knob is nil")
	}

	kb := &KnobBank{
		KnobInput: knob,
		lastValue: knob.ReadVoltage(),
	}

	for _, opt := range opts {
		if err := opt(kb); err != nil {
			return nil, err
		}
	}

	return kb, nil
}

func (kb *KnobBank) Configure(config hal.AnalogInputConfig) error {
	// Configure call on a KnobBank is not allowed
	return errors.New("unsupported")
	//return kb.knob.Configure(config)
}

func (kb *KnobBank) CurrentName() string {
	if len(kb.bank) == 0 {
		return ""
	}
	return kb.bank[kb.current].name
}

func (kb *KnobBank) CurrentIndex() int {
	return kb.current
}

func (kb *KnobBank) Current() hal.KnobInput {
	return kb
}

func (kb *KnobBank) ReadVoltage() float32 {
	value := kb.KnobInput.ReadVoltage()
	if len(kb.bank) == 0 {
		return value
	}

	cur := &kb.bank[kb.current]
	percent := kb.Percent()
	kb.lastValue = cur.update(percent, value, kb.lastValue)
	return cur.Value()
}

func (kb *KnobBank) Percent() float32 {
	percent := kb.KnobInput.Percent()
	if len(kb.bank) == 0 {
		return percent
	}

	cur := &kb.bank[kb.current]
	value := kb.KnobInput.ReadVoltage()
	kb.lastValue = cur.update(percent, value, kb.lastValue)
	return cur.Percent()
}

func (kb *KnobBank) Next() {
	if len(kb.bank) == 0 {
		kb.current = 0
		return
	}

	cur := &kb.bank[kb.current]
	cur.lock(kb.KnobInput, kb.lastValue)

	kb.current++
	if kb.current >= len(kb.bank) {
		kb.current = 0
	}
	kb.bank[kb.current].unlock()
}
