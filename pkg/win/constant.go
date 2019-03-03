package win

// GetSysColor constants
const (
	COLOR_WINDOW = 5
)

// Window class styles
const (
	CS_HREDRAW = 0x0002
	CS_VREDRAW = 0x0001
)

const (
	CW_USEDEFAULT = ^0x7fffffff
)

// LoadCursor constants
const (
	IDC_ARROW = 32512
)

// LoadIcon constants
const (
	IDI_APPLICATION = 32512
)

// ShowWindow constants
const (
	SW_HIDE        = 0
	SW_SHOW        = 5
	SW_SHOWDEFAULT = 10
	SW_SHOWNORMAL  = 1
)

// Window messages
const (
	WM_DESTROY = 0x0002
)

// Window styles
const (
	WS_CAPTION          = 0x00c00000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_MINIMIZEBOX      = 0x00020000
	WS_OVERLAPPED       = 0x00000000
	WS_SYSMENU          = 0x00080000
	WS_THICKFRAME       = 0x00040000
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
)

// WTS API constants
const (
	WTS_CURRENT_SERVER_HANDLE = 0
)
