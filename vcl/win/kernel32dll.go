// +build windows

package win

import (
	"syscall"
	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
)

// winapi 非libvcl的
var (

	// kernel32.dll
	kernel32dll         = syscall.NewLazyDLL("kernel32.dll")
	isWow64Process      = kernel32dll.NewProc("IsWow64Process")
	getCurrentProcess   = kernel32dll.NewProc("GetCurrentProcess")
	getModuleHandleW    = kernel32dll.NewProc("GetModuleHandleW")
	getVersionExW       = kernel32dll.NewProc("GetVersionExW")
	getNativeSystemInfo = kernel32dll.NewProc("GetNativeSystemInfo")
)

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

// 获取系统版本信息
func GetVersionExW(lpVersionInformation *TOSVersionInfoExW) bool {
	if lpVersionInformation != nil {
		lpVersionInformation.OSVersionInfoSize = uint32(unsafe.Sizeof(*lpVersionInformation))
	}
	r, _, _ := getVersionExW.Call(uintptr(unsafe.Pointer(lpVersionInformation)))
	return r != 0
}

// GetNativeSystemInfo
func GetNativeSystemInfo(lpSystemInformation *TSystemInfo) bool {
	r, _, _ := getNativeSystemInfo.Call(uintptr(unsafe.Pointer(lpSystemInformation)))
	return r != 0
}
