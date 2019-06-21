package id_gen

import (
	"sync"
	"time"
)

var (
	//为节点分配留16位
	_SpecNodeBits uint8 = 16

	//同一个毫秒冲突后的步长，这里只剩6位，相当于1毫秒可以产生64个步长, 1秒产生64000个步长
	_SpecStepBits uint8 = 6

	//step的掩码 -- step不能超过63
	_SpecStepMask int64 = -1 ^ (-1 << _SpecStepBits)

	//时间戳的位移，22位
	_SpecTimeShift = _SpecNodeBits + _SpecStepBits
	//step的位移，16位
	_SpecStepShift = _SpecNodeBits

	//hash基数 - 65536 = 64k -- 最多可以将数据分配到65536个节点上
	_SpecHashBase int64 = 1 << _SpecNodeBits
)

type SnowFlakeSpecNode struct {
	mu   sync.Mutex
	time int64
	step int64
}

// 生成ID -- 此hash与传入的int64值在64k基数上的hash是一致的
func (n *SnowFlakeSpecNode) Generate(outId int64) int64 {

	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1000000

	if n.time == now {
		n.step = (n.step + 1) & _SpecStepMask

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
	r := ((now - _Epoch) << _SpecTimeShift) | (n.step << _SpecStepShift) | (outId % _SpecHashBase)

	return r
}
