package httpx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/pineal-niwan/sensor/cache"
	"github.com/pineal-niwan/sensor/logger"
	"net/http"
	"net/http/httputil"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// gin handler结构体
type GinHandler struct {
	*gin.Engine
	logger.ILogger
	prefix string
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
func NewGinHandler(iLogger logger.ILogger) *GinHandler {
	var ginLogger logger.ILogger

	if iLogger == nil {
		ginLogger = logger.DefaultLogger
	} else {
		ginLogger = iLogger
	}

	ginHandler := &GinHandler{
		Engine:  gin.New(),
		ILogger: ginLogger,
	}
	ginHandler.Use(ginHandler.Logger)
	ginHandler.Use(ginHandler.Recovery)
	return ginHandler
}

//添加前缀
func (g *GinHandler) SetPrefix(prefix string) {
	g.prefix = prefix
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
func (g *GinHandler) GET(url string, handlerList ...HandleGinUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.GET(url, ginHandlerFuncList...)
}

//添加POST处理函数
func (g *GinHandler) POST(url string, handlerList ...HandleGinUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.POST(url, ginHandlerFuncList...)
}

// gin handler带cache service和db
type GinDataHandler struct {
	*GinHandler
	db          *gorm.DB
	cacheClient cache.IStringKeyCacheClient
}

//新建
func NewGinDataHandler(db *gorm.DB, cacheClient cache.IStringKeyCacheClient, iLogger logger.ILogger) *GinDataHandler {
	ginDataHandler := &GinDataHandler{}
	ginDataHandler.GinHandler = NewGinHandler(iLogger)
	ginDataHandler.db = db
	ginDataHandler.cacheClient = cacheClient
	return ginDataHandler
}

//带logger的http处理函数
type HandleGinDataUrlFunc func(c *gin.Context, db *gorm.DB, cacheClient cache.IStringKeyCacheClient, iLogger logger.ILogger)

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
func (g *GinDataHandler) GET(url string, handlerList ...HandleGinDataUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.GET(url, ginHandlerFuncList...)
}

//添加POST处理函数
func (g *GinDataHandler) POST(url string, handlerList ...HandleGinDataUrlFunc) {
	groupGin := g.Group(g.prefix)
	ginHandlerFuncList := g.convertHandler(handlerList)
	groupGin.POST(url, ginHandlerFuncList...)
}
