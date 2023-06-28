package main

import (
	"unsafe"

	"github.com/google/uuid"
)

// ByteArrayToString 字节数组转string
func ByteArrayToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// StringToByteArray 字符串转字节数组
func StringToByteArray(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// GeneralUUID 生成uuid字符串
func GeneralUUID() string {
	return uuid.New().String()
}
