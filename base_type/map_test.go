package base_type

import "testing"

//注意整数之间的转换，如果要转成小长度的无符号整数，会出问题。
//负数的浮点数转换成无符号整数会出问题。
//float64与float32转换后，会有精度问题，误差可能在0.00000002左右，小数位后8位

func TestConvertMapVal2Bool(t *testing.T) {
	key := "test"

	if !ConvertMapVal2Bool(map[string]interface{}{key: true}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: false}, key) {
		t.Fail()
	}

	k1, k2, k3, k4, k5 := int(1), int(2), int(-1), int(-2), int(0)

	if !ConvertMapVal2Bool(map[string]interface{}{key: k1}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: k2}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: k3}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: k4}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: k5}, key) {
		t.Fail()
	}

	if !ConvertMapVal2Bool(map[string]interface{}{key: uint(k1)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint(k2)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint(k3)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint(k4)}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: uint(k5)}, key) {
		t.Fail()
	}

	if !ConvertMapVal2Bool(map[string]interface{}{key: int8(k1)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int8(k2)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int8(k3)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int8(k4)}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: int8(k5)}, key) {
		t.Fail()
	}

	if !ConvertMapVal2Bool(map[string]interface{}{key: uint8(k1)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint8(k2)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint8(k3)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: uint8(k4)}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: uint8(k5)}, key) {
		t.Fail()
	}

	if !ConvertMapVal2Bool(map[string]interface{}{key: int32(k1)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int32(k2)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int32(k3)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int32(k4)}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: int32(k5)}, key) {
		t.Fail()
	}

	if !ConvertMapVal2Bool(map[string]interface{}{key: int64(k1)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int64(k2)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int64(k3)}, key) {
		t.Fail()
	}
	if !ConvertMapVal2Bool(map[string]interface{}{key: int64(k4)}, key) {
		t.Fail()
	}
	if ConvertMapVal2Bool(map[string]interface{}{key: int64(k5)}, key) {
		t.Fail()
	}
}

func TestConvertMapVal2Int(t *testing.T) {
	key := "test"

	k1, k2, k3, k4, k5 := int(1), int(2), int(-1), int(-2), int(0)

	if ConvertMapVal2Int(map[string]interface{}{key: k1}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: k2}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: k3}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: k4}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: k5}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int(map[string]interface{}{key: uint(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: uint(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: uint(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: uint(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: uint(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int(map[string]interface{}{key: int8(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int8(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int8(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int8(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int8(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int(map[string]interface{}{key: int32(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int32(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int32(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int32(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int32(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int(map[string]interface{}{key: int64(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int64(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int64(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int64(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int(map[string]interface{}{key: int64(k5)}, key) != k5 {
		t.Fail()
	}

}

func TestConvertMapVal2Int64(t *testing.T) {
	key := "test"

	k1, k2, k3, k4, k5 := int64(1), int64(2), int64(-1), int64(-2), int64(0)

	if ConvertMapVal2Int64(map[string]interface{}{key: k1}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: k2}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: k3}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: k4}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: k5}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int64(map[string]interface{}{key: uint(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: uint(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: uint(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: uint(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: uint(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int64(map[string]interface{}{key: int8(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int8(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int8(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int8(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int8(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int64(map[string]interface{}{key: int32(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int32(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int32(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int32(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int32(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int64(map[string]interface{}{key: int64(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int64(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int64(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int64(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int64(map[string]interface{}{key: int64(k5)}, key) != k5 {
		t.Fail()
	}
}

func TestConvertMapVal2Int32(t *testing.T) {
	key := "test"

	k1, k2, k3, k4, k5 := int32(1), int32(2), int32(-1), int32(-2), int32(0)

	if ConvertMapVal2Int32(map[string]interface{}{key: k1}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: k2}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: k3}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: k4}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: k5}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int32(map[string]interface{}{key: uint(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: uint(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: uint(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: uint(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: uint(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int32(map[string]interface{}{key: int8(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int8(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int8(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int8(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int8(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int32(map[string]interface{}{key: int32(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int32(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int32(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int32(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int32(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Int32(map[string]interface{}{key: int64(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int64(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int64(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int64(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Int32(map[string]interface{}{key: int64(k5)}, key) != k5 {
		t.Fail()
	}
}

func TestConvertMapVal2Float64(t *testing.T) {
	key := "test"

	k1, k2, k3, k4, k5 := float64(1), float64(2), float64(-1), float64(-2), float64(0)

	if ConvertMapVal2Float64(map[string]interface{}{key: k1}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: k2}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: k3}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: k4}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: k5}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Float64(map[string]interface{}{key: int8(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int8(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int8(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int8(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int8(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Float64(map[string]interface{}{key: int32(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int32(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int32(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int32(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int32(k5)}, key) != k5 {
		t.Fail()
	}

	if ConvertMapVal2Float64(map[string]interface{}{key: int64(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int64(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int64(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int64(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: int64(k5)}, key) != k5 {
		t.Fail()
	}
}

func TestConvertMapVal2Float64_2(t *testing.T) {
	key := "test"

	k1, k2, k3, k4, k5 := float64(1), float64(2), float64(-1), float64(-2), float64(0)

	if ConvertMapVal2Float64(map[string]interface{}{key: float32(k1)}, key) != k1 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: float32(k2)}, key) != k2 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: float32(k3)}, key) != k3 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: float32(k4)}, key) != k4 {
		t.Fail()
	}
	if ConvertMapVal2Float64(map[string]interface{}{key: float32(k5)}, key) != k5 {
		t.Fail()
	}

	k1, k2, k3, k4, k5 = float64(1.1), float64(2.1), float64(-1.1), float64(-2.1), float64(0.0)
	if !FloatEquals(ConvertMapVal2Float64(map[string]interface{}{key: float32(k1)}, key), k1) {
		t.Fail()
	}
	if !FloatEquals(ConvertMapVal2Float64(map[string]interface{}{key: float32(k2)}, key), k2) {
		t.Fail()
	}
	if !FloatEquals(ConvertMapVal2Float64(map[string]interface{}{key: float32(k3)}, key), k3) {
		t.Fail()
	}
	if !FloatEquals(ConvertMapVal2Float64(map[string]interface{}{key: float32(k4)}, key), k4) {
		t.Fail()
	}
	if !FloatEquals(ConvertMapVal2Float64(map[string]interface{}{key: float32(k5)}, key), k5) {
		t.Fail()
	}

}

func TestConvertMapVal2String(t *testing.T) {
	key := "test"
	if ConvertMapVal2String(map[string]interface{}{key: key}, key) != key {
		t.Fail()
	}
	if ConvertMapVal2String(map[string]interface{}{key: []byte(key)}, key) != key {
		t.Fail()
	}
}
