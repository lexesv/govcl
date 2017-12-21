package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

var (
	frmAbout *vcl.TForm
)

func initfrmAbout() {
	frmAbout = vcl.Application.CreateForm()
	frmAbout.ScreenCenter()
	frmAbout.SetCaption("关于")
	frmAbout.SetBorderStyle(types.BsSingle)
	frmAbout.EnabledMaximize(false)
	frmAbout.EnabledMinimize(false)
	frmAbout.SetWidth(405)
	frmAbout.SetHeight(210)

	btn := vcl.NewButton(frmAbout)
	btn.SetParent(frmAbout)
	btn.SetCaption("OK")
	btn.SetModalResult(types.MbOK)
	btn.SetLeft(frmAbout.ClientWidth() - btn.Width() - 10)
	btn.SetTop(frmAbout.ClientHeight() - btn.Height() - 10)
}
