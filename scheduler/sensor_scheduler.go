package scheduler

import (
	"time"
)

var (
	ticker = time.NewTicker(5 * time.Second)
	quit   = make(chan struct{})
)

// Starts the scheduler. It periodically calls the passed function
func Start(toCall func()) {
	go func() {
		for {
			select {
			case <-ticker.C:
				toCall()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// Stops the scheduler
func Stop() {
	close(quit)
}
