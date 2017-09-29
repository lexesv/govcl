// +build linux,darwin !cgo

package api

import (
	"unsafe"
)

func DStrToGoStr(ustr uintptr) string {
	if ustr == 0 {
		return ""
	}
	str := make([]byte, DStrLen(ustr))
	DMove(ustr, uintptr(unsafe.Pointer(&str[0])), len(str))
	return string(str)
}

func GoStrToDStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(&([]byte(s)[0])))
}
