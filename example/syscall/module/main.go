package main

import (
	"golang.org/x/sys/windows"
)

func main() {
	caption, err := windows.UTF16PtrFromString("Go Windows Programming")
	if err != nil {
		panic(err)
	}
	text, err := windows.UTF16PtrFromString("Calling MessageBox from Go!")
	if err != nil {
		panic(err)
	}

	messageBox(0, text, caption, 0)
}
