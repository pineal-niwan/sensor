package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BadRequestCode     = 400
	SessionFailCode    = 401
	PermissionDenyCode = 403
	NotFoundCode       = 404

	ServerFailCode = 500
)

const (
	BadRequestMsg     = `common.badrequest`
	SessionFailMsg    = `common.sessionfail`
	PermissionDenyMsg = `common.permissiondeny`
	NotFoundMsg       = `common.notfound`

	ServerFailMsg = `common.serverfail`
)

//回应错误的参数请求
func ResponseBadRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": BadRequestCode,
		"errMsg": GetLocale(c, GinI18n, BadRequestMsg, GinLogger),
	})
}

//回应没有登录的请求
func ResponseNeedLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": SessionFailCode,
		"errMsg": GetLocale(c, GinI18n, SessionFailMsg, GinLogger),
	})
}

//回应权限不足
func ResponsePermissionDeny(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": PermissionDenyCode,
		"errMsg": GetLocale(c, GinI18n, PermissionDenyMsg, GinLogger),
	})
}

//回应not found
func ResponseNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": NotFoundCode,
		"errMsg": GetLocale(c, GinI18n, NotFoundMsg, GinLogger),
	})
}

//回应服务器内部错
func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": ServerFailCode,
		"errMsg": GetLocale(c, GinI18n, ServerFailMsg, GinLogger),
	})
}

//回应正确的响应
func ResponseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": http.StatusOK,
		"errMsg": "",
	})
}

//回应正确的数据
func ResponseData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": http.StatusOK,
		"errMsg": "",
		"data":   data,
	})
}
