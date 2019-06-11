package cypher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/go-errors/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
	"hash"
)

var (
	ErrWriteHashStop = errors.New("write to hash buffer stopped")
)

//原始的V1版本uuid
func RawUuid() string {
	id := uuid.NewV1()
	return hex.EncodeToString(id[:])
}

//md5版本的uuid
func Md5Uuid() (string, error) {
	id := uuid.NewV1()
	h := md5.New()
	return write2Hash(h, id[:])
}

//sha1版本的uuid
func Sha1Uuid() (string, error) {
	id := uuid.NewV1()
	h := sha1.New()
	return write2Hash(h, id[:])
}

//sha2 256版本的uuid
func Sha2Uuid256() (string, error) {
	id := uuid.NewV1()
	h := sha256.New()
	return write2Hash(h, id[:])
}

//sha2 512版本的uuid
func Sha2Uuid512() (string, error) {
	id := uuid.NewV1()
	h := sha512.New()
	return write2Hash(h, id[:])
}

//sha3 256版本的uuid
func Sha3Uuid256() (string, error) {
	id := uuid.NewV1()
	h := sha3.New256()
	return write2Hash(h, id[:])
}

//sha3 512版本的uuid
func Sha3Uuid512() (string, error) {
	id := uuid.NewV1()
	h := sha3.New512()
	return write2Hash(h, id[:])
}

//生成相关的hash
func write2Hash(h hash.Hash, buf []byte) (string, error) {
	n, err := h.Write(buf)
	if n != len(buf) {
		err = ErrWriteHashStop
	}
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
