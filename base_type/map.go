package base_type

import (
	"time"
)

//Map中是否包含了key
func IsKeyInMap(m map[string]interface{}, k string) bool {
	_, ok := m[k]
	return ok
}

//将map中的value转换为bool
func ConvertMapVal2Bool(m map[string]interface{}, k string) bool {
	x, ok := m[k]
	if !ok {
		return false
	}
	switch x.(type) {
	case bool:
		return x.(bool)
	}
	intVal := Convert2Int(x)
	if intVal != 0 {
		return true
	}
	return false
}

//将map中的value转换为int
func ConvertMapVal2Int(m map[string]interface{}, k string) int {
	x, ok := m[k]
	if !ok {
		return 0
	}
	return Convert2Int(x)
}

//将map中的value转换为int64
func ConvertMapVal2Int64(m map[string]interface{}, k string) int64 {
	x, ok := m[k]
	if !ok {
		return 0
	}
	return Convert2Int64(x)
}

//将map中的value转换为int32
func ConvertMapVal2Int32(m map[string]interface{}, k string) int32 {
	x, ok := m[k]
	if !ok {
		return 0
	}
	return Convert2Int32(x)
}

//将map中的value转换为float64
func ConvertMapVal2Float64(m map[string]interface{}, k string) float64 {
	x, ok := m[k]
	if !ok {
		return 0
	}
	return Convert2Float64(x)
}

//将map中的value转换为string
func ConvertMapVal2String(m map[string]interface{}, k string) string {
	x, ok := m[k]
	if !ok {
		return ""
	}
	return Convert2String(x)
}

//将map中的value转换为time
func ConvertMapVal2Time(m map[string]interface{}, k string) (time.Time, bool) {
	x, ok := m[k]
	if !ok {
		return time.Time{}, false
	}
	return Convert2Time(x)
}
