package httpx

import (
	"github.com/gin-gonic/gin"
	"pineal/sensor/i18n"
	"pineal/sensor/logger"
)

const (
	LangKey = "lang"
	DataKey = "data"
)

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
