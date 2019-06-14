package sql_convert

import (
	"database/sql/driver"
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

//纳秒级的unix timestamp
type UnixNanoStamp struct {
	time.Time
}

//从数据库读取后组装
func (timestamp *UnixNanoStamp) Scan(src interface{}) error {
	switch src := src.(type) {
	case int64:
		timestamp.Time = time.Unix(0, src)
		return nil
	default:
		timestamp.Time = time.Time{}
		return ErrTimestampIsNotInt64
	}
}

//写回数据库时的组装
func (timestamp *UnixNanoStamp) Value() (driver.Value, error) {
	timestampUnix := timestamp.UnixNano()
	return timestampUnix, nil
}
