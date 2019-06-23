package lru

import (
	"testing"
	"time"
)

func judgeHashNotExistK64(t *testing.T, lruHash *LruInt64AsKeyHash, key int64) {
	_, ok := lruHash.Get(key)
	if ok {
		t.Error(`judgeHashNotExist`)
		t.Fail()
	}
}

func judgeHashValK64(t *testing.T, lruHash *LruInt64AsKeyHash, key int64, val int) {
	v, ok := lruHash.Get(key)
	if !ok {
		t.Errorf(`judgeHashVal not ok`)
		t.Fail()
		return
	}
	if v.(*_HashT).v != val {
		t.Errorf(`judgeHashVal not equal`)
		t.Fail()
	}
}

func judgeHashValWithRefreshK64(t *testing.T, lruHash *LruInt64AsKeyHash, key int64, val int, timeout int64) {
	v, ok := lruHash.GetThenRefreshTimeout(key, timeout)
	if !ok {
		t.Error(`judgeHashValWithRefresh not ok`)
		t.Fail()
		return
	}
	if v.(*_HashT).v != val {
		t.Error(`judgeHashValWithRefresh not equal`)
		t.Fail()
	}
}

func judgeHashLenK64(t *testing.T, lruHash *LruInt64AsKeyHash, length int) {
	ll, ml := lruHash.Len()
	if ll != ml || ll != length {
		t.Error(`judgeHashLen`)
		t.Fail()
	}
}

func TestLruHash_Get_Set1K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.Set(1, &_HashT{1})
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.Set(2, &_HashT{2})
	judgeHashLenK64(t, lruHash, 2)

	lruHash.Set(3, &_HashT{3})
	judgeHashLenK64(t, lruHash, 3)

	lruHash.Set(4, &_HashT{4})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashNotExistK64(t, lruHash, 2)
	judgeHashValK64(t, lruHash, 3, 2)
	judgeHashValK64(t, lruHash, 3, 3)
	judgeHashValK64(t, lruHash, 4, 4)

	judgeHashLenK64(t, lruHash, 3)
}

func TestLruHash_Get_Set2K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.Set(2, &_HashT{1})
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 2, 1)

	lruHash.Set(2, &_HashT{2})
	judgeHashLenK64(t, lruHash, 2)

	lruHash.Set(3, &_HashT{3})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.Set(4, &_HashT{4})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashNotExistK64(t, lruHash, 2)
	judgeHashValK64(t, lruHash, 1, 1)
	judgeHashValK64(t, lruHash, 3, 3)
	judgeHashValK64(t, lruHash, 4, 4)

	judgeHashLenK64(t, lruHash, 3)
}

func TestLruHash_Get_Set3K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.Set(1, &_HashT{1})
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.Set(2, &_HashT{2})
	judgeHashLenK64(t, lruHash, 2)

	lruHash.Set(3, &_HashT{3})
	judgeHashLenK64(t, lruHash, 3)

	lruHash.Set(1, &_HashT{11})
	lruHash.Set(4, &_HashT{4})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashNotExistK64(t, lruHash, 2)
	judgeHashValK64(t, lruHash, 1, 11)
	judgeHashValK64(t, lruHash, 3, 3)
	judgeHashValK64(t, lruHash, 4, 4)

	judgeHashLenK64(t, lruHash, 3)
}

func TestLruHash_Get_Set4K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.Set(1, &_HashT{1})
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.Set(2, &_HashT{2})
	judgeHashLenK64(t, lruHash, 2)

	time.Sleep(time.Second * 10)
	lruHash.SetWithTimeout(3, &_HashT{3}, 1000)
	judgeHashLenK64(t, lruHash, 3)

	time.Sleep(time.Second * 10)
	lruHash.Set(1, &_HashT{11})
	lruHash.Set(4, &_HashT{4})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashNotExistK64(t, lruHash, 2)
	judgeHashValK64(t, lruHash, 1, 11)
	time.Sleep(time.Second * 10)
	judgeHashNotExistK64(t, lruHash, 3)
	judgeHashValK64(t, lruHash, 4, 4)

	judgeHashLenK64(t, lruHash, 2)
}

func TestLruHash_Get_Set5K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.SetWithTimeout(1, &_HashT{1}, 1)
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.SetWithTimeout(2, &_HashT{2}, 1)
	judgeHashLenK64(t, lruHash, 2)
	judgeHashValWithRefreshK64(t, lruHash, 2, 2, 3)

	t.Log(`before sleep`)
	time.Sleep(2 * time.Millisecond)
	t.Log(`after sleep`)
	judgeHashNotExistK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 2, 2)

	judgeHashLenK64(t, lruHash, 1)
}

func TestLruHash_Get_Set6K64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.SetWithTimeout(1, &_HashT{1}, 1)
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.SetWithTimeout(2, &_HashT{2}, 1)
	judgeHashLenK64(t, lruHash, 2)
	judgeHashValWithRefreshK64(t, lruHash, 2, 2, 3)

	t.Log(`before sleep`)
	time.Sleep(5 * time.Millisecond)
	t.Log(`after sleep`)
	judgeHashNotExistK64(t, lruHash, 1)
	judgeHashNotExistK64(t, lruHash, 2)

	judgeHashLenK64(t, lruHash, 0)
}

func TestLruHash_Get_Set_RemoveK64(t *testing.T) {
	lruHash := NewLruInt64AsKeyHash(3)

	lruHash.Set(1, &_HashT{1})
	judgeHashLenK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 1, 1)

	lruHash.Set(2, &_HashT{2})
	judgeHashLenK64(t, lruHash, 2)

	lruHash.Set(3, &_HashT{3})
	judgeHashLenK64(t, lruHash, 3)

	lruHash.Set(1, &_HashT{11})
	lruHash.Remove(1)
	lruHash.Set(4, &_HashT{4})
	judgeHashLenK64(t, lruHash, 3)

	judgeHashNotExistK64(t, lruHash, 1)
	judgeHashValK64(t, lruHash, 2, 2)
	judgeHashValK64(t, lruHash, 3, 3)
	judgeHashValK64(t, lruHash, 4, 4)

	judgeHashLenK64(t, lruHash, 3)
}
