package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")

	procWTSEnumerateProcessesW = libwtsapi32.NewProc("WTSEnumerateProcessesW")
	procWTSFreeMemory          = libwtsapi32.NewProc("WTSFreeMemory")

	lastError error
)

func GetLastError() error {
	if lastError.(windows.Errno) != 0 {
		return lastError
	}
	return nil
}

func WTSEnumerateProcesses(hServer uintptr, Reserved, Version uint32, ppProcessInfo **WTS_PROCESS_INFO, pCount *uint32) int32 {
	var r1 uintptr
	r1, _, lastError = procWTSEnumerateProcessesW.Call(hServer, uintptr(Reserved), uintptr(Version), uintptr(unsafe.Pointer(ppProcessInfo)), uintptr(unsafe.Pointer(pCount)))
	return int32(r1)
}

func WTSFreeMemory(pMemory unsafe.Pointer) {
	procWTSFreeMemory.Call(uintptr(pMemory))
}
