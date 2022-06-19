package util

import "unsafe"

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	p := (*[]uintptr)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&p))
}
