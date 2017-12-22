package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
)

func main() {

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	vcl.Application.CreateFormFromFile("Form1.gfm", &Form1)
	fmt.Println(Form1.Button1)
	Form1.Button1.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("Hello!")
	})

	Form1.CheckBox1.SetOnClick(func(sender vcl.IObject) {
		Form1.Button1.SetEnabled(Form1.CheckBox1.Checked())
		vcl.ShowMessage(Form1.ComboBox1.Text())
	})

	Form1.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		fmt.Println("关闭。")
	})

	vcl.Application.Run()
}
