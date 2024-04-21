package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Data: data,
		Msg:  e.GetMsg(errCode),
	})
}
