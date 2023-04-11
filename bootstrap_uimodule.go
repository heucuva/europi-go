package europi

import (
	"context"
	"machine"
	"sync"
	"time"

	"github.com/heucuva/europi/input"
)

type uiModule struct {
	screen  UserInterface
	repaint chan struct{}
	stop    context.CancelFunc
	wg      sync.WaitGroup
}

func (u *uiModule) wait() {
	u.wg.Wait()
}

func (u *uiModule) run(e *EuroPi, interval time.Duration) {
	defer u.wg.Done()

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

	const longDuration = time.Millisecond * 650

	r.Handler(func(p machine.Pin) {
		startDown := r.LastChange()
		deltaTime := time.Now().Sub(startDown)
		if deltaTime < longDuration {
			onShort(e, p)
		} else {
			onLong(e, p)
		}
	})
}
