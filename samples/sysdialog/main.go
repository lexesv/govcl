package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(400)

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
		if dlgOpen.Execute() {
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
		if dlSave.Execute() {
			fmt.Println("filename: ", dlSave.FileName())
		}
	})

	dlFont := vcl.NewFontDialog(mainForm)

	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Font Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlFont.Execute() {
			fmt.Println("Name: ", dlFont.Font().Name())
		}
	})

	dlColor := vcl.NewColorDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Color Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlColor.Execute() {
			fmt.Println("Color: ", dlColor.Color())
		}
	})

	dlPicOpen := vcl.NewOpenPictureDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("OpenPic Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlPicOpen.Execute() {
			fmt.Println("Name: ", dlPicOpen.FileName())
		}
	})

	dlPicSave := vcl.NewSavePictureDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("SavePic Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlPicSave.Execute() {
			fmt.Println("Name: ", dlPicSave.FileName())
		}
	})

	dlTxtOpen := vcl.NewOpenTextFileDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Open Text Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlTxtOpen.Execute() {
			fmt.Println("Name: ", dlTxtOpen.FileName())
		}
	})

	dlTxtSave := vcl.NewSaveTextFileDialog(mainForm)
	btn = vcl.NewButton(mainForm)
	btn.SetAlign(types.AlTop)
	btn.SetParent(mainForm)
	btn.SetCaption("Save Text Dialog")
	btn.SetOnClick(func(vcl.IObject) {
		if dlTxtSave.Execute() {
			fmt.Println("Name: ", dlTxtSave.FileName())
		}
	})

	vcl.Application.Run()
}
