package controllers

import (
	"github.com/astaxie/beego"
	"movie/library"
)

type ResponseController struct {
	beego.Controller
}

type OutputResponse struct {
	ErrNo  int         `json:"errNo"`
	ErrStr string      `json:"errStr"`
	Data   interface{} `json:"data"`
}

func ReturnError(errNo int,) *OutputResponse {
	return &OutputResponse{
		ErrNo:  errNo,
		ErrStr: library.ErrMap(errNo),
		Data:   nil,
	}
}

func ReturnResponse(data interface{}) *OutputResponse {
	return &OutputResponse{
		ErrNo:  0,
		ErrStr: "success",
		Data:   data,
	}
}

func ReturnStr(errStr string) *OutputResponse {
	return &OutputResponse{
		ErrNo:  -1,
		ErrStr: errStr,
		Data:   nil,
	}
}
