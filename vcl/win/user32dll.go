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
	user32dll         = syscall.NewLazyDLL("user32.dll")
	_MessageBox       = user32dll.NewProc("MessageBoxW")
	_LoadIcon         = user32dll.NewProc("LoadIconW")
	_GetClientRect    = user32dll.NewProc("GetClientRect")
	_GetSystemMetrics = user32dll.NewProc("GetSystemMetrics")

	_CallWindowProc = user32dll.NewProc("CallWindowProcW")
)

// MessageBox 消息框
func MessageBox(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	r, _, _ := _MessageBox.Call(hWnd, api.GoStrToDStr(lpText), api.GoStrToDStr(lpCaption), uintptr(uType))
	return int32(r)
}

// LoadIcon 从实例资源中加载icon
func LoadIcon(hInstance uintptr, lpIconName int) types.HICON {
	r, _, _ := _LoadIcon.Call(hInstance, uintptr(lpIconName))
	return types.HICON(r)
}

// GetClientRect 获取指定句柄客户区矩形
func GetClientRect(hWnd types.HWND) types.TRect {
	r := types.TRect{}
	_GetClientRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&r)))
	return r
}

// GetSystemMetrics
func GetSystemMetrics(nIndex int32) int32 {
	r, _, _ := _GetSystemMetrics.Call(uintptr(nIndex))
	return int32(r)
}

// GetWindowLongPtr
func GetWindowLongPtr(hWnd types.HWND, nIndex int32) int {
	r, _, _ := _GetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex))
	return int(r)
}

// SetWindowLongPtr
func SetWindowLongPtr(hWnd types.HWND, nIndex int32, dwNewLong uintptr) uintptr {
	r, _, _ := _SetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex), dwNewLong)
	return r
}

// CallWindowProc
func CallWindowProc(lpPrevWndFunc uintptr, hWnd types.HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := _CallWindowProc.Call(lpPrevWndFunc, uintptr(hWnd), uintptr(Msg), wParam, lParam)
	return r
}
