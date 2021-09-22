package utils

import (
	"crypto/md5"
	"unsafe"
)

func NewMd5(key string) [16]byte {
	return md5.Sum([]byte(key))
}

// func NewMD5String(key string) string {
// 	md5 := NewMD5(key)
// 	return string(md5[:])
// }

// func NewMD5Int(key string) int {
// 	md5 := NewMD5(key)
// 	return byteArrayToInt(md5)
// }

func Md5HashToInt(arr [16]byte) int {
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
