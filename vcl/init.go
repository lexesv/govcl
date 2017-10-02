package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

var (
	// 四个实例类，不需要Create和Free即可访问
	Application *TApplication
	Screen      *TScreen
	Mouse       *TMouse
	Clipboard   *TClipboard
	// StyleManager 没有实例类的，属于静态类
	StyleManager *TStyleManager
)

func init() {
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(callbackStdcall)

	// 导入四个实例类
	Application = ApplicationFromInst(Application_Instance())
	Screen = ScreenFromInst(Screen_Instance())
	Mouse = MouseFromInst(Mouse_Instance())
	Clipboard = ClipboardFromInst(Clipboard_Instance())
	StyleManager = NewStyleManager()

}
