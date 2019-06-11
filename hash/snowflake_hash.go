package hash

import (
	"sync"
	"time"
)

var (
	// 以北京时间2019-01-01 00:00:00 开始，比twitter的2010年多了近9年
	_Epoch int64 = 1546275600000

	//为hash分配留16位
	_NodeBits uint8 = 16

	//同一个毫秒冲突后的步长，这里只剩6位，相当于1毫秒可以产生64个步长, 1秒产生64000个步长
	_StepBits uint8 = 6

	//step的掩码 -- step不能超过63
	_StepMask int64 = -1 ^ (-1 << _StepBits)

	//时间戳的位移，22位
	_TimeShift uint8 = _NodeBits + _StepBits
	//step的位移，16位
	_StepShift uint8 = _NodeBits

	//hash基数 - 65536 = 64k -- 最多可以将数据分配到65536个节点上
	_HashBase int64 = 1 << _NodeBits
)

type Node struct {
	mu   sync.Mutex
	time int64
	step int64
}

// 生成ID -- 此hash与传入的int64值在64k基数上的hash是一致的
func (n *Node) Generate(outId int64) int64 {

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

	//前面的位数是时间戳，中间为step，最后为外部传入数据以64k为基础计算的hash值
	r := (now-_Epoch)<<_TimeShift | (n.step << _StepShift) | (outId % _HashBase)

	n.mu.Unlock()
	return r
}
