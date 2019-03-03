package win

import (
	"unicode/utf16"
	"unsafe"
)

func UTF16PtrToString(p *uint16) string {
	var s []uint16
	for i := 0; i < 65536; i++ {
		c := *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(unsafe.Sizeof(uint16(0))*uintptr(i))))
		if c == 0 {
			break
		}
		s = append(s, c)
	}
	return string(utf16.Decode(s))
}
