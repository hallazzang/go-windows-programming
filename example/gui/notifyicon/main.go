package main

import (
	"log"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

func wndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case notifyIconMsg:
		switch nmsg := win.LOWORD(uint32(lParam)); nmsg {
		case win.NIN_BALLOONUSERCLICK:
			log.Print("User has clicked the balloon message")
		case win.WM_LBUTTONDOWN:
			clickHandler()
		}
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func createMainWindow() (uintptr, error) {
	hInstance := win.GetModuleHandle(nil)

	wndClass := windows.StringToUTF16Ptr("MyWindow")

	var wcex win.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wcex.LpfnWndProc = windows.NewCallback(wndProc)
	wcex.HInstance = hInstance
	wcex.HbrBackground = win.COLOR_WINDOW + 1
	wcex.LpszClassName = wndClass
	if win.RegisterClassEx(&wcex) == 0 {
		return 0, win.GetLastError()
	}

	hwnd := win.CreateWindowEx(0, wndClass, windows.StringToUTF16Ptr("NotifyIcon Example"), win.WS_OVERLAPPEDWINDOW, win.CW_USEDEFAULT, win.CW_USEDEFAULT, 400, 300, 0, 0, hInstance, nil)
	if hwnd == win.NULL {
		return 0, win.GetLastError()
	}
	win.ShowWindow(hwnd, win.SW_SHOW)

	return hwnd, nil
}

func loadIcon(name string) (uintptr, error) {
	pname, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	hIcon := win.LoadImage(win.NULL, pname, win.IMAGE_ICON, 0, 0, win.LR_DEFAULTSIZE|win.LR_LOADFROMFILE)
	if hIcon == win.NULL {
		return 0, win.GetLastError()
	}

	return hIcon, nil
}

func clickHandler() {
	log.Print("User has clicked the notify icon")
}

func main() {
	hwnd, err := createMainWindow()
	if err != nil {
		panic(err)
	}

	hIcon, err := loadIcon("icon.ico")
	defer win.DestroyIcon(hIcon)

	win.SendMessage(hwnd, win.WM_SETICON, win.ICON_BIG, hIcon)
	win.SendMessage(hwnd, win.WM_SETICON, win.ICON_SMALL, hIcon)

	ni, err := newNotifyIcon(hwnd)
	if err != nil {
		panic(err)
	}
	defer ni.Dispose()

	ni.SetIcon(hIcon)
	ni.SetTooltip("NotifyIcon Example")
	ni.ShowNotification("Hello", "NotifyIcon!")

	var msg win.MSG
	for win.GetMessage(&msg, 0, 0, 0) != 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}
