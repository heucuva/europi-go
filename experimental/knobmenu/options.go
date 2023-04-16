package knobmenu

import (
	"fmt"

	"github.com/heucuva/europi/experimental/knobbank"
	"github.com/heucuva/europi/units"
)

type KnobMenuOption func(km *KnobMenu) ([]knobbank.KnobBankOption, error)

func WithItem(name, label string, stringFn func() string, valueFn func() units.CV, updateFn func(value units.CV)) KnobMenuOption {
	return func(km *KnobMenu) ([]knobbank.KnobBankOption, error) {
		for _, it := range km.items {
			if it.name == name {
				return nil, fmt.Errorf("item %q already exists", name)
			}
		}

		km.items = append(km.items, item{
			name:     name,
			label:    label,
			stringFn: stringFn,
			updateFn: updateFn,
		})

		return []knobbank.KnobBankOption{
			knobbank.WithLockedKnob(name, knobbank.InitialPercentageValue(valueFn().ToFloat32())),
		}, nil
	}
}

func WithPosition(x, y int16) KnobMenuOption {
	return func(km *KnobMenu) ([]knobbank.KnobBankOption, error) {
		km.x = x
		km.y = y
		return nil, nil
	}
}

func WithYAdvance(yadvance int16) KnobMenuOption {
	return func(km *KnobMenu) ([]knobbank.KnobBankOption, error) {
		km.yadvance = yadvance
		return nil, nil
	}
}