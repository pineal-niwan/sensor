package httpx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/pineal-niwan/sensor/i18n"
	"github.com/pineal-niwan/sensor/logger"
	"net/http"
	"net/http/httputil"
	"time"
)

var (
	GinLogger logger.ILogger
	GinI18n   *i18n.LangStringGroup
)

var (
	//gin模式错误
	ErrHttpXInvalidMode = errors.New("httpx invalid mode")
	//没有设置logger
	ErrHttpXLoggerNil = errors.New("nil httpx logger")
)

//记录日志
func Logger(c *gin.Context) {
	if GinLogger.GetLevel() >= logger.InfoLevel {
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

		GinLogger.Infof("[GIN] %v | %3d | %13v | %s |%s  %s",
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
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			goErr := errors.Wrap(err, 3)
			reset := string([]byte{27, 91, 48, 109})
			errMsg := fmt.Sprintf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s",
				httprequest, goErr.Error(), goErr.Stack(), reset)
			GinLogger.Errorf(errMsg)
			//回调
			c.AbortWithStatus(http.StatusInternalServerError)
		}

	}()
	c.Next()
}

//新建gin handler
func newGinHandler(mode string) *gin.Engine {
	gin.SetMode(mode)
	handler := gin.New()
	handler.Use(Logger)
	handler.Use(Recovery)
	return handler
}

//新建gin handler
func NewGinHandler(mode string, ginLogger logger.ILogger, prefix string, defaultLang string, i18nFiles ...string) (
	*gin.RouterGroup, error) {

	//设置模式
	if mode != gin.DebugMode && mode != gin.ReleaseMode && mode != gin.TestMode {
		return nil, ErrHttpXInvalidMode
	}

	//设置翻译
	GinI18n := &i18n.LangStringGroup{}
	GinI18n.Init(defaultLang)
	defaultI18nConfig, err := i18n.LoadI18nConfigFromTomlBuffer([]byte(DefaultI18nTomlConfig))
	if err != nil {
		return nil, err
	}
	GinI18n.ConfigBy(defaultI18nConfig)
	for _, i18nFile := range i18nFiles {
		i18nConfig, err := i18n.LoadI18nConfigFromYamlFile(i18nFile)
		if err != nil {
			return nil, err
		}
		GinI18n.ConfigBy(i18nConfig)
	}

	//设置logger
	if ginLogger == nil {
		return nil, ErrHttpXLoggerNil
	}
	GinLogger = ginLogger

	//设置handler及url前缀
	handler := newGinHandler(mode)
	handlerGrp := handler.Group(prefix)
	return handlerGrp, nil
}
