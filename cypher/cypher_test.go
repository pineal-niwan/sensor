package cypher

import (
	"testing"
)

func testDecryptDecrypt(t *testing.T, key []byte, message string) {
	x, err := Encrypt(key, message)
	if err != nil {
		t.Errorf("encrypt error:%+v", err)
	}else{
		t.Logf("encrypt message:%+v", x)
	}

	y, err := Decrypt(key, x)
	if err != nil {
		t.Errorf("encrypt error:%+v", err)
	}else {
		t.Logf("decrypt message:%+v", y)
		if y != message {
			t.Fail()
		}
	}
}

func TestEncryptDecrypt(t *testing.T) {
	testDecryptDecrypt(t, []byte("coqz24HKzk21k0091NAMZLAB"), "a encrypt message test.")
	testDecryptDecrypt(t, []byte("coqz24HKzk21k0091NAMZLAB"), "hello world")
	testDecryptDecrypt(t, []byte("coqz24HKzk21k0091NAMZLAC"), "hello world")

	b16 := make([]byte, 16)
	testDecryptDecrypt(t, b16, "!@#$%^&*()")
	b24 := make([]byte, 24)
	testDecryptDecrypt(t, b24, "!@#$%^&*()")
	b32 := make([]byte, 32)
	testDecryptDecrypt(t, b32, "!@#$%^&*()")
}