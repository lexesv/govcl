package win

import (
	"syscall"
	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
)

// winapi 非libvcl的
var (
	// user32.dll
	user32dll   = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32dll.NewProc("MessageBoxW")

	// kernel32.dll
	kernel32dll       = syscall.NewLazyDLL("kernel32.dll")
	isWow64Process    = kernel32dll.NewProc("IsWow64Process")
	getCurrentProcess = kernel32dll.NewProc("GetCurrentProcess")
)

// MessageBoxW 消息框
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := messageBoxW.Call(hWnd, api.GoStrToDStr(lpText), api.GoStrToDStr(lpCaption), uintptr(uType))
	return int32(r)
}

// IsWow64Process 检测进程是否运行在64位下
func IsWow64Process(hProcess uintptr) bool {
	if isWow64Process.Find() != nil {
		return false
	}
	var wow64Process int32
	r, _, _ := isWow64Process.Call(hProcess, uintptr(unsafe.Pointer(&wow64Process)))
	if r != 0 && wow64Process != 0 {
		return true
	}
	return false
}

// GetCurrentProcess 返回当前进程伪句柄
func GetCurrentProcess() uintptr {
	r, _, _ := getCurrentProcess.Call()
	return r
}

// IsWow64 判断当前进程是否运行在64上
func IsWow64() bool {
	return IsWow64Process(GetCurrentProcess())
}
