package cache

import(
	"github.com/patrickmn/go-cache"
	"time"
)

const state = "state"
const timestamp = "time"

var lightCache = cache.New(60*time.Minute, time.Minute)

func StoreState(on bool) {
	if x, found := lightCache.Get(state); found {
		current := x.(bool)
		if(current != on) {
			lightCache.Set(state, on, cache.NoExpiration)
			lightCache.Set(timestamp, time.Now(), cache.NoExpiration)
		}
	}
}

func GetState() (bool, time.Duration) {
	currentState := false
	if x, found := lightCache.Get(state); found {
		currentState = x.(bool)
	}

	currentSince := time.Now()
	if y, found := lightCache.Get(state); found {
		currentSince = y.(time.Time)
	}

	return currentState, calculateDuration(currentSince)
}

func calculateDuration(from time.Time) time.Duration {
	return time.Since(from);
}
