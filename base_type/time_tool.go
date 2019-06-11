package base_type

import (
	"time"
)

// 根据日期字符串获取本地日期
func GetLocalTimeByDateStr(dateStr string) (t time.Time, err error) {
	return time.ParseInLocation("2006-01-02", dateStr, time.Local)
}

// 根据时间字符串获取当地时间
func GetLocalTimeByTimeStr(timeStr string) (t time.Time, err error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
}
