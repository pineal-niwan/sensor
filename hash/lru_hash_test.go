package hash

import (
	"testing"
)

type _HashT struct {
	v int
}

func judgeHashNotExist(t *testing.T, lruHash *LruHash, key string) {
	_, ok := lruHash.Get(key)
	if ok {
		t.Fail()
	}
}

func judgeHashVal(t *testing.T, lruHash *LruHash, key string, val int) {
	v, ok := lruHash.Get(key)
	if !ok {
		t.Fail()
		return
	}
	if v.(*_HashT).v != val {
		t.Fail()
	}
}

func judgeHashLen(t *testing.T, lruHash *LruHash, length int) {
	ll := lruHash.Len()
	ml := lruHash.HashLen()
	if ll != ml || ll != length {
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
