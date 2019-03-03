package win

import "unsafe"

func MAKEINTRESOURCE(i uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(i))
}
