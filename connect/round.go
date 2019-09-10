/*
   @Time : 2019-07-31 11:13
   @Author : frozenChen
   @File : round
   @Software: KingMaxWMS_APP_API
*/
package connect

import (
	"DailysServer/pkg/timer"
)

const (
	_timer     = 32
	_timerSize = 2048
)

type Round struct {
	timers []timer.Timer
}

func NewRound() *Round {
	r := &Round{}
	r.timers = make([]timer.Timer, _timer)
	for i := 0; i < _timer; i++ {
		r.timers[i].Init(_timerSize)
	}
	return r
}

func (r *Round) Timer(rn int) *timer.Timer {
	return &(r.timers[rn%_timer])
}
