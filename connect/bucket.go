/*
   @Time : 2019-06-12 14:41
   @Author : frozenchen
   @File : bucket
   @Software: DailyServer
*/
package connect

import "sync"

type Bucket struct {
	lock sync.RWMutex
	chs  map[string]*Channel
}

func NewBucket() (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[string]*Channel)
	return b
}

func (b *Bucket) Online(key string, c *Channel) {
	b.lock.Lock()
	//该id已登录 下线
	if ch, ok := b.chs[key]; ok {
		ch.Close()
	}
	b.chs[key] = c
	//重复判断.
	b.lock.Unlock()
	return
}
func (b *Bucket) Get(key string) (c *Channel) {
	b.lock.RLock()
	c = b.chs[key]
	b.lock.RUnlock()
	return
}

//下线&离线
func (b *Bucket) Offline(Key string) {
	var (
		ok bool
	)
	b.lock.Lock()
	if _, ok = b.chs[Key]; ok {
		delete(b.chs, Key)
	}
	b.lock.Unlock()
}
