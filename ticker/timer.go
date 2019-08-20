package ticker

import (
	log "github.com/Sirupsen/logrus"
	"sync/atomic"
	"time"
)

type WbTimer struct {
	timer *Timer
	p     *TimerParams
	ts    int64
}

type timerCB func(args interface{}) //ticker call back

type TimerParams struct {
	Seconds int
	Args    interface{}
	Cb      timerCB
}

var tw *TimingWheel

func init() {
	tw = NewTimingWheel(10*time.Millisecond, 1000)
	tw.Start()
}

func StartTimer(params TimerParams) *WbTimer {
	tm := new(WbTimer)
	tm.timer = tw.AfterFunc(time.Second*time.Duration(params.Seconds), func() {
		params.Cb(params.Args)
	})
	tm.p = &params
	log.Debugf("start ticker, args:%+v, second:%d", params.Args, params.Seconds)
	return tm
}

func (r *WbTimer) Stop() {
	atomic.StoreInt64(&r.ts, 0)
}

func (r *WbTimer) Reset() {
	atomic.StoreInt64(&r.ts, time.Now().UnixNano())
}

func (r *WbTimer) ReStart() {
	r.timer = tw.AfterFunc(time.Second*time.Duration(r.p.Seconds), func() {
		r.p.Cb(r.p.Args)
	})
}
func (r *WbTimer) GetTS() int64 {
	return atomic.LoadInt64(&r.ts)
}

// 检查 ticker 是否还生效
func (r *WbTimer) CheckTimer(limit int64) bool {
	ts := r.GetTS()
	if time.Now().UnixNano()-ts < limit {
		log.Debugf("Timer.Invalid ts:%v,limit:%v", ts, limit)
		return true
	}
	return false
}
