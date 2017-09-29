package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

var (
	Application *TApplication
	Screen      *TScreen
	Mouse       *TMouse
	Clipboard   *TClipboard
)

func init() {
	SetEventCallback(callbackStdcall)
	Application = ApplicationFromInst(Application_Instance())
	Screen = ScreenFromInst(Screen_Instance())
	Mouse = MouseFromInst(Mouse_Instance())
	Clipboard = ClipboardFromInst(Clipboard_Instance())
}
