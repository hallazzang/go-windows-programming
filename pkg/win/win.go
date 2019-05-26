package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	libshell32  = windows.NewLazySystemDLL("shell32.dll")
	libuser32   = windows.NewLazySystemDLL("user32.dll")
	libwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")

	procCreateWindowExW        = libuser32.NewProc("CreateWindowExW")
	procDefWindowProcW         = libuser32.NewProc("DefWindowProcW")
	procDestroyIcon            = libuser32.NewProc("DestroyIcon")
	procDispatchMessageW       = libuser32.NewProc("DispatchMessageW")
	procGetMessageW            = libuser32.NewProc("GetMessageW")
	procGetModuleHandleW       = libkernel32.NewProc("GetModuleHandleW")
	procLoadCursorW            = libuser32.NewProc("LoadCursorW")
	procLoadIconW              = libuser32.NewProc("LoadIconW")
	procLoadImageW             = libuser32.NewProc("LoadImageW")
	procPostQuitMessage        = libuser32.NewProc("PostQuitMessage")
	procRegisterClassExW       = libuser32.NewProc("RegisterClassExW")
	procSendMessageW           = libuser32.NewProc("SendMessageW")
	procShell_NotifyIconW      = libshell32.NewProc("Shell_NotifyIconW")
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

func DefWindowProc(hWnd uintptr, Msg uint32, wParam, lParam uintptr) uintptr {
	var r1 uintptr
	r1, _, _ = procDefWindowProcW.Call(hWnd, uintptr(Msg), wParam, lParam)
	return r1
}

func DestroyIcon(hIcon uintptr) BOOL {
	var r1 uintptr
	r1, _, lastError = procDestroyIcon.Call(hIcon)
	return BOOL(r1)
}

func DispatchMessage(lpMsg *MSG) uintptr {
	var r1 uintptr
	r1, _, _ = procDispatchMessageW.Call(uintptr(unsafe.Pointer(lpMsg)))
	return r1
}

func GetMessage(lpMsg *MSG, hWnd uintptr, uMsgFilterMin, uMsgFilterMax uint32) BOOL {
	var r1 uintptr
	r1, _, lastError = procGetMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(uMsgFilterMin),
		uintptr(uMsgFilterMax))
	return BOOL(r1)
}

func GetModuleHandle(lpModuleName *uint16) uintptr {
	var r1 uintptr
	r1, _, lastError = procGetModuleHandleW.Call(uintptr(unsafe.Pointer(lpModuleName)))
	return r1
}

func LoadCursor(hInstance uintptr, lpCursorName *uint16) uintptr {
	var r1 uintptr
	r1, _, lastError = procLoadCursorW.Call(hInstance, uintptr(unsafe.Pointer(lpCursorName)))
	return r1
}

func LoadIcon(hInstance uintptr, lpIconName *uint16) uintptr {
	var r1 uintptr
	r1, _, lastError = procLoadIconW.Call(hInstance, uintptr(unsafe.Pointer(lpIconName)))
	return r1
}

func LoadImage(hInst uintptr, name *uint16, type_ uint32, cx, cy int32, fuLoad uint32) uintptr {
	var r1 uintptr
	r1, _, lastError = procLoadImageW.Call(hInst, uintptr(unsafe.Pointer(name)), uintptr(type_), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	return r1
}

func PostQuitMessage(nExitCode int32) {
	procPostQuitMessage.Call(uintptr(nExitCode))
}

func RegisterClassEx(Arg1 *WNDCLASSEX) uint16 {
	var r1 uintptr
	r1, _, lastError = procRegisterClassExW.Call(uintptr(unsafe.Pointer(Arg1)))
	return uint16(r1)
}

func SendMessage(hWnd uintptr, Msg uint32, wParam, lParam uintptr) uintptr {
	var r1 uintptr
	r1, _, lastError = procSendMessageW.Call(hWnd, uintptr(Msg), wParam, lParam)
	return r1
}

func Shell_NotifyIcon(dwMessage uint32, lpData *NOTIFYICONDATA) BOOL {
	var r1 uintptr
	r1, _, _ = procShell_NotifyIconW.Call(uintptr(dwMessage), uintptr(unsafe.Pointer(lpData)))
	return BOOL(r1)
}

func ShowWindow(hWnd uintptr, nCmdShow int32) BOOL {
	var r1 uintptr
	r1, _, _ = procShowWindow.Call(hWnd, uintptr(nCmdShow))
	return BOOL(r1)
}

func TranslateMessage(lpMsg *MSG) BOOL {
	var r1 uintptr
	r1, _, _ = procTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return BOOL(r1)
}

func UpdateWindow(hWnd uintptr) BOOL {
	var r1 uintptr
	r1, _, lastError = procUpdateWindow.Call(hWnd)
	return BOOL(r1)
}

func WTSEnumerateProcesses(
	hServer uintptr,
	Reserved, Version uint32,
	ppProcessInfo **WTS_PROCESS_INFO,
	pCount *uint32) BOOL {
	var r1 uintptr
	r1, _, lastError = procWTSEnumerateProcessesW.Call(
		hServer,
		uintptr(Reserved),
		uintptr(Version),
		uintptr(unsafe.Pointer(ppProcessInfo)),
		uintptr(unsafe.Pointer(pCount)))
	return BOOL(r1)
}

func WTSFreeMemory(pMemory unsafe.Pointer) {
	procWTSFreeMemory.Call(uintptr(pMemory))
}
