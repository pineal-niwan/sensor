package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/pineal-niwan/sensor/logger"
	"github.com/xeipuuv/gojsonschema"
	"time"
)

const (
	DataKey = "data"
)

//Hash Key
type HashKeyWithTime struct {
	HashKey string
	TimeTTL time.Time
}

//设置post中的json请求数据
func SetPostJsonDataFromReq(c *gin.Context, iLogger logger.ILogger, data interface{}) {
	err := c.BindJSON(data)
	if err != nil {
		ResponseBadRequest(c)
		c.Abort()
		iLogger.Errorf("错误的json请求%+v", err)
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
func SetPostJsonDataWithSchemaFromReq(c *gin.Context, iLogger logger.ILogger, schema *gojsonschema.Schema) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		iLogger.Errorf("错误的json请求%+v", err)
		ResponseBadRequest(c)
		c.Abort()
		return
	} else {
		if schema != nil {
			loader := gojsonschema.NewGoLoader(data)
			ret, err := schema.Validate(loader)
			if err != nil {
				iLogger.Error(`验证json schema出错:%+v`, err)
				ResponseBadRequest(c)
				c.Abort()
				return
			}
			if !ret.Valid() {
				iLogger.Error(`验证json schema出错内容:%+v`, ret.Errors())
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
