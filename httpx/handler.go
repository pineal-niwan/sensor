package httpx

import (
	"crypto/aes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	"github.com/pineal-niwan/sensor/cache"
	"github.com/pineal-niwan/sensor/httpx/consts"
	"github.com/pineal-niwan/sensor/logger"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"net/http/httputil"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

//Hash Key
type HashKeyWithTime struct {
	HashKey string
	TimeTTL time.Time
}

var (
	//缺省的cookie codec
	defaultCookieCodec = securecookie.New([]byte(`zACJq*(lFao11n&@lt)#$qoGNHu3zjo6!`),
		[]byte(`6!q*(lWQ1T8P$q1zj1ao7n3RS&@lt9)#`))
	//空hash
	__emptyHash = make(map[string]interface{})
)

// gin handler结构体
type GinHandler struct {
	*gin.Engine
	logger.ILogger
	prefix         string
	dataKey        string
	sessionKey     string
	cookieCodec    *securecookie.SecureCookie
	jsonSchemaHash map[string]*gojsonschema.Schema
}

//记录日志
func (g *GinHandler) Logger(c *gin.Context) {
	if g.ILogger.GetLevel() >= logger.InfoLevel {
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		g.ILogger.Infof("[GIN] %v | %3d | %13v | %s |%s  %s",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}

//恢复
func (g *GinHandler) Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			goErr := errors.Wrap(err, 3)
			reset := string([]byte{27, 91, 48, 109})
			errMsg := fmt.Sprintf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s",
				httprequest, goErr.Error(), goErr.Stack(), reset)
			g.ILogger.Errorf(errMsg)
			//回调
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}

//新建gin handler
func NewGinHandler(iLogger logger.ILogger, jsonSchemaHash map[string]*gojsonschema.Schema) *GinHandler {
	var ginLogger logger.ILogger

	if iLogger == nil {
		ginLogger = logger.DefaultLogger
	} else {
		ginLogger = iLogger
	}

	newJsonSchemaHash := jsonSchemaHash
	if newJsonSchemaHash == nil {
		newJsonSchemaHash = make(map[string]*gojsonschema.Schema)
	}

	ginHandler := &GinHandler{
		Engine:         gin.New(),
		ILogger:        ginLogger,
		dataKey:        `data`,
		jsonSchemaHash: newJsonSchemaHash,
	}
	ginHandler.Use(ginHandler.Logger)
	ginHandler.Use(ginHandler.Recovery)
	return ginHandler
}

//添加前缀
func (g *GinHandler) SetPrefix(prefix string) {
	g.prefix = prefix
}

//设置session key
func (g *GinHandler) SetSessionKey(sessionKey string) {
	g.sessionKey = sessionKey
}

//回应错误的参数请求
func (g *GinHandler) ResponseBadRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: consts.BadRequestCode,
		consts.NameErrMsg: consts.BadRequestMsg,
	})
}

//回应没有登录的请求
func (g *GinHandler) ResponseNeedLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: consts.SessionFailCode,
		consts.NameErrMsg: consts.SessionFailMsg,
	})
}

//回应权限不足
func (g *GinHandler) ResponsePermissionDeny(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: consts.PermissionDenyCode,
		consts.NameErrMsg: consts.PermissionDenyMsg,
	})
}

//回应not found
func (g *GinHandler) ResponseNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: consts.NotFoundCode,
		consts.NameErrMsg: consts.NotFoundMsg,
	})
}

//回应服务器内部错
func (g *GinHandler) ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: consts.ServerFailCode,
		consts.NameErrMsg: consts.ServerFailMsg,
	})
}

//回应传入的result和error message
func (g *GinHandler) ResponseErrMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: code,
		consts.NameErrMsg: msg,
	})
}

//回应正确的响应
func (g *GinHandler) ResponseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: http.StatusOK,
		consts.NameErrMsg: "",
	})
}

//回应正确的数据
func (g *GinHandler) ResponseData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		consts.NameResult: http.StatusOK,
		consts.NameErrMsg: "",
		"data":   data,
	})
}

//使用json schema解析数据
func (g *GinHandler) ParseJson(c *gin.Context, schemaKey string) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		g.ResponseBadRequest(c)
		c.Abort()
		g.ILogger.Errorf("错误的json请求:%+v 错误信息:%+v", c.Request, err)
		return
	} else {
		if schemaKey != "" {
			schema := g.jsonSchemaHash[schemaKey]
			if schema != nil {
				loader := gojsonschema.NewGoLoader(data)
				ret, err := schema.Validate(loader)
				if err != nil {
					g.ResponseBadRequest(c)
					c.Abort()
					g.ILogger.Errorf("错误的json请求:%+v 错误信息:%+v", c.Request, err)
					g.ILogger.Errorf(`验证json schema出错:%+v`, err)
					return
				} else {
					if !ret.Valid() {
						g.ResponseBadRequest(c)
						c.Abort()
						g.ILogger.Errorf("错误的json请求:%+v 错误信息:%+v", c.Request, err)
						g.ILogger.Errorf(`验证json schema出错内容:%+v`, ret.Errors())
						return
					}
				}
			}
		}
		c.Set(g.dataKey, data)
	}
}

