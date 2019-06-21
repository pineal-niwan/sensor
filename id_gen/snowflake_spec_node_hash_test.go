package id_gen

import (
	"testing"
	"time"
)

func TestSpecShow16Bit(t *testing.T) {
	t.Log(0xFFFF)

	t.Log("_SpecStepMask", _SpecStepMask)
	t.Log("_SpecTimeShift", _SpecTimeShift)
	t.Log("_SpecStepShift", _SpecStepShift)
	t.Log("_SpecHashBase", _SpecHashBase)
}

func TestSpecStepMask(t *testing.T) {
	count := int64(64)
	showX := make([]int64, count)

	start := int64(0)
	for i := start; i < start+count; i++ {
		showX[i] = i & _SpecStepMask
	}
	t.Log(showX)

	start = int64(64)
	for i := start; i < start+count; i++ {
		showX[i&_SpecStepMask] = i & _SpecStepMask
	}
	t.Log(showX)
}

func TestSpecNode_Generate(t *testing.T) {
	node := &SnowFlakeSpecNode{}
	count := _SpecHashBase * 16

	t1 := time.Now()
	for i := int64(0); i < count; i++ {
		id := node.Generate(i)
		if i%_SpecHashBase != id%_SpecHashBase {
			t.Errorf("fail: i:%+v id:%+v", i, id)
		}
	}
	t2 := time.Now()

	t.Logf("each :%+v", t2.Sub(t1)/time.Duration(count))
}
