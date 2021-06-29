package model

type Ret struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RetOk(data interface{}) *Ret {
	return &Ret{
		Code: "0",
		Msg:  "",
		Data: data,
	}
}
