package sql_convert

import (
	"github.com/go-errors/errors"
	"github.com/pineal-niwan/sensor/json"
)

var (
	//不支持json序列化的类型
	ErrNotJsonSupportType = errors.New(`the data type is not support json`)
)

type Map map[string]interface{}

//json对应的map
type JsonMap struct {
	Map
}

//从数据库读取后组装
func (jsonMap *JsonMap) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		if src == "" {
			jsonMap.Map = make(Map)
			return nil
		} else {
			return json.Unmarshal([]byte(src), &jsonMap.Map)
		}
	case []byte:
		if len(src) == 0 {
			jsonMap.Map = make(Map)
			return nil
		} else {
			return json.Unmarshal(src, &jsonMap.Map)
		}
	default:
		return ErrNotJsonSupportType
	}
}
