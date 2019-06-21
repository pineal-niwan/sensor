package id_gen

import (
	"sync"
	"time"
)

var (
	// 以北京时间2019-01-01 00:00:00 开始，比twitter的2010年多了近9年
	_Epoch int64 = 1546275600000

	//为节点分配8位
	_NodeBits uint8 = 8

	//为步长分配12位
	_StepBits uint8 = 12

	//时间戳的位移, 20位
	_TimeShift = _NodeBits + _StepBits

	//步长的掩码
	_StepMask int64 = -1 ^ (-1 << _StepBits)
)

type SnowFlakeNode struct {
	mu   sync.Mutex
	time int64
	step int64
}

// 生成ID -- 传入的是nodeId(0-255)
func (n *SnowFlakeNode) Generate(nodeId uint8) int64 {

	n.mu.Lock()

	now := time.Now().UnixNano() / 1000000

	if n.time == now {
		n.step = (n.step + 1) & _StepMask

		if n.step == 0 {
			for now <= n.time {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	//前面的位数是时间戳，中间为node，最后为step值
	r := (now-_Epoch)<<_TimeShift | (int64(nodeId) << _StepBits) | n.step

	n.mu.Unlock()
	return r
}
