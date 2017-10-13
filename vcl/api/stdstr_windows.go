package api

import (
	"syscall"
	"unsafe"
)

func DStrToGoStr(ustr uintptr) string {
	l := DStrLen(ustr)
	if l == 0 {
		return ""
	}
	str := make([]uint16, l)
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), len(str)*2)
	return syscall.UTF16ToString(str)
}

func GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
