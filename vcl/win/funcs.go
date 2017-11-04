package win

import (
	"gitee.com/ying32/govcl/vcl/api"
)

// MessageBoxW 消息框 
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType uint32) int32 {
	return api.MessageBoxW(hWnd, lpText, lpCaption, uType)
}
