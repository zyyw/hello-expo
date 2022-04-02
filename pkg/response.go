package pkg

import (
	"github.com/gin-gonic/gin"
	e "github.com/zyyw/hello-expo/pkg/error"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetCodeMessage(errCode),
		Data: data,
	})
	return
}
