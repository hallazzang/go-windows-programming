package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	libuser32   = windows.NewLazySystemDLL("user32.dll")
	libwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")

	procCreateWindowExW        = libuser32.NewProc("CreateWindowExW")
	procDispatchMessageW       = libuser32.NewProc("DispatchMessageW")
	procGetMessageW            = libuser32.NewProc("GetMessageW")
	procGetModuleHandleW       = libkernel32.NewProc("GetModuleHandleW")
	procRegisterClassExW       = libuser32.NewProc("RegisterClassExW")
	procShowWindow             = libuser32.NewProc("ShowWindow")
	procTranslateMessage       = libuser32.NewProc("TranslateMessage")
	procUpdateWindow           = libuser32.NewProc("UpdateWindow")
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

func CreateWindowEx(
	dwExStyle uint32,
	lpClassName, lpWindowName *uint16,
	dwStyle uint32,
	X, Y, nWidth, nHeight int32,
	hWndParent, hMenu, hInstance uintptr,
	lpParam unsafe.Pointer) uintptr {
	var r1 uintptr
	r1, _, lastError = procCreateWindowExW.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(X),
		uintptr(Y),
		uintptr(nWidth),
		uintptr(nHeight),
		hWndParent,
		hMenu,
		hInstance,
		uintptr(lpParam))
	return r1
}

func DispatchMessage(lpMsg *MSG) uintptr {
	var r1 uintptr
	r1, _, _ = procDispatchMessageW.Call(uintptr(unsafe.Pointer(lpMsg)))
	return r1
}

func GetMessage(lpMsg *MSG, hWnd uintptr, uMsgFilterMin, uMsgFilterMax uint32) int32 {
	var r1 uintptr
	r1, _, lastError = procGetMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(uMsgFilterMin),
		uintptr(uMsgFilterMax))
	return int32(r1)
}

func GetModuleHandle(lpModuleName *uint16) uintptr {
	var r1 uintptr
	r1, _, lastError = procGetModuleHandleW.Call(uintptr(unsafe.Pointer(lpModuleName)))
	return r1
}

func RegisterClassEx(Arg1 *WNDCLASSEX) uint16 {
	var r1 uintptr
	r1, _, lastError = procRegisterClassExW.Call(uintptr(unsafe.Pointer(Arg1)))
	return uint16(r1)
}

func ShowWindow(hWnd uintptr, nCmdShow int32) int32 {
	var r1 uintptr
	r1, _, _ = procShowWindow.Call(hWnd, uintptr(nCmdShow))
	return int32(r1)
}

func TranslateMessage(lpMsg *MSG) int32 {
	var r1 uintptr
	r1, _, _ = procTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return int32(r1)
}

func UpdateWindow(hWnd uintptr) int32 {
	var r1 uintptr
	r1, _, lastError = procUpdateWindow.Call(hWnd)
	return int32(r1)
}

func WTSEnumerateProcesses(
	hServer uintptr,
	Reserved, Version uint32,
	ppProcessInfo **WTS_PROCESS_INFO,
	pCount *uint32) int32 {
	var r1 uintptr
	r1, _, lastError = procWTSEnumerateProcessesW.Call(
		hServer,
		uintptr(Reserved),
		uintptr(Version),
		uintptr(unsafe.Pointer(ppProcessInfo)),
		uintptr(unsafe.Pointer(pCount)))
	return int32(r1)
}

func WTSFreeMemory(pMemory unsafe.Pointer) {
	procWTSFreeMemory.Call(uintptr(pMemory))
}
