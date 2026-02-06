//go:build windows

package common

import (
	"runtime"
	"sync"
)

type Sleeper struct {
	wg sync.WaitGroup
}

func NewSleeper() *Sleeper {
	s := &Sleeper{}
	s.wg.Add(1)
	go s.run()
	return s
}

func (s *Sleeper) run() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	PreventSleep()
	s.wg.Wait()
	AllowSleep()
}

func (s *Sleeper) Close() {
	s.wg.Done()
}
