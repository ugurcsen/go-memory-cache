package MemoryCache

import "time"

type MemoryCacheMap map[string]*MemoryCache

func NewMemoryCacheMap() MemoryCacheMap {
	return make(MemoryCacheMap)
}

func (mcm MemoryCacheMap) Get(key string) interface{} {
	if v ,ok:= mcm[key]; ok {
		return v.Get()
	}
	return nil
}

func (mcm MemoryCacheMap) Set(key string, ttl time.Duration, f func() interface{})  {
	mcm[key] = NewMemoryCache(ttl, f)
}

func (mcm MemoryCacheMap) Remove(key string)  {
	delete(mcm, key)
}

func (mcm *MemoryCacheMap) Flush()  {
	*mcm = NewMemoryCacheMap()
}

func (mcm MemoryCacheMap) Refresh(key string)  {
	mcm[key].Refresh()
}

func (mcm MemoryCacheMap) RefreshAll()  {
	for s, _ := range mcm {
		mcm[s].Refresh()
	}
}