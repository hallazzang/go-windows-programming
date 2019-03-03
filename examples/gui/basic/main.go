package main

import (
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

func wndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}

func main() {
	hInstance := win.GetModuleHandle(nil)

	windowClass, err := windows.UTF16PtrFromString("MyWindow")
	if err != nil {
		panic(err)
	}

	var wcex win.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wcex.LpfnWndProc = windows.NewCallback(wndProc)
	wcex.CbClsExtra = 0
	wcex.CbWndExtra = 0
	wcex.HInstance = hInstance
	wcex.HIcon = win.LoadIcon(hInstance, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	wcex.HCursor = win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))
	wcex.HbrBackground = win.COLOR_WINDOW + 1
	wcex.LpszMenuName = nil
	wcex.LpszClassName = windowClass
	wcex.HIconSm = win.LoadIcon(hInstance, win.MAKEINTRESOURCE(win.IDI_APPLICATION))

	if win.RegisterClassEx(&wcex) == 0 {
		panic(win.GetLastError())
	}
}
