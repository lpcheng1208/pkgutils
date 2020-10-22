package pkg

import "encoding/json"

type CommonResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCommonResponse() *CommonResponse {
	return &CommonResponse{
		Code: 0,
		Msg:  "ok",
		Data: json.RawMessage("{}"),
	}
}