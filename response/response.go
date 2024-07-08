package response

import (
	"github.com/IanZC0der/kubecenter/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, msg string, data any) {

	c.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})

}

func Failed(c *gin.Context, err error) {
	defer c.Abort()
	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	c.JSON(e.HttpCode, &Response{
		Code: e.HttpCode,
		Msg:  "failed",
		Data: e,
	})

}
