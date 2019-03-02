package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libuser32 *windows.LazyDLL

	procMessageBoxW *windows.LazyProc
)

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")

	procMessageBoxW = libuser32.NewProc("MessageBoxW")
}

func messageBox(hWnd uintptr, lpText *uint16, lpCaption *uint16, uType uint32) int32 {
	r1, _, _ := procMessageBoxW.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpText)), uintptr(unsafe.Pointer(lpCaption)), uintptr(uType))
	return int32(r1)
}
