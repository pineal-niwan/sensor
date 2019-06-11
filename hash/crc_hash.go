package hash

import (
	"hash/crc32"
)

//对字符串做crc hash
func CRCHashString(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//对字符串list做crc hash
func CRCHashStringList(keyList ...string) uint32 {
	length := 0
	for _, key := range keyList {
		length += len(key)
	}
	buf := make([]byte, length)
	index := 0
	for _, key := range keyList {
		index = copy(buf[index:], key)
	}
	return crc32.ChecksumIEEE(buf)
}