//获取解析的json data
func (g *GinHandler) GetData(c *gin.Context) (map[string]interface{}, bool) {
	iData, ok := c.Get(g.dataKey)
	if !ok {
		return __emptyHash, false
	} else {
		data, ok := iData.(map[string]interface{})
		return data, ok
	}
}

//编码session
func (g *GinHandler) encodeSession(key string, hashKeyWithTime HashKeyWithTime) (hashKey string, err error) {
	cookieCodec := defaultCookieCodec
	if g.cookieCodec != nil {
		cookieCodec = g.cookieCodec
	}
	return cookieCodec.Encode(key, hashKeyWithTime)
}

//解码session
func (g *GinHandler) decodeSession(key string, hashKey string) (hashKeyWithTime HashKeyWithTime, err error) {
	cookieCodec := defaultCookieCodec
	if g.cookieCodec != nil {
		cookieCodec = g.cookieCodec
	}
	err = cookieCodec.Decode(key, hashKey, &hashKeyWithTime)
	return
}

//设置codec
//其中blockKey长度需要为16/24/32
func (g *GinHandler) SetCookieCodec(hashKey string, blockKey string) error {
	l := len(blockKey)
	switch l {
	default:
		return aes.KeySizeError(l)
	case 16, 24, 32:
		g.cookieCodec = securecookie.New([]byte(hashKey), []byte(blockKey))
	}
	return nil
}

//获取session
func (g *GinHandler) GetSessionFromKey(c *gin.Context) (HashKeyWithTime, error) {
	var sessionValue HashKeyWithTime
	cookies, err := c.Request.Cookie(g.sessionKey)
	if err != nil {
		g.ILogger.Errorf("获取session key失败:%+v", err)
		return sessionValue, err
	}

	sessionValue, err = g.decodeSession(g.sessionKey, cookies.Value)
	if err != nil {
		g.ILogger.Errorf("解码session key失败:%+v", err)
		return sessionValue, err
	}
	return sessionValue, err
}

//设置session
func (g *GinHandler) SetSessionWithKey(c *gin.Context, hashKeyWithTime HashKeyWithTime) error {
	encode, err := g.encodeSession(g.sessionKey, hashKeyWithTime)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     g.sessionKey,
		Value:    encode,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	return nil
}

//添加GET处理函数
func (g *GinHandler) GET(url string, handlers ...gin.HandlerFunc) {
	groupGin := g.Group(g.prefix)
	groupGin.GET(url, handlers...)
}

//添加POST处理函数
func (g *GinHandler) POST(url string, handlers ...gin.HandlerFunc) {
	groupGin := g.Group(g.prefix)
	groupGin.POST(url, handlers...)
}

//带logger的http处理函数
type HandleGinUrlFunc func(c *gin.Context, iLogger logger.ILogger)

//利用闭包，转换handler函数，加入logger支持
func (g *GinHandler) convertHandler(handlerList []HandleGinUrlFunc) []gin.HandlerFunc {
	chainLen := len(handlerList)
	ginHandlerFuncList := make([]gin.HandlerFunc, chainLen)
	for i := 0; i < chainLen; i++ {
		ginHandlerFuncList[i] = func(c *gin.Context) {
			handlerList[i](c, g.ILogger)
		}
	}
	return ginHandlerFuncList
}

//添加GET处理函数
func (g *GinHandler) GETx(url string, handlerList ...HandleGinUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.GET(url, ginHandlerFuncList...)
}

//添加POST处理函数
func (g *GinHandler) POSTx(url string, handlerList ...HandleGinUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.POST(url, ginHandlerFuncList...)
}

// gin handler带cache service和db
type GinDataHandler struct {
	*GinHandler
	db          *gorm.DB
	cacheClient cache.IStringAsKeyCacheClient
}

//新建
func NewGinDataHandler(db *gorm.DB, cacheClient cache.IStringAsKeyCacheClient,
	iLogger logger.ILogger, jsonSchema map[string]*gojsonschema.Schema) *GinDataHandler {

	ginDataHandler := &GinDataHandler{}
	ginDataHandler.GinHandler = NewGinHandler(iLogger, jsonSchema)
	ginDataHandler.db = db
	ginDataHandler.cacheClient = cacheClient
	return ginDataHandler
}

//带logger的http处理函数
type HandleGinDataUrlFunc func(c *gin.Context, db *gorm.DB, cacheClient cache.IStringAsKeyCacheClient, iLogger logger.ILogger)

//利用闭包，转换handler函数，加入logger支持
func (g *GinDataHandler) convertHandler(handlerList []HandleGinDataUrlFunc) []gin.HandlerFunc {
	chainLen := len(handlerList)
	ginHandlerFuncList := make([]gin.HandlerFunc, chainLen)
	for i := 0; i < chainLen; i++ {
		ginHandlerFuncList[i] = func(c *gin.Context) {
			handlerList[i](c, g.db, g.cacheClient, g.ILogger)
		}
	}
	return ginHandlerFuncList
}

//添加GET处理函数
func (g *GinDataHandler) GETx(url string, handlerList ...HandleGinDataUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.GET(url, ginHandlerFuncList...)
}

//添加POST处理函数
func (g *GinDataHandler) POSTx(url string, handlerList ...HandleGinDataUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.POST(url, ginHandlerFuncList...)
}
