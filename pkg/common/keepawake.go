package common

import (
	"sync"
	"time"
)

var (
	keepAwakeMutex sync.Mutex
	lastActivity   time.Time
	globalSleeper  *Sleeper
	timerStarted   bool
)

// KeepAwake marks that user or client activity has occurred.
// This will prevent the system from sleeping for the next 10 minutes.
func KeepAwake() {
	keepAwakeMutex.Lock()
	defer keepAwakeMutex.Unlock()

	lastActivity = time.Now()

	if globalSleeper == nil {
		globalSleeper = NewSleeper()
	}

	if !timerStarted {
		timerStarted = true
		go keepAwakeLoop()
	}
}

func keepAwakeLoop() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		keepAwakeMutex.Lock()
		if time.Since(lastActivity) >= 10*time.Minute {
			if globalSleeper != nil {
				globalSleeper.Close()
				globalSleeper = nil
			}
			timerStarted = false
			keepAwakeMutex.Unlock()
			return
		}
		keepAwakeMutex.Unlock()
	}
}
