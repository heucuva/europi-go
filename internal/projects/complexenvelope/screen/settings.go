package screen

import (
	"fmt"
	"machine"
	"time"

	"github.com/heucuva/europi"
	"github.com/heucuva/europi/experimental/knobmenu"
	"github.com/heucuva/europi/internal/projects/complexenvelope/module"
	"github.com/heucuva/europi/output"
	"github.com/heucuva/europi/units"
)

type Settings struct {
	km  *knobmenu.KnobMenu
	Env *module.ComplexEnvelope
	cur int
}

func (m *Settings) modeString() string {
	return module.ModeString(m.Env.Mode(m.cur))
}

func (m *Settings) modeValue() units.CV {
	return module.ModeToCV(m.Env.Mode(m.cur))
}

func (m *Settings) setModeValue(value units.CV) {
	m.Env.SetMode(m.cur, module.CVToMode(value))
}

func (m *Settings) attackModeString() string {
	return module.AttackModeString(m.Env.AttackMode(m.cur))
}

func (m *Settings) attackModeValue() units.CV {
	return module.AttackModeToCV(m.Env.AttackMode(m.cur))
}

func (m *Settings) setAttackModeValue(value units.CV) {
	m.Env.SetAttackMode(m.cur, module.CVToAttackMode(value))
}

func (m *Settings) releaseModeString() string {
	return module.ReleaseModeString(m.Env.ReleaseMode(m.cur))
}

func (m *Settings) releaseModeValue() units.CV {
	return module.ReleaseModeToCV(m.Env.ReleaseMode(m.cur))
}

func (m *Settings) setReleaseModeValue(value units.CV) {
	m.Env.SetReleaseMode(m.cur, module.CVToReleaseMode(value))
}

func (m *Settings) attackString() string {
	return module.AttackString(m.Env.Attack(m.cur))
}

func (m *Settings) attackValue() units.CV {
	return module.AttackToCV(m.Env.Attack(m.cur))
}

func (m *Settings) setAttackValue(value units.CV) {
	m.Env.SetAttack(m.cur, module.CVToAttack(value))
}

func (m *Settings) decayString() string {
	return module.DecayString(m.Env.Decay(m.cur))
}

func (m *Settings) decayValue() units.CV {
	return module.DecayToCV(m.Env.Decay(m.cur))
}

func (m *Settings) setDecayValue(value units.CV) {
	m.Env.SetDecay(m.cur, module.CVToDecay(value))
}

func (m *Settings) Start(e *europi.EuroPi) {
	m.setupMenu(e)
}

func (m *Settings) setupMenu(e *europi.EuroPi) {
	km, err := knobmenu.NewKnobMenu(e.K1,
		knobmenu.WithItem("mode", "Mode", m.modeString, m.modeValue, m.setModeValue),
		knobmenu.WithItem("attackMode", "AttackMode", m.attackModeString, m.attackModeValue, m.setAttackModeValue),
		knobmenu.WithItem("releaseMode", "ReleaseMode", m.releaseModeString, m.releaseModeValue, m.setReleaseModeValue),
		knobmenu.WithItem("attack", "Attack", m.attackString, m.attackValue, m.setAttackValue),
		knobmenu.WithItem("decay", "Decay", m.decayString, m.decayValue, m.setDecayValue),
	)
	if err != nil {
		panic(err)
	}

	m.km = km
}

func (m *Settings) Button1Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Settings) Button1(e *europi.EuroPi, p machine.Pin) {
	m.km.Next()
}

func (m *Settings) Button2Debounce() time.Duration {
	return time.Millisecond * 200
}

func (m *Settings) Button2(e *europi.EuroPi, p machine.Pin) {
	m.cur = (m.cur + 1) % 2
	m.setupMenu(e)
}

func (m *Settings) Paint(e *europi.EuroPi, deltaTime time.Duration) {
	m.km.Paint(e, deltaTime)
	disp := e.Display
	disp.WriteLineInverseAligned(fmt.Sprint(m.cur+1), 0, line1y, output.AlignRight, output.AlignTop)
}
