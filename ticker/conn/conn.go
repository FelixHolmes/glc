package conn

import (
	"github.com/FelixHolmes/glc/ticker"
	"time"
)

var wheel *ticker.TimingWheel

type Ticker struct {
}

func NewTimingWheel(tick time.Duration, wheelSize int64) {
	if wheel == nil {
		wheel = ticker.NewTimingWheel(tick, wheelSize)
	}
}

func (me *Ticker) Start() {

}
func (me *Ticker) Stop() {

}
func (me *Ticker) ReSet() {

}
