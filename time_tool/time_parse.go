package time_tool

import "time"

//解析日期 - 格式为`YYYY-MM-DD`
func ParseDate(timeStr string) (time.Time, error) {
	return time.ParseInLocation(`2006-01-02`, timeStr, time.Local)
}

//解析时间 - 格式为`YYYY-MM-DD HH:MM:SS`
func ParseTime(timeStr string) (time.Time, error) {
	return time.ParseInLocation(`2006-01-02 15:04:05`, timeStr, time.Local)
}
