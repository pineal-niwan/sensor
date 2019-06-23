package cypher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// md5 hash
func Md5Bytes(buf []byte) (string, error) {
	h := md5.New()
	return write2Hash(h, buf)
}

// sha1 hash
func Sha1Bytes(buf []byte) (string, error) {
	h := sha1.New()
	return write2Hash(h, buf)
}

// sha256 hash
func Sha256Bytes(buf []byte) (string, error) {
	h := sha256.New()
	return write2Hash(h, buf)
}

// sha512 hash
func Sha512Bytes(buf []byte) (string, error) {
	h := sha512.New()
	return write2Hash(h, buf)
}

// md5 hash
func Md5String(buf string) (string, error) {
	h := md5.New()
	return write2Hash(h, []byte(buf))
}

// sha1 hash
func Sha1String(buf string) (string, error) {
	h := sha1.New()
	return write2Hash(h, []byte(buf))
}

// sha256 hash
func Sha256String(buf string) (string, error) {
	h := sha256.New()
	return write2Hash(h, []byte(buf))
}

// sha512 hash
func Sha512String(buf string) (string, error) {
	h := sha512.New()
	return write2Hash(h, []byte(buf))
}
