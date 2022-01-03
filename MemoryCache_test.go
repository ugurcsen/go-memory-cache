package MemoryCache

import (
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	mcm := NewMemoryCacheMap()
	a := "str1"
	mcm.Set("c1", time.Second*10, func() interface{} {
		return a
	})
	str, ok := mcm.Get("c1").(string)
	if !ok {
		t.Fatal("Cache could not cast to string")
	}
	t.Log(str)
	a ="str2"
	mcm.RefreshAll()
	str = mcm.Get("c1").(string)
	if !ok {
		t.Fatal("Cache could not cast to string")
	}
	t.Log(str)
	mcm.Flush()
}