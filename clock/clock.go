package clock

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func NewClock() Clock {
	clockTmp := &realClock{}
	return clockTmp
}

func (c realClock) Now() time.Time {
	return time.Now().UTC()
}
