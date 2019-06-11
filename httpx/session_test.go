package httpx

import (
	"pineal/sensor/cypher"
	"testing"
	"time"
)

func TestSessionKeyEncodeDecode(t *testing.T) {
	uuid := cypher.RawUuid()
	t.Logf("uuid - %+v", uuid)
	encode, err := encodeSession("tData", HashKeyWithTime{`12345678`, time.Now()})
	t.Logf("encode - %+v", encode)
	if err != nil {
		t.Error(err)
	} else {
		hashKeyWithTime, err := decodeSession("tData", encode)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(hashKeyWithTime)
		}
	}
}
