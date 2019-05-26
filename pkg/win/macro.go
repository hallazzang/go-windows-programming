package win

import "unsafe"

func MAKEINTRESOURCE(i uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(i))
}

func LOWORD(dwValue uint32) uint16 {
	return uint16(dwValue)
}

func HIWORD(dwValue uint32) uint16 {
	return uint16((dwValue >> 16) & 0xffff)
}
