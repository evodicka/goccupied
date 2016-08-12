package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const state = "state"
const timestamp = "time"

var lightCache = cache.New(60*time.Minute, time.Minute)

// Stores the received value in a cache if it has changed and updates the time
func StoreState(on bool) {
	if x, found := lightCache.Get(state); found {
		current := x.(bool)
		if current != on {
			store(on)
		}
	} else {
		store(on)
	}
}

func store(on bool) {
	lightCache.Set(state, on, cache.NoExpiration)
	lightCache.Set(timestamp, time.Now(), cache.NoExpiration)
}

// Retrieves the current state and duration from the cache
func GetState() (bool, time.Duration) {
	currentState := false
	if x, found := lightCache.Get(state); found {
		currentState = x.(bool)
	}

	currentSince := time.Now()
	if y, found := lightCache.Get(timestamp); found {
		currentSince = y.(time.Time)
	}

	return currentState, calculateDuration(currentSince)
}

func calculateDuration(from time.Time) time.Duration {
	return time.Since(from)
}
