package win

type SID struct {
	Revision            byte
	SubAuthorityCount   byte
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
}

type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]byte
}

type WTS_PROCESS_INFO struct {
	SessionId    uint32
	ProcessId    uint32
	PProcessName *uint16
	PUserSid     *SID
}
