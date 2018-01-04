package rtl

import "gitee.com/ying32/govcl/vcl/win"

// MainInstance EXE自身的实例
func MainInstance() uintptr {
	return win.GetSelfModuleHandle()
}
