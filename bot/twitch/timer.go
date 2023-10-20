package twitch

import "time"

type intervalTimer struct {
	interval time.Duration
	handler  func()
	quit     chan bool
}

func newIntervalTimer(interval time.Duration, handler func()) *intervalTimer {
	return &intervalTimer{
		interval: interval,
		handler:  handler,
		quit:     make(chan bool, 1),
	}
}

func (t *intervalTimer) start() {
	go func() {
		for {
			select {
			case <-t.quit:
				return
			default:
				t.handler()
				time.Sleep(t.interval * time.Second)
			}
		}
	}()
}

func (t *intervalTimer) stop() {
	t.quit <- true
}
