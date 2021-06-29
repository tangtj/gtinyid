package api

import (
	"encoding/json"
	"github.com/tangtj/gtinyid/server/model"
	"github.com/tangtj/gtinyid/server/service"
	"log"
	"net/http"
	"strconv"
)

var generator service.IdGenService = service.NewMemoryGenerator()

var tokenService = service.NewIdTokenService()

func BatchNext(writer http.ResponseWriter, request *http.Request) {
	if ok := canGenerate(getAuth(request)); !ok {
		resp, _ := json.Marshal(model.RetErr("403", "token异常", nil))
		writer.Write(resp)
		return
	}
	size := 20
	if sizeStr := request.FormValue("size"); len(sizeStr) > 0 {
		if i, err := strconv.Atoi(sizeStr); err != nil {
			log.Printf("size参数解析异常 %s", err)
		} else if i > 0 && i <= 100 {
			size = i
		}
	}
	ret, _ := generator.BatchNext(size)
	resp, _ := json.Marshal(model.RetOk(ret))
	writer.Write(resp)
}

func Next(writer http.ResponseWriter, request *http.Request) {
	if ok := canGenerate(getAuth(request)); !ok {
		resp, _ := json.Marshal(model.RetErr("403", "token异常", nil))
		writer.Write(resp)
		return
	}
	ret, _ := generator.Next()
	resp, _ := json.Marshal(model.RetOk(ret))
	writer.Write(resp)
}

func canGenerate(biz, token string) bool {
	return tokenService.CanGenerate(biz, token)
}

func getAuth(request *http.Request) (string, string) {
	return request.Header.Get("biz_type"), request.Header.Get("token")
}
