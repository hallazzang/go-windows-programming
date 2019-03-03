package win

type MSG struct {
	Hwnd     uintptr
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       POINT
	LPrivate uint32
}

type POINT struct {
	X int32
	Y int32
}

type SID struct {
	Revision            byte
	SubAuthorityCount   byte
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
}

type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]byte
}

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         uintptr
	HCursor       uintptr
	HbrBackground uintptr
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       uintptr
}

type WTS_PROCESS_INFO struct {
	SessionId    uint32
	ProcessId    uint32
	PProcessName *uint16
	PUserSid     *SID
}
