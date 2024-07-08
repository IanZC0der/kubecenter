package response

import (
	"github.com/IanZC0der/kubecenter/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data any) {

	c.JSON(http.StatusOK, data)

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

	c.JSON(e.HttpCode, e)

}
