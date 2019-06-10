package base_type

import (
	"fmt"
	"strconv"
)

//将基本类型转换为字符串
func Convert2String(v interface{}) string {

	//检查整数
	bInteger := true
	var iv int
	switch v.(type) {
	case int8:
		iv = int(v.(int8))
	case uint8:
		iv = int(v.(uint8))
	case int16:
		iv = int(v.(int16))
	case uint16:
		iv = int(v.(uint16))
	case int32:
		iv = int(v.(int32))
	case uint32:
		iv = int(v.(uint32))
	case int64:
		iv = int(v.(int64))
	case uint64:
		iv = int(v.(uint64))
	case int:
		iv = v.(int)
	case uint:
		iv = int(v.(uint))
	default:
		bInteger = false
	}
	if bInteger {
		return strconv.Itoa(iv)
	}

	//检查字符串
	bString := true
	var sv string
	switch v.(type) {
	case string:
		sv = v.(string)
	case ([]byte):
		sv = string(v.([]byte))
	default:
		bString = false
	}
	if bString {
		return sv
	}

	//检查浮点数
	bFloat := true
	var fv float64
	switch v.(type) {
	case float32:
		fv = float64(v.(float32))
	case float64:
		fv = v.(float64)
	default:
		bFloat = false
	}
	if bFloat {
		iv = int(fv)
		return strconv.FormatFloat(fv, 'g', 3, 64)
	}

	return fmt.Sprint(v)
}

//将基本类型转换为整型
func Convert2Int(v interface{}) int {
	var iv int
	switch v.(type) {
	case int8:
		iv = int(v.(int8))
	case uint8:
		iv = int(v.(uint8))
	case int16:
		iv = int(v.(int16))
	case uint16:
		iv = int(v.(uint16))
	case int32:
		iv = int(v.(int32))
	case uint32:
		iv = int(v.(uint32))
	case int64:
		iv = int(v.(int64))
	case uint64:
		iv = int(v.(uint64))
	case int:
		iv = v.(int)
	case uint:
		iv = int(v.(uint))
	case float32:
		iv = int(v.(float32))
	case float64:
		iv = int(v.(float64))
	case string:
		iv, _ = strconv.Atoi(v.(string))
	default:
		iv = 0
	}
	return iv
}

//将基本类型转换为32位整型
func Convert2Int32(v interface{}) int32 {
	var iv int32
	switch v.(type) {
	case int8:
		iv = int32(v.(int8))
	case uint8:
		iv = int32(v.(uint8))
	case int16:
		iv = int32(v.(int16))
	case uint16:
		iv = int32(v.(uint16))
	case int32:
		iv = v.(int32)
	case uint32:
		iv = int32(v.(uint32))
	case int64:
		iv = int32(v.(int64))
	case uint64:
		iv = int32(v.(uint64))
	case int:
		iv = int32(v.(int))
	case uint:
		iv = int32(v.(uint))
	case float32:
		iv = int32(v.(float32))
	case float64:
		iv = int32(v.(float64))
	case string:
		iiv, _ := strconv.Atoi(v.(string))
		iv = int32(iiv)
	default:
		iv = 0
	}
	return iv
}

//将基本类型转换为32位整型
func Convert2Int64(v interface{}) int64 {
	var iv int64
	switch v.(type) {
	case int8:
		iv = int64(v.(int8))
	case uint8:
		iv = int64(v.(uint8))
	case int16:
		iv = int64(v.(int16))
	case uint16:
		iv = int64(v.(uint16))
	case int32:
		iv = int64(v.(int32))
	case uint32:
		iv = int64(v.(uint32))
	case int64:
		iv = v.(int64)
	case uint64:
		iv = int64(v.(uint64))
	case int:
		iv = int64(v.(int))
	case uint:
		iv = int64(v.(uint))
	case float32:
		iv = int64(v.(float32))
	case float64:
		iv = int64(v.(float64))
	case string:
		iiv, _ := strconv.Atoi(v.(string))
		iv = int64(iiv)
	default:
		iv = 0
	}
	return iv
}

//将基本类型转换为浮点型
func Convert2Float64(v interface{}) float64 {
	var iv float64
	switch v.(type) {
	case int8:
		iv = float64(v.(int8))
	case uint8:
		iv = float64(v.(uint8))
	case int16:
		iv = float64(v.(int16))
	case uint16:
		iv = float64(v.(uint16))
	case int32:
		iv = float64(v.(int32))
	case uint32:
		iv = float64(v.(uint32))
	case int64:
		iv = float64(v.(int64))
	case uint64:
		iv = float64(v.(uint64))
	case int:
		iv = float64(v.(int))
	case uint:
		iv = float64(v.(uint))
	case float32:
		iv = float64(v.(float32))
	case float64:
		iv = v.(float64)
	case string:
		iv, _ = strconv.ParseFloat(v.(string), 64)
	default:
		iv = 0
	}
	return iv
}

//浮点数近似相等
func FloatEquals(x, y float64) bool {
	delta := x - y
	if delta > 0.000001 || delta < -0.000001 {
		return false
	}
	return true
}
