package api

import (
	"syscall"
	"unsafe"
)

func DStrToGoStr(ustr uintptr) string {
	str := make([]uint16, DStrLen(ustr))
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), len(str)*2)
	return syscall.UTF16ToString(str)
}

func GoStrToDStr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}
