package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
)

// 类似一个继承的，用于自动填充字段
type TMainForm struct {
	*vcl.TForm
	Button1   *vcl.TButton
	CheckBox1 *vcl.TCheckBox
	ComboBox1 *vcl.TComboBox
	ListBox1  *vcl.TListBox
}

var mainForm *TMainForm

func main() {

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	vcl.Application.CreateFormFromFile("./Form1.gfm", &mainForm)
	fmt.Println(mainForm.Button1)
	mainForm.Button1.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("Hello!")
	})

	mainForm.CheckBox1.SetOnClick(func(sender vcl.IObject) {
		mainForm.Button1.SetEnabled(mainForm.CheckBox1.Checked())
		vcl.ShowMessage(mainForm.ComboBox1.Text())
	})

	mainForm.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		fmt.Println("关闭。")
	})

	vcl.Application.Run()
}
