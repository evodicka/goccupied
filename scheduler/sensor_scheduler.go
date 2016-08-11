package scheduler

import (
	"time"
)

var (
	ticker = time.NewTicker(5 * time.Second)
	quit = make(chan struct{})
)

func Start(call func()) {
	go func() {
		for {
			select {
			case <- ticker.C:
				call()
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func Stop() {
	close(quit)
}


