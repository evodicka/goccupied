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
	go runner(ticker, toCall)
}

func StartPeriod(toCall func(), period int) {
	ticker = time.NewTicker(time.Duration(period) * time.Second)
	go runner(ticker, toCall)
}

func runner(ticker *time.Ticker, toCall func()) {
	for {
		select {
		case <-ticker.C:
			toCall()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

// Stops the scheduler
func Stop() {
	close(quit)
}
