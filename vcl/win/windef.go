package win

import (
	"syscall"
	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/types"
)

// winapi 非libvcl的
var (
	// user32.dll
	user32dll   = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32dll.NewProc("MessageBoxW")
	loadIconW   = user32dll.NewProc("LoadIconW")

	// kernel32.dll
	kernel32dll       = syscall.NewLazyDLL("kernel32.dll")
	isWow64Process    = kernel32dll.NewProc("IsWow64Process")
	getCurrentProcess = kernel32dll.NewProc("GetCurrentProcess")
	getModuleHandleW  = kernel32dll.NewProc("GetModuleHandleW")
)

// MessageBoxW 消息框
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := messageBoxW.Call(hWnd, api.GoStrToDStr(lpText), api.GoStrToDStr(lpCaption), uintptr(uType))
	return int32(r)
}

// LoadIconW 从实例资源中加载icon
func LoadIconW(hInstance uintptr, lpIconName string) types.HICON {
	r, _, _ := loadIconW.Call(hInstance, api.GoStrToDStr(lpIconName))
	return types.HICON(r)
}

// LoadIconW 从实例资源中加载icon
func LoadIconId(hInstance uintptr, id int) types.HICON {
	r, _, _ := loadIconW.Call(hInstance, uintptr(id))
	return types.HICON(r)
}

// GetModuleHandleW 获取当前是实例句柄，可传空
func GetModuleHandleW(lpModuleName string) uintptr {
	r, _, _ := getModuleHandleW.Call(api.GoStrToDStr(lpModuleName))
	return r
}

// GetSelfModuleHandle 获取自身模块实例句柄
func GetSelfModuleHandle() uintptr {
	r, _, _ := getModuleHandleW.Call(0)
	return r
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
