package lru

import (
	"testing"
	"time"
)

type _HashT struct {
	v int
}

func judgeHashNotExist(t *testing.T, lruHash *LruHash, key string) {
	_, ok := lruHash.Get(key)
	if ok {
		t.Error(`judgeHashNotExist`)
		t.Fail()
	}
}

func judgeHashVal(t *testing.T, lruHash *LruHash, key string, val int) {
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

func judgeHashValWithRefresh(t *testing.T, lruHash *LruHash, key string, val int, timeout int64) {
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

func judgeHashLen(t *testing.T, lruHash *LruHash, length int) {
	ll, ml := lruHash.Len()
	if ll != ml || ll != length {
		t.Error(`judgeHashLen`)
		t.Fail()
	}
}

func TestLruHash_Get_Set1(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.Set("1", &_HashT{1})
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("2", &_HashT{2})
	judgeHashLen(t, lruHash, 2)

	lruHash.Set("3", &_HashT{3})
	judgeHashLen(t, lruHash, 3)

	lruHash.Set("4", &_HashT{4})
	judgeHashLen(t, lruHash, 3)

	judgeHashNotExist(t, lruHash, "1")
	judgeHashVal(t, lruHash, "2", 2)
	judgeHashVal(t, lruHash, "3", 3)
	judgeHashVal(t, lruHash, "4", 4)

	judgeHashLen(t, lruHash, 3)
}

func TestLruHash_Get_Set2(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.Set("1", &_HashT{1})
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("2", &_HashT{2})
	judgeHashLen(t, lruHash, 2)

	lruHash.Set("3", &_HashT{3})
	judgeHashLen(t, lruHash, 3)

	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("4", &_HashT{4})
	judgeHashLen(t, lruHash, 3)

	judgeHashNotExist(t, lruHash, "2")
	judgeHashVal(t, lruHash, "1", 1)
	judgeHashVal(t, lruHash, "3", 3)
	judgeHashVal(t, lruHash, "4", 4)

	judgeHashLen(t, lruHash, 3)
}

func TestLruHash_Get_Set3(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.Set("1", &_HashT{1})
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("2", &_HashT{2})
	judgeHashLen(t, lruHash, 2)

	lruHash.Set("3", &_HashT{3})
	judgeHashLen(t, lruHash, 3)

	lruHash.Set("1", &_HashT{11})
	lruHash.Set("4", &_HashT{4})
	judgeHashLen(t, lruHash, 3)

	judgeHashNotExist(t, lruHash, "2")
	judgeHashVal(t, lruHash, "1", 11)
	judgeHashVal(t, lruHash, "3", 3)
	judgeHashVal(t, lruHash, "4", 4)

	judgeHashLen(t, lruHash, 3)
}

func TestLruHash_Get_Set4(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.Set("1", &_HashT{1})
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("2", &_HashT{2})
	judgeHashLen(t, lruHash, 2)

	time.Sleep(time.Second * 10)
	lruHash.SetWithTimeout("3", &_HashT{3}, 1000)
	judgeHashLen(t, lruHash, 3)

	time.Sleep(time.Second * 10)
	lruHash.Set("1", &_HashT{11})
	lruHash.Set("4", &_HashT{4})
	judgeHashLen(t, lruHash, 3)

	judgeHashNotExist(t, lruHash, "2")
	judgeHashVal(t, lruHash, "1", 11)
	time.Sleep(time.Second * 10)
	judgeHashNotExist(t, lruHash, "3")
	judgeHashVal(t, lruHash, "4", 4)

	judgeHashLen(t, lruHash, 2)
}

func TestLruHash_Get_Set5(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.SetWithTimeout("1", &_HashT{1}, 1)
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.SetWithTimeout("2", &_HashT{2}, 1)
	judgeHashLen(t, lruHash, 2)
	judgeHashValWithRefresh(t, lruHash, "2", 2, 3)

	t.Log(`before sleep`)
	time.Sleep(2 * time.Millisecond)
	t.Log(`after sleep`)
	judgeHashNotExist(t, lruHash, "1")
	judgeHashVal(t, lruHash, "2", 2)

	judgeHashLen(t, lruHash, 1)
}

func TestLruHash_Get_Set6(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.SetWithTimeout("1", &_HashT{1}, 1)
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.SetWithTimeout("2", &_HashT{2}, 1)
	judgeHashLen(t, lruHash, 2)
	judgeHashValWithRefresh(t, lruHash, "2", 2, 3)

	t.Log(`before sleep`)
	time.Sleep(5 * time.Millisecond)
	t.Log(`after sleep`)
	judgeHashNotExist(t, lruHash, "1")
	judgeHashNotExist(t, lruHash, "2")

	judgeHashLen(t, lruHash, 0)
}

func TestLruHash_Get_Set_Remove(t *testing.T) {
	lruHash := NewLruHash(3)

	lruHash.Set("1", &_HashT{1})
	judgeHashLen(t, lruHash, 1)
	judgeHashVal(t, lruHash, "1", 1)

	lruHash.Set("2", &_HashT{2})
	judgeHashLen(t, lruHash, 2)

	lruHash.Set("3", &_HashT{3})
	judgeHashLen(t, lruHash, 3)

	lruHash.Set("1", &_HashT{11})
	lruHash.Remove("1")
	lruHash.Set("4", &_HashT{4})
	judgeHashLen(t, lruHash, 3)

	judgeHashNotExist(t, lruHash, "1")
	judgeHashVal(t, lruHash, "2", 2)
	judgeHashVal(t, lruHash, "3", 3)
	judgeHashVal(t, lruHash, "4", 4)

	judgeHashLen(t, lruHash, 3)
}
