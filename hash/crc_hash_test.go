package hash

import (
	"testing"
)

func TestCRCHashStringList(t *testing.T) {
	s1 := "test_hash1"
	s2 := "012Hash2T3"
	s3 := s1 + s2

	u1 := CRCHashStringList(s1, s2)
	u2 := CRCHashString(s3)

	if u1 != u2 {
		t.Fail()
	}

	t.Logf("u1 :%+v u2:%+v", u1, u2)
}
