package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessReturn struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Error int         `json:"error"`
}

type ErrorReturn struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error int         `json:"error"`
}

func MakeSuccessReturn(data interface{}) (int, interface{}) {
	return http.StatusOK, SuccessReturn{
		Msg:   "success",
		Data:  data,
		Error: 0,
	}
}

func MakeErrorReturn(status int, msg string, code int, data interface{}, error int) (int, interface{}) {
	return status, ErrorReturn{
		Msg:   msg,
		Code:  code,
		Data:  data,
		Error: error,
	}
}

func Response(f func(c *gin.Context) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(f(c))
	}
}
