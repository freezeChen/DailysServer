package server

import (
	"errors"
)

var (
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
)

type Ring struct {
	rp   uint64
	wp   uint64
	num  uint64
	data []Proto
}

func NewRing() (r *Ring) {
	r = new(Ring)
	r.num = 7
	r.data = make([]Proto, 8)
	return
}

func (r *Ring) Get() (msg *Proto, err error) {
	if r.rp == r.wp {
		return nil, ErrRingEmpty
	}

	msg = &(r.data[r.rp&r.num])
	return
}

func (r *Ring) GetAdv() {
	r.rp++
}

func (r *Ring) Set() (msg *Proto, err error) {
	if r.wp-r.rp >= r.num {
		return nil, ErrRingFull
	}
	msg = &(r.data[r.wp&r.num])
	return
}

func (r *Ring) SetAdv() {
	r.wp++
}
