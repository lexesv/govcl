// +build windows

package win

var (
	_GetWindowLongPtrW = user32dll.NewProc("GetWindowLongW")
	_SetWindowLongPtrW = user32dll.NewProc("SetWindowLongW")
)
