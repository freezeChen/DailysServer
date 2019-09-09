/*
   @Time : 2019-07-25 14:38
   @Author : frozenChen
   @File : channel
   @Software: KingMaxWMS_APP_API
*/
package server

import "bufio"

type Channel struct {
	Id     string
	Ring   *Ring
	msg    chan *Proto
	Writer bufio.Writer
	Reader bufio.Reader
}

func NewChannel() *Channel {
	return &Channel{
		Ring: NewRing(),
		msg:  make(chan *Proto,1024),
	}
}

func (c *Channel) Push(msg *Proto) {
	select {
	case c.msg <- msg:
	default:

	}
	return
}

func (c *Channel) Ready() *Proto {
	return <-c.msg
}

func (c *Channel) Signal() {
	c.msg <- ProtoReady
}

func (c *Channel) Close() {
	c.msg <- ProtoFinish
}
