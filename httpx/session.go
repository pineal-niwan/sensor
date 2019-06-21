package httpx

import (
	"crypto/aes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pineal-niwan/sensor/logger"
	"net/http"
)

var (
	CookieCodec = securecookie.New([]byte(`zACJq*(lFao11n&@lt)#$qoGNHu3zjo6!`),
		[]byte(`6!q*(lWQ1T8P$q1zj1ao7n3RS&@lt9)#`))
)

//编码session
func encodeSession(key string, hashKeyWithTime HashKeyWithTime) (hashKey string, err error) {
	return CookieCodec.Encode(key, hashKeyWithTime)
}

//解码session
func decodeSession(key string, hashKey string) (hashKeyWithTime HashKeyWithTime, err error) {
	err = CookieCodec.Decode(key, hashKey, &hashKeyWithTime)
	return
}

//设置codec
//其中blockKey长度需要为16/24/32
func SetCookieCodec(hashKey string, blockKey string) error {
	l := len(blockKey)
	switch l {
	default:
		return aes.KeySizeError(l)
	case 16, 24, 32:
		CookieCodec = securecookie.New([]byte(hashKey), []byte(blockKey))
	}
	return nil
}

//获取session
func GetSessionFromKey(c *gin.Context, iLogger logger.ILogger, key string) (HashKeyWithTime, error) {
	var sessionValue HashKeyWithTime
	cookies, err := c.Request.Cookie(key)
	if err != nil {
		iLogger.Errorf("获取session key失败:%+v", err)
		return sessionValue, err
	}

	sessionValue, err = decodeSession(key, cookies.Value)
	if err != nil {
		iLogger.Errorf("解码ession key失败:%+v", err)
		return sessionValue, err
	}
	return sessionValue, err
}

//设置session
func SetSessionWithKey(c *gin.Context, key string, hashKeyWithTime HashKeyWithTime) error {
	encode, err := encodeSession(key, hashKeyWithTime)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     key,
		Value:    encode,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	return nil
}
