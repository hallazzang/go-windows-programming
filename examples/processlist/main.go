package main

import (
	"fmt"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
)

type processInfo struct {
	id        uint32
	name      string
	sessionID uint32
}

func processList() ([]processInfo, error) {
	var pProcessInfo *win.WTS_PROCESS_INFO
	var count uint32
	var ps []processInfo
	if win.WTSEnumerateProcesses(win.WTS_CURRENT_SERVER_HANDLE, 0, 1, &pProcessInfo, &count) == 0 {
		return nil, win.GetLastError()
	}
	defer win.WTSFreeMemory(unsafe.Pointer(pProcessInfo))
	size := unsafe.Sizeof(win.WTS_PROCESS_INFO{})
	for i := uint32(0); i < count; i++ {
		p := *(*win.WTS_PROCESS_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(pProcessInfo)) + uintptr(size)*uintptr(i)))
		ps = append(ps, processInfo{
			id:        p.ProcessId,
			name:      win.UTF16PtrToString(p.PProcessName),
			sessionID: p.SessionId,
		})
	}
	return ps, nil
}

func main() {
	ps, err := processList()
	if err != nil {
		panic(err)
	}
	fmt.Printf("process list:\n%+v\n", ps)
}
