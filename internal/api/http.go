package api

import "net/http"

func HttpEnable() {
	http.HandleFunc("/next", Next)
	http.HandleFunc("/batchNext", BatchNext)
	http.HandleFunc("/segment", Segment)
	http.ListenAndServe(":8080", nil)
}
