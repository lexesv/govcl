package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

var (
	mainForm *vcl.TForm
	richEdit *vcl.TRichEdit
)

func main() {
	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	initMainForm()
	initMainMenu()
	tlbar := vcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)

	richEdit = vcl.NewRichEdit(mainForm)
	richEdit.SetParent(mainForm)
	richEdit.SetAlign(types.AlClient)
	richEdit.Lines().Add("这是一段文字红色，粗体，斜體")
	richEdit.SetSelStart(6)
	richEdit.SetSelLength(2)
	richEdit.SelAttributes().SetColor(types.ClRed)

	richEdit.SetSelStart(9)
	richEdit.SetSelLength(2)
	style := types.TFontStyles(rtl.Include(0, types.FsBold))
	richEdit.SelAttributes().SetStyle(style)

	richEdit.SetSelStart(12)
	richEdit.SetSelLength(2)
	style = types.TFontStyles(rtl.Include(0, types.FsItalic))
	richEdit.SelAttributes().SetStyle(style)

	richEdit.SetSelStart(15)

	initRichEditPopupMenu()

	stabar := vcl.NewStatusBar(mainForm)
	stabar.SetParent(mainForm)

	vcl.Application.Run()
}

func initMainForm() {
	mainForm = vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
}

func initMainMenu() {
	mainMenu := vcl.NewMainMenu(mainForm)

	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("&File")
	mainMenu.Items().Add(item)

	item = vcl.NewMenuItem(mainForm)
	item.SetCaption("&Help")
	mainMenu.Items().Add(item)
}

func initRichEditPopupMenu() {
	pm := vcl.NewPopupMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("&Clear")
	pm.Items().Add(item)

	richEdit.SetPopupMenu(pm)
}
