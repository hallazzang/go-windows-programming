package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libuser32   *windows.LazyDLL
	libkernel32 *windows.LazyDLL

	procMessageBoxW  *windows.LazyProc
	procGetLastError *windows.LazyProc

	lastError error
)

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procMessageBoxW = libuser32.NewProc("MessageBoxW")
	procGetLastError = libkernel32.NewProc("GetLastError")
}

func messageBox(hWnd uintptr, lpText *uint16, lpCaption *uint16, uType uint32) int32 {
	var r1 uintptr
	r1, _, lastError = procMessageBoxW.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpText)), uintptr(unsafe.Pointer(lpCaption)), uintptr(uType))
	return int32(r1)
}

func getLastError() error {
	if lastError.(windows.Errno) != 0 {
		return lastError
	}
	return nil
}
