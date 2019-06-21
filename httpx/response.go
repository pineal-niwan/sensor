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
	BadRequestMsg     = `common.badRequest`
	SessionFailMsg    = `common.sessionFail`
	PermissionDenyMsg = `common.permissionDeny`
	NotFoundMsg       = `common.notfound`

	ServerFailMsg = `common.serverFail`
)

//回应错误的参数请求
func ResponseBadRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": BadRequestCode,
		"errMsg": BadRequestMsg,
	})
}

//回应没有登录的请求
func ResponseNeedLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": SessionFailCode,
		"errMsg": SessionFailMsg,
	})
}

//回应权限不足
func ResponsePermissionDeny(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": PermissionDenyCode,
		"errMsg": PermissionDenyMsg,
	})
}

//回应not found
func ResponseNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": NotFoundCode,
		"errMsg": NotFoundMsg,
	})
}

//回应服务器内部错
func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": ServerFailCode,
		"errMsg": ServerFailMsg,
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
