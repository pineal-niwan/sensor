package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/pineal/sensor/i18n"
	"github.com/pineal/sensor/logger"
	"github.com/xeipuuv/gojsonschema"
	"time"
)

const (
	LangKey = "lang"
	DataKey = "data"
)

//Hash Key
type HashKeyWithTime struct {
	HashKey string
	TimeTTL time.Time
}

//获取i18n中对应语言的翻译
func GetLocale(c *gin.Context, i18n *i18n.LangStringGroup, key string, logger logger.Logger) string {
	if i18n == nil {
		logger.Error("i18n为空")
		return key
	}
	lang, ok := c.Get(LangKey)
	if !ok {
		//上下文中不带语言项
		defaultLang := i18n.GetDefaultLang()
		langString, _ := i18n.GetLocale(defaultLang, key)
		if langString == key {
			logger.Errorf("lang:%+v key没有翻译 %+v", defaultLang, key)
		}
		return langString
	} else {
		langLocal, ok := lang.(string)
		if !ok {
			//语言项有问题
			langString, _ := i18n.GetLocale(i18n.GetDefaultLang(), key)
			logger.Errorf("错误的LangKey %+v", langLocal)
			return langString
		} else {
			//查找翻译
			langString, _ := i18n.GetLocale(langLocal, key)
			if langString == key {
				logger.Errorf("lang:%+v key没有翻译 %+v", langLocal, key)
			}
			return langString
		}
	}
}

//设置post中的json请求数据
func SetPostJsonDataFromReq(c *gin.Context, data interface{}, logger logger.Logger) {
	err := c.BindJSON(data)
	if err != nil {
		ResponseBadRequest(c)
		c.Abort()
		logger.Errorf("错误的json请求%+v", err)
	} else {
		c.Set(DataKey, data)
	}
}

//生成json schema
func GenJsonSchema(schema interface{}) (*gojsonschema.Schema, error) {
	loader := gojsonschema.NewGoLoader(schema)
	return gojsonschema.NewSchema(loader)
}

//设置post中的json请求数据并带json schema检查
func SetPostJsonDataWithSchemaFromReq(c *gin.Context, schema *gojsonschema.Schema, logger logger.Logger) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		logger.Errorf("错误的json请求%+v", err)
		ResponseBadRequest(c)
		c.Abort()
		return
	} else {
		if schema != nil {
			loader := gojsonschema.NewGoLoader(data)
			ret, err := schema.Validate(loader)
			if err != nil {
				logger.Error(`验证json schema出错:%+v`, err)
				ResponseBadRequest(c)
				c.Abort()
				return
			}
			if !ret.Valid() {
				logger.Error(`验证json schema出错内容:%+v`, ret.Errors())
				ResponseBadRequest(c)
				c.Abort()
				return
			}
		}
		c.Set(DataKey, data)
	}
}

//获取post中的请求数据
func GetPostJsonDataFromReq(c *gin.Context) (interface{}, bool) {
	data, ok := c.Get(DataKey)
	return data, ok
}
