package service

import (
	"sync"
	"testing"
)

func TestMemoryIdGenerator_BatchNext(t *testing.T) {
	g := NewMemoryGenerator()

	ids, _ := g.BatchNext(100)
	if len(ids) < 100 {
		t.Error("获取id异常")
	}
	if ids[len(ids)-1] != 100 {
		t.Error("id序号异常")
	}
}

//BenchmarkMemoryIdGenerator_BatchNext-4   	 2680452	       420.5 ns/op
func BenchmarkMemoryIdGenerator_BatchNext(b *testing.B) {
	g := NewMemoryGenerator()
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.BatchNext(100)
		}()
	}
	wg.Wait()
}
