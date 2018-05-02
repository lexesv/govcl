// +build windows

package win

var (
	_GetWindowLongPtrW = user32dll.NewProc("GetWindowLongPtrW")
	_SetWindowLongPtrW = user32dll.NewProc("SetWindowLongPtrW")
)
