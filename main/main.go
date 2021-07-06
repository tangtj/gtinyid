package main

import (
	"github.com/tangtj/gtinyid/internal/api"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		api.HttpEnable(wg)
	}()

	go func() {
		api.GrpcEnable(wg)
	}()

	wg.Wait()
}
