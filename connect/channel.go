/*
   @Time : 2019-07-25 14:38
   @Author : frozenChen
   @File : channel
   @Software: KingMaxWMS_APP_API
*/
package connect

import (
	"bufio"

	"DailysServer/proto"
)

type Channel struct {
	Id     string
	Ring   *Ring
	msg    chan *proto.Proto
	Writer bufio.Writer
	Reader bufio.Reader
}

func NewChannel() *Channel {
	return &Channel{
		Ring: NewRing(),
		msg:  make(chan *proto.Proto,1024),
	}
}

func (c *Channel) Push(msg *proto.Proto) {
	select {
	case c.msg <- msg:
	default:

	}
	return
}

func (c *Channel) Ready() *proto.Proto {
	return <-c.msg
}

func (c *Channel) Signal() {
	c.msg <- proto.ProtoReady
}

func (c *Channel) Close() {
	c.msg <- proto.ProtoFinish
}
