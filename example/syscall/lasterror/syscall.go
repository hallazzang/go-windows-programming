package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libuser32 = windows.NewLazySystemDLL("user32.dll")

	procMessageBoxW = libuser32.NewProc("MessageBoxW")

	lastError error
)

func messageBox(hWnd uintptr, lpText *uint16, lpCaption *uint16, uType uint32) int32 {
	var r1 uintptr
	r1, _, lastError = procMessageBoxW.Call(hWnd, uintptr(unsafe.Pointer(lpText)), uintptr(unsafe.Pointer(lpCaption)), uintptr(uType))
	return int32(r1)
}

func getLastError() error {
	if lastError.(windows.Errno) != 0 {
		return lastError
	}
	return nil
}
