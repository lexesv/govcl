package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/win"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	//    vcl.SetIcon()

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(api.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	vcl.ShowMessage("消息")
	if vcl.MessageDlg("消息", api.MtConfirmation, api.MbYes, api.MbNo) == vcl.IdYes {
		vcl.ShowMessage("你点击了“是")
	}
	if vcl.Application.MessageBox("消息", "标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == vcl.MrOk {
		vcl.ShowMessage("你点击了“是")
	}

	vcl.Application.Run()
}
