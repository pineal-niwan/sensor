package id_gen

import (
	"testing"
	"time"
)

func TestShow16Bit(t *testing.T) {
	t.Log(0xFFFF)

	t.Log("_StepMask", _StepMask)
	t.Log("_TimeShift", _TimeShift)
}

func TestNode_Generate(t *testing.T) {
	node := &SnowFlakeNode{}
	count := _SpecHashBase * 16
	x := make(map[int64]struct{})

	t1 := time.Now()
	for i := int64(0); i < count; i++ {
		id := node.Generate(uint8(i%256))
		_, ok := x[id]
		if ok {
			t.Fail()
			return
		}
		x[id] = struct{}{}
	}
	t2 := time.Now()

	t.Logf("each :%+v", t2.Sub(t1)/time.Duration(count))
}
