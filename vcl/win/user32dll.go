// +build windows

package win

import (
	"syscall"
	"unsafe"

	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/types"
)

var (
	// user32.dll
	user32dll        = syscall.NewLazyDLL("user32.dll")
	messageBoxW      = user32dll.NewProc("MessageBoxW")
	loadIconW        = user32dll.NewProc("LoadIconW")
	getClientRect    = user32dll.NewProc("GetClientRect")
	getSystemMetrics = user32dll.NewProc("GetSystemMetrics")
)

// MessageBoxW 消息框
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := messageBoxW.Call(hWnd, api.GoStrToDStr(lpText), api.GoStrToDStr(lpCaption), uintptr(uType))
	return int32(r)
}

// LoadIconW 从实例资源中加载icon
func LoadIconW(hInstance uintptr, lpIconName int) types.HICON {
	r, _, _ := loadIconW.Call(hInstance, uintptr(lpIconName))
	return types.HICON(r)
}

// GetClientRect 获取指定句柄客户区矩形
func GetClientRect(hWnd types.HWND) types.TRect {
	r := types.TRect{}
	getClientRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&r)))
	return r
}

// GetSystemMetrics
func GetSystemMetrics(nIndex int32) int32 {
	r, _, _ := getSystemMetrics.Call(uintptr(nIndex))
	return int32(r)
}
