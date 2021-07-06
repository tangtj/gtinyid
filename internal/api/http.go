package api

import (
	"log"
	"net/http"
	"sync"
)

func HttpEnable(wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/next", Next)
	http.HandleFunc("/batchNext", BatchNext)
	http.HandleFunc("/segment", Segment)
	err := http.ListenAndServe(":8080", nil)
	log.Printf("http 服务端口监听异常 : %s", err)
}
