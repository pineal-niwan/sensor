package sql_convert

import (
	"encoding/json"
	"github.com/pineal-niwan/sensor/time_tool"
	"testing"
	"time"
)

func TestUnixStamp_Scan(t *testing.T) {
	var unixStamp UnixStamp

	longTime, err := time_tool.ParseTime(`2788-11-16 23:59:59`)
	if err != nil {
		t.Errorf("解析错误%+v", err)
	} else {
		err = unixStamp.Scan(longTime.Unix())
		if err != nil {
			t.Errorf("scan失败%+v", err)
		} else {
			t.Logf("unix stamp：%+v", unixStamp)
		}
	}
}

func TestUnixStamp_Value(t *testing.T) {
	var unixStamp UnixStamp

	longTime, err := time_tool.ParseTime(`2788-11-16 23:59:59`)
	longTimeUnix := longTime.Unix()
	if err != nil {
		t.Errorf("解析错误%+v", err)
	} else {
		unixStamp.Time = longTime
		finalUnix, err := unixStamp.Value()
		if err != nil {
			t.Errorf("unix stamp.value error:%+v", err)
		} else {
			t.Log(longTimeUnix)
			t.Log(finalUnix)
		}
	}
}

func TestUnixNanoStamp_Scan(t *testing.T) {
	var unixNanoStamp UnixNanoStamp

	longTime, err := time_tool.ParseTime(`2188-11-16 23:59:59`)
	if err != nil {
		t.Errorf("解析错误%+v", err)
	} else {
		err = unixNanoStamp.Scan(longTime.UnixNano())
		if err != nil {
			t.Errorf("scan失败%+v", err)
		} else {
			t.Logf("unix stamp：%+v", unixNanoStamp)
		}
	}
}

func TestUnixNanoStamp_Value(t *testing.T) {
	var unixNanoStamp UnixNanoStamp

	timeNow := time.Now()
	unixNanoStamp.Time = timeNow

	finalUnix, err := unixNanoStamp.Value()
	if err != nil {
		t.Errorf("unix stamp.value error:%+v", err)
	} else {
		t.Log(timeNow.UnixNano())
		t.Log(finalUnix)
	}
}

func TestMarshalUnixStamp(t *testing.T) {
	t1 := time.Now()
	t2 := UnixStamp{Time:t1}
	x1, err1 := json.Marshal(t1)
	x2, err2 := json.Marshal(t2)

	t.Logf("time is %+v", t1)

	t.Logf("x1 is :%+v err1 is :%+v", x1, err1)
	t.Logf("x2 is :%+v err2 is :%+v", x2, err2)

	var y1 time.Time
	var y2 UnixStamp

	err1 = json.Unmarshal(x1, &y1)
	err2 = json.Unmarshal(x2, &y2)
	t.Logf("y1 is :%+v err1 is :%+v", y1, err1)
	t.Logf("y2 is :%+v err2 is :%+v", y2, err2)
}