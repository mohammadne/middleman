package utils

import (
	"crypto/md5"
	"encoding/binary"
)

func NewMd5(key string) uint64 {
	checksum := md5.New().Sum([]byte(key))
	return binary.BigEndian.Uint64(checksum)
}
