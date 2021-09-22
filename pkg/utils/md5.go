package utils

import (
	"crypto/md5"
	"strconv"
	"unsafe"
)

func NewMD5(key string) string {
	md5Key := md5.Sum([]byte(key))
	md5Int := byteArrayToInt(md5Key)
	return strconv.Itoa(md5Int)
}

func byteArrayToInt(arr [16]byte) int {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	if val < 0 {
		val = val * -1
	}

	return int(val)
}
