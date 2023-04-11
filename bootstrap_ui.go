package europi

import (
	"context"
	"machine"
	"time"

	"github.com/heucuva/europi/input"
)

type UserInterface interface {
	Start(e *EuroPi)
	Paint(e *EuroPi, deltaTime time.Duration)
}

type UserInterfaceButton1 interface {
	Button1(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton1Long interface {
	Button1Long(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton2 interface {
	Button2(e *EuroPi, p machine.Pin)
}

type UserInterfaceButton2Long interface {
	Button2Long(e *EuroPi, p machine.Pin)
}

type uiModule struct {
	screen  UserInterface
	repaint chan struct{}
	stop    context.CancelFunc
}

var (
	ui uiModule
)

func enableUI(e *EuroPi, screen UserInterface, interval time.Duration) {
	ui.screen = screen
	if ui.screen == nil {
		return
	}

	ui.repaint = make(chan struct{}, 1)

	var (
		inputB1  func(e *EuroPi, p machine.Pin)
		inputB1L func(e *EuroPi, p machine.Pin)
	)
	if in, ok := screen.(UserInterfaceButton1); ok {
		inputB1 = in.Button1
	}
	if in, ok := screen.(UserInterfaceButton1Long); ok {
		inputB1L = in.Button1Long
	}
	ui.setupButton(e, e.B1, inputB1, inputB1L)

	var (
		inputB2  func(e *EuroPi, p machine.Pin)
		inputB2L func(e *EuroPi, p machine.Pin)
	)
	if in, ok := screen.(UserInterfaceButton2); ok {
		inputB2 = in.Button2
	}
	if in, ok := screen.(UserInterfaceButton2Long); ok {
		inputB2L = in.Button2Long
	}
	ui.setupButton(e, e.B2, inputB2, inputB2L)

	go ui.run(e, interval)
}

func startUI(e *EuroPi) {
	if ui.screen == nil {
		return
	}

	ui.screen.Start(e)
}

// ForceRepaintUI schedules a forced repaint of the UI (if it is configured and running)
func ForceRepaintUI(e *EuroPi) {
	if ui.repaint != nil {
		ui.repaint <- struct{}{}
	}
}

func disableUI(e *EuroPi) {
	if ui.stop != nil {
		ui.stop()
	}

	if ui.repaint != nil {
		close(ui.repaint)
	}
}

func (u *uiModule) run(e *EuroPi, interval time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())
	ui.stop = cancel
	defer ui.stop()

	t := time.NewTicker(interval)
	defer t.Stop()

	lastTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			return

		case <-ui.repaint:
			now := time.Now()
			deltaTime := now.Sub(lastTime)
			lastTime = now
			u.screen.Paint(e, deltaTime)

		case now := <-t.C:
			deltaTime := now.Sub(lastTime)
			lastTime = now
			u.screen.Paint(e, deltaTime)
		}
	}
}

func (u *uiModule) setupButton(e *EuroPi, r input.DigitalReader, onShort, onLong func(e *EuroPi, p machine.Pin)) {
	if onShort == nil && onLong == nil {
		return
	}

	if onShort == nil {
		// no-op
		onShort = func(e *EuroPi, p machine.Pin) {}
	}

	// if no long-press handler present, just reuse short-press handler
	if onLong == nil {
		onLong = onShort
	}

	const longDuration = time.Second

	r.HandlerEx(machine.PinFalling, func(p machine.Pin) {
		startDown := r.LastInput()
		deltaTime := time.Now().Sub(startDown)
		if deltaTime < longDuration {
			onShort(e, p)
		} else {
			onLong(e, p)
		}
	})
}
