package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BadRequestCode = 400
	ServerFail = 500

)

const (
	BadRequestMsg = `common.badrequest`
	ServerFailMsg = `common.serverfail`
)

//回应错误的参数请求
func ResponseBadRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": BadRequestCode,
		"errMsg": GetLocale(c, GinI18n, BadRequestMsg, GinLogger),
	})
}

//回应服务器内部错
func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": ServerFail,
		"errMsg": GetLocale(c, GinI18n, ServerFailMsg, GinLogger),
	})
}