package sql_convert

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/go-errors/errors"
	"time"
)

var (
	ErrTimestampIsNotInt64 = errors.New("time stamp is not int64")
)

//秒级的unix timestamp
type UnixStamp struct {
	time.Time
}

//新建unix stamp
func NewUnixStamp(t time.Time) UnixStamp {
	return UnixStamp{t}
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

//将interface转换为UnixStamp
func Convert2UnixStamp(v interface{}) (UnixStamp, bool) {
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

//纳秒级的unix timestamp
type UnixNanoStamp struct {
	time.Time
}

//新建unix nano stamp
func NewUnixNanoStamp(t time.Time) UnixNanoStamp {
	return UnixNanoStamp{t}
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

//将interface转换为UnixNanoStamp
func Convert2UnixNanoStamp(v interface{}) (UnixNanoStamp, bool) {
	j, ok := v.(UnixNanoStamp)
	return j, ok
}

//转换map中的key为UnixStamp
func ConvertMapVal2UnixNanoStamp(m map[string]interface{}, k string) (UnixNanoStamp, bool) {
	x, ok := m[k]
	if !ok {
		return UnixNanoStamp{}, false
	}
	return Convert2UnixNanoStamp(x)
}
