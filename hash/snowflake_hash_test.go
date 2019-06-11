package hash

import (
	"testing"
	"time"
)

func TestShow16Bit(t *testing.T) {
	t.Log(0xFFFF)

	t.Log("_StepMask", _StepMask)
	t.Log("_TimeShift", _TimeShift)
	t.Log("_StepShift", _StepShift)
	t.Log("_HashBase", _HashBase)
}

func TestStepMask(t *testing.T) {
	count := int64(64)
	showX := make([]int64, count)

	start := int64(0)
	for i := start; i < start+count; i++ {
		showX[i] = i & _StepMask
	}
	t.Log(showX)

	start = int64(64)
	for i := start; i < start+count; i++ {
		showX[i&_StepMask] = i & _StepMask
	}
	t.Log(showX)
}

func TestNode_Generate(t *testing.T) {
	node := &Node{}
	count := _HashBase * 16

	t1 := time.Now()
	for i := int64(0); i < count; i++ {
		id := node.Generate(i)
		if i%_HashBase != id%_HashBase {
			t.Errorf("fail: i:%+v id:%+v", i, id)
		}
	}
	t2 := time.Now()

	t.Logf("each :%+v", t2.Sub(t1)/time.Duration(count))
}
