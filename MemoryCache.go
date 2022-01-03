package MemoryCache

import (
	"sync"
	"time"
)

type MemoryCache struct {
	sync.RWMutex
	data       interface{}
	createDate time.Time
	ttl        time.Duration
	dataCreator func() interface{}
}

func NewMemoryCache(ttl time.Duration, dataCreator func() interface{}) *MemoryCache {
	mc := new(MemoryCache)
	mc.dataCreator = dataCreator
	mc.ttl = ttl
	mc.Refresh()
	return mc
}

func (mc *MemoryCache) Refresh()  {
	mc.Lock()
	defer mc.Unlock()
	mc.createDate = time.Now()
	mc.data = mc.dataCreator()
}

func (mc *MemoryCache) Get() interface{} {
	if time.Now().Sub(mc.createDate) > mc.ttl {
		mc.Refresh()
	}
	return mc.data
}
