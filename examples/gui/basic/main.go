package main

// https://docs.microsoft.com/en-us/cpp/windows/walkthrough-creating-windows-desktop-applications-cpp?view=vs-2017

import (
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

func wndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func main() {
	hInstance := win.GetModuleHandle(nil)

	windowClass, err := windows.UTF16PtrFromString("MyWindow")
	if err != nil {
		panic(err)
	}
	title, err := windows.UTF16PtrFromString("Go Windows Programming")
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

	hWnd := win.CreateWindowEx(
		0,
		windowClass,
		title,
		win.WS_OVERLAPPEDWINDOW,
		win.CW_USEDEFAULT,
		win.CW_USEDEFAULT,
		500,
		100,
		0,
		0,
		hInstance,
		nil)
	if hWnd == 0 {
		panic(win.GetLastError())
	}

	win.ShowWindow(hWnd, win.SW_SHOWNORMAL)
	win.UpdateWindow(hWnd)

	var msg win.MSG
	for win.GetMessage(&msg, 0, 0, 0) != 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}
