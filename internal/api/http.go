package api

import "net/http"

func HttpEnable() {
	http.HandleFunc("/next", Next)
	http.HandleFunc("/batchNext", BatchNext)
	//http.HandleFunc("/segment",Req)
	http.ListenAndServe(":8080", nil)
}
