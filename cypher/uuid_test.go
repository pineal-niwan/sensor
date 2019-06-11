package cypher

import "testing"

func TestRawUuid(t *testing.T) {
	u1 := RawUuid()
	u2 := RawUuid()
	t.Logf("%+v - %+v", u1, u2)
}

func TestMd5Uuid(t *testing.T) {
	u1, err := Md5Uuid()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Md5Uuid()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}

func TestSha1Uuid(t *testing.T) {
	u1, err := Sha1Uuid()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Sha1Uuid()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}

func TestSha2Uuid256(t *testing.T) {
	u1, err := Sha2Uuid256()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Sha2Uuid256()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}

func TestSha2Uuid512(t *testing.T) {
	u1, err := Sha2Uuid512()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Sha2Uuid512()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}

func TestSha3Uuid256(t *testing.T) {
	u1, err := Sha2Uuid256()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Sha2Uuid256()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}

func TestSha3Uuid512(t *testing.T) {
	u1, err := Sha2Uuid512()
	if err != nil {
		t.Errorf("%+v", err)
	}
	u2, err := Sha2Uuid512()
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v - %+v", u1, u2)
}
