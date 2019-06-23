package httpx

const (
	//错误的请求
	BadRequestCode = 400
	//需要重新登录
	SessionFailCode = 401
	//权限不够
	PermissionDenyCode = 403
	//没有找到对应的资源
	NotFoundCode = 404
	//服务器内部错误
	ServerFailCode = 500
)

const (
	//错误的请求
	BadRequestMsg = `common.badRequest`
	//需要重新登录
	SessionFailMsg = `common.sessionFail`
	//权限不够
	PermissionDenyMsg = `common.permissionDeny`
	//没有找到对应的资源
	NotFoundMsg = `common.notfound`
	//服务器内部错误
	ServerFailMsg = `common.serverFail`
)

var (
	//空hash
	__emptyHash = make(map[string]interface{})
)