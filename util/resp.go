package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func OKResp(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, res)
}

func RespFail(c *gin.Context, msg string) {
	Resp(c, -1, nil, msg)
}
func RespOk(c *gin.Context, data interface{}, msg string) {
	Resp(c, 0, data, msg)
}
func RespOkList(c *gin.Context, lists interface{}, total interface{}) {
	RespList(c, 0, lists, total)
}
func Resp(c *gin.Context, code int, data interface{}, msg string) {
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, h)
}
func RespList(c *gin.Context, code int, data interface{}, total interface{}) {
	h := &H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	c.JSON(http.StatusOK, h)
}
