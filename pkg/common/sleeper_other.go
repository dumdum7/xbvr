//go:build !windows

package common

type Sleeper struct {
}

func NewSleeper() *Sleeper {
	return &Sleeper{}
}

func (s *Sleeper) Close() {
}
