package time_tool

import "testing"

func TestParseDate(t *testing.T) {
	day, err := ParseDate(`2239-12-05`)
	if err != nil{
		t.Errorf("解析错误:%+v", err)
	}else{
		t.Log(day)
	}
}

func TestParseTime(t *testing.T) {
	day, err := ParseTime(`2239-12-05   15:25:56`)
	if err != nil{
		t.Errorf("解析错误:%+v", err)
	}else{
		t.Log(day)
	}
}