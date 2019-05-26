package main

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

const notifyIconMsg = win.WM_APP + 1

var errShellNotifyIcon = errors.New("Shell_NotifyIcon error")

func init() {
	rand.Seed(time.Now().UnixNano())
}

type notifyIcon struct {
	hwnd uintptr
	guid win.GUID
}

func newNotifyIcon(hwnd uintptr) (*notifyIcon, error) {
	ni := &notifyIcon{
		hwnd: hwnd,
		guid: newGUID(),
	}
	data := ni.newData()
	data.UFlags |= win.NIF_MESSAGE
	data.UCallbackMessage = notifyIconMsg
	if win.Shell_NotifyIcon(win.NIM_ADD, data) == win.FALSE {
		return nil, errShellNotifyIcon
	}
	return ni, nil
}

func (ni *notifyIcon) Dispose() {
	win.Shell_NotifyIcon(win.NIM_DELETE, ni.newData())
}

func (ni *notifyIcon) SetTooltip(tooltip string) error {
	data := ni.newData()
	data.UFlags |= win.NIF_TIP
	copy(data.SzTip[:], windows.StringToUTF16(tooltip))
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func (ni *notifyIcon) SetIcon(hIcon uintptr) error {
	data := ni.newData()
	data.UFlags |= win.NIF_ICON
	data.HIcon = hIcon
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func (ni *notifyIcon) ShowNotification(title, text string) error {
	data := ni.newData()
	data.UFlags |= win.NIF_INFO
	copy(data.SzInfoTitle[:], windows.StringToUTF16(title))
	copy(data.SzInfo[:], windows.StringToUTF16(text))
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func (ni *notifyIcon) newData() *win.NOTIFYICONDATA {
	var nid win.NOTIFYICONDATA
	nid.CbSize = uint32(unsafe.Sizeof(nid))
	nid.UFlags = win.NIF_GUID
	nid.HWnd = ni.hwnd
	nid.GuidItem = ni.guid
	return &nid
}

func newGUID() win.GUID {
	var buf [16]byte
	rand.Read(buf[:])
	return *(*win.GUID)(unsafe.Pointer(&buf[0]))
}
