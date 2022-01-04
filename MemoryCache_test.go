package main

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

	v, err := mcm.Get("c1")
	if err != nil {
		t.Fatal("c1 not set")
	}
	str, ok := v.(string)
	if !ok {
		t.Fatal("Cache could not cast to string")
	}
	t.Log(str)
	a ="str2"
	mcm.RefreshAll()
	v2, err := mcm.Get("c1")
	if err != nil {
		t.Fatal("c1 not set")
	}
	var str2 string
	str2, ok = v2.(string)
	if !ok {
		t.Fatal("Cache could not cast to string")
	}
	t.Log(str2)

	if str == str2 {
		t.Fatal("Cache did not refresh")
	}
	mcm.Flush()
	t.Log("flushed")
}