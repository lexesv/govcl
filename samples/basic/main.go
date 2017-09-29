package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
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
	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose uintptr) {
		rtl.SetFormCanClose(CanClose, vcl.MessageDlg("是否退出？", api.MtConfirmation, api.MbYes, api.MbNo) == vcl.IdYes)
	})

	vcl.Application.Run()
}
