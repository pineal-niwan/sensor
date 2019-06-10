package json

import "github.com/gin-gonic/gin"

//获取json参数
func Json(c *gin.Context) {
	var mData map[string]interface{}
	err := c.BindJSON(&mData)
	if err != nil {
		gin.DefaultWriter.Debug("错误的json请求", err)
		ResponseBadRequest(c)
		c.Abort()
		return
	} else {
		key := getValidatePath(c)
		if SchemaHandler.Validate(key, mData) <= json_handler.VAL_OK {
			c.Set(REQ_DATA_KEY, mData)
		} else {
			logrus.Debug("jsonschema检查出错了")
			ResponseBadRequest(c)
			c.Abort()
		}
	}
}