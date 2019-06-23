package sql_convert

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/go-errors/errors"
	"time"
)

var (
	ErrTimestampIsNotInt64 = errors.New("time stamp is not int64")
	_UnixZeroTimeStamp     = time.Unix(0, 0)
)

//秒级的unix timestamp
type UnixStamp struct {
	time.Time
}

//新建unix stamp
func NewUnixStamp(t time.Time) UnixStamp {
	return UnixStamp{t}
}

//新建zero值的timestamp
func NewUnixZeroStamp() UnixStamp {
	var unixStamp UnixStamp
	unixStamp.Time = _UnixZeroTimeStamp
	return unixStamp
}

//从数据库读取后组装
func (timestamp *UnixStamp) Scan(src interface{}) error {
	switch src := src.(type) {
	case int64:
		timestamp.Time = time.Unix(src, 0)
		return nil
	default:
		timestamp.Time = time.Time{}
		return ErrTimestampIsNotInt64
	}
}

//写回数据库时的组装
func (timestamp *UnixStamp) Value() (driver.Value, error) {
	timestampUnix := timestamp.Unix()
	return timestampUnix, nil
}

//Json marshal
func (timeStamp UnixStamp) MarshalJSON() ([]byte, error) {
	timestampUnix := timeStamp.Unix()
	return json.Marshal(timestampUnix)
}

//Json marshal
func (timeStamp *UnixStamp) UnmarshalJSON(data []byte) error {
	var timestampUnix int64
	err := json.Unmarshal(data, &timestampUnix)
	if err != nil {
		return err
	}
	timeStamp.Time = time.Unix(timestampUnix, 0)
	return nil
}

//纳秒级的unix timestamp
type UnixNanoStamp struct {
	time.Time
}

//新建unix nano stamp
func NewUnixNanoStamp(t time.Time) UnixNanoStamp {
	return UnixNanoStamp{t}
}

//新建zero值的nano timestamp
func NewUnixNanoZeroStamp() UnixNanoStamp {
	var unixNanoStamp UnixNanoStamp
	unixNanoStamp.Time = _UnixZeroTimeStamp
	return unixNanoStamp
}

//从数据库读取后组装
func (timestampNano *UnixNanoStamp) Scan(src interface{}) error {
	switch src := src.(type) {
	case int64:
		timestampNano.Time = time.Unix(0, src)
		return nil
	default:
		timestampNano.Time = time.Time{}
		return ErrTimestampIsNotInt64
	}
}

//写回数据库时的组装
func (timestampNano *UnixNanoStamp) Value() (driver.Value, error) {
	timestampUnixNano := timestampNano.UnixNano()
	return timestampUnixNano, nil
}

//Json marshal
func (timestampNano UnixNanoStamp) MarshalJSON() ([]byte, error) {
	timestampUnixNano := timestampNano.UnixNano()
	return json.Marshal(timestampUnixNano)
}

//Json marshal
func (timestampNano *UnixNanoStamp) UnmarshalJSON(data []byte) error {
	var timestampUnixNano int64
	err := json.Unmarshal(data, &timestampUnixNano)
	if err != nil {
		return err
	}
	timestampNano.Time = time.Unix(0, timestampUnixNano)
	return nil
}

//将interface转换为UnixStamp
func Convert2UnixStamp(v interface{}) (UnixStamp, bool) {
	var stamp UnixStamp
	ok := true

	switch v.(type) {
	case UnixStamp:
		stamp = v.(UnixStamp)
	case UnixNanoStamp:
		uj := v.(UnixNanoStamp)
		stamp.Time = uj.Time
	case time.Time:
		stamp.Time = v.(time.Time)
	case int64:
		uj := v.(int64)
		stamp.Time = time.Unix(uj, 0)
	case int32:
		uj := v.(int32)
		stamp.Time = time.Unix(int64(uj), 0)
	case int:
		uj := v.(int)
		stamp.Time = time.Unix(int64(uj), 0)
	case uint64:
		uj := v.(uint64)
		stamp.Time = time.Unix(int64(uj), 0)
	case uint32:
		uj := v.(uint32)
		stamp.Time = time.Unix(int64(uj), 0)
	case uint:
		uj := v.(uint)
		stamp.Time = time.Unix(int64(uj), 0)
	case float64:
		uj := v.(float64)
		stamp.Time = time.Unix(int64(uj), 0)
	case float32:
		uj := v.(float32)
		stamp.Time = time.Unix(int64(uj), 0)
	default:
		ok = false
	}
	return stamp, ok

	j, ok := v.(UnixStamp)
	return j, ok
}

//转换map中的key为UnixStamp
func ConvertMapVal2UnixStamp(m map[string]interface{}, k string) (UnixStamp, bool) {
	x, ok := m[k]
	if !ok {
		return UnixStamp{}, false
	}
	return Convert2UnixStamp(x)
}

//将interface转换为UnixNanoStamp
func Convert2UnixNanoStamp(v interface{}) (UnixNanoStamp, bool) {
	var stamp UnixNanoStamp
	ok := true

	switch v.(type) {
	case UnixNanoStamp:
		stamp = v.(UnixNanoStamp)
	case UnixStamp:
		uj := v.(UnixStamp)
		stamp.Time = uj.Time
	case time.Time:
		stamp.Time = v.(time.Time)
	case int64:
		uj := v.(int64)
		stamp.Time = time.Unix(0, uj)
	case int32:
		uj := v.(int32)
		stamp.Time = time.Unix(0, int64(uj))
	case int:
		uj := v.(int)
		stamp.Time = time.Unix(0, int64(uj))
	case uint64:
		uj := v.(uint64)
		stamp.Time = time.Unix(0, int64(uj))
	case uint32:
		uj := v.(uint32)
		stamp.Time = time.Unix(0, int64(uj))
	case uint:
		uj := v.(uint)
		stamp.Time = time.Unix(0, int64(uj))
	case float64:
		uj := v.(float64)
		stamp.Time = time.Unix(0, int64(uj))
	case float32:
		uj := v.(float32)
		stamp.Time = time.Unix(0, int64(uj))
	default:
		ok = false
	}
	return stamp, ok
}

//转换map中的key为UnixStamp
func ConvertMapVal2UnixNanoStamp(m map[string]interface{}, k string) (UnixNanoStamp, bool) {
	x, ok := m[k]
	if !ok {
		return UnixNanoStamp{}, false
	}
	return Convert2UnixNanoStamp(x)
}
