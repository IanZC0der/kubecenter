package response

import (
	"github.com/IanZC0der/kubecenter/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, msg string, data any) {

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
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

	c.JSON(e.HttpCode, gin.H{
		"code": e.HttpCode,
		"msg":  "failed",
		"data": e,
	})

}
