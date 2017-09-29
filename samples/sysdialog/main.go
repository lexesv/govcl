package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	//    vcl.SetIcon()

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	dlgOpen := vcl.NewOpenDialog(mainForm)
	dlgOpen.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	//    dlgOpen.SetInitialDir()
	//	dlgOpen.SetFilterIndex()

	dlgOpen.SetOptions(types.TOpenOptions(rtl.Include(uint32(dlgOpen.Options()), types.OfShowHelp)))
	dlgOpen.SetTitle("打开")

	btn := vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Open Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlgOpen.Execute(mainForm.Handle()) {
			fmt.Println("filename: ", dlgOpen.FileName())
		}
	})

	dlSave := vcl.NewSaveDialog(mainForm)
	dlSave.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	dlSave.SetOptions(types.TOpenOptions(rtl.Include(uint32(dlSave.Options()), types.OfShowHelp)))
	dlSave.SetTitle("保存")

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Save Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlSave.Execute(mainForm.Handle()) {
			fmt.Println("filename: ", dlSave.FileName())
		}
	})

	dlFont := vcl.NewFontDialog(mainForm)

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Font Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlFont.Execute(mainForm.Handle()) {
			fmt.Println("Name: ", dlFont.Font().Name())
		}
	})

	vcl.Application.Run()
}
