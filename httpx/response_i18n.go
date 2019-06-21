package httpx

const (
	DefaultI18nYamlConfig = `
zh:
  common.badRequest: 错误的请求
  common.sessionFail: 您太久没有操作页面，需要重新登录验证您的身份
  common.permissionDeny: 您没有对应的权限
  common.notfound: 没有找到对应的资源
  common.serverFail: 服务器内部错
en:
  common.badRequest: Bad request
  common.sessionFail: You have not operate the site for long time, please re-login the site again
  common.permissionDeny: You have not the corresponding permissions
  common.notfound: No corresponding resources found
  common.serverFail: Server internal error
`

	DefaultI18nTomlConfig = `
[zh]
  "common.badRequest" = "错误的请求"
  "common.sessionFail" = "您太久没有操作页面，需要重新登录验证您的身份"
  "common.permissionDeny" = "您没有对应的权限"
  "common.notfound" = "没有找到对应的资源"
  "common.serverFail" = "服务器内部错"
[en]
  "common.badRequest" = "Bad request"
  "common.sessionFail" = "You have not operate the site for long time, please re-login the site again"
  "common.permissionDeny" = "You have not the corresponding permissions"
  "common.notfound" = "No corresponding resources found"
  "common.serverFail" = "Server internal error"
`
)
