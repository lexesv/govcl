package api

import (
	"syscall"
)

// winapi 非libvcl的
var (
	user32dll   = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32dll.NewProc("MessageBoxW")
)

// MessageBoxW 消息框
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := messageBoxW.Call(hWnd, GoStrToDStr(lpText), GoStrToDStr(lpCaption), uintptr(uType))
	return int32(r)
}
