package timer

import "time"

type TimeWheel interface {
	Init(conf TimeConfig) error
	AddJob(ts time.Duration, f func()) *Timer
}

type Timer interface {
	Start()
	Stop()
	ReSet()
}
type TimeConfig struct {
}
