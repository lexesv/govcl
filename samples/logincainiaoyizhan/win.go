package main

import (
	"syscall"

	"gitee.com/ying32/govcl/vcl/types"
)

var (
	user32dll           = syscall.NewLazyDLL("user32.dll")
	openIcon            = user32dll.NewProc("OpenIcon")
	setForegroundWindow = user32dll.NewProc("SetForegroundWindow")
	closeWindow         = user32dll.NewProc("CloseWindow")
)

const (
	WM_SYSCOMMAND = 0x0112
	SC_RESTORE    = 61728
)

// function OpenIcon(hWnd: HWND): BOOL; stdcall;
// function OpenIcon; external user32 name 'OpenIcon';
// OpenIcon
func OpenIcon(hWnd types.HWND) bool {
	r, _, _ := openIcon.Call(uintptr(hWnd))
	return r != 0
}

// function SetForegroundWindow(hWnd: HWND): BOOL; stdcall;
// function SetForegroundWindow; external user32 name 'SetForegroundWindow';
func SetForegroundWindow(hWnd types.HWND) bool {
	r, _, _ := setForegroundWindow.Call(uintptr(hWnd))
	return r != 0
}

// function CloseWindow(hWnd: HWND): BOOL; stdcall;
// function CloseWindow; external user32 name 'CloseWindow';
func CloseWindow(hWnd types.HWND) bool {
	r, _, _ := closeWindow.Call(uintptr(hWnd))
	return r != 0
}
