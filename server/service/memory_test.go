package service

import (
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
