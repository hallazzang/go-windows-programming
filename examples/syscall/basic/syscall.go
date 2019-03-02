package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	// Prepare libraries and procedures
	libuser32 := windows.NewLazySystemDLL("user32.dll")
	procMessageBoxW := libuser32.NewProc("MessageBoxW")

	// Encode Go string with UTF-16
	caption, err := windows.UTF16PtrFromString("Go Windows Programming")
	if err != nil {
		panic(err)
	}
	text, err := windows.UTF16PtrFromString("Calling MessageBox from Go!")
	if err != nil {
		panic(err)
	}

	// Call the procedure
	procMessageBoxW.Call(
		0,                                // HWND hWnd
		uintptr(unsafe.Pointer(text)),    // LPCWSTR lpText
		uintptr(unsafe.Pointer(caption)), // LPCWSTR lpCaption
		0)                                // UINT uType
}
