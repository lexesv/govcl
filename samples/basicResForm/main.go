package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/exts/tools"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	// mac下发布时去掉，只在测试时使用
	tools.RunWithMacOSApp()

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		vcl.ShowMessage(e.Message())
	})
	//   Form1.gfm
	vcl.Application.CreateFormFromFile("Form1.gfm", &Form1)
	// 文件加载方式
	//	vcl.Application.CreateFormFromFile("Form2.gfm", &Form2)
	// 字节加载方式
	vcl.Application.CreateFormFromBytes(form2Bytes, &Form2)

	fmt.Println(Form1.Caption())
	Form1.Button1.SetOnClick(func(sender vcl.IObject) {
		//vcl.ShowMessage("Hello!")
		jpg := vcl.NewJPEGImage()
		defer jpg.Free()
		jpg.LoadFromFile("a.jpg")
	})

	Form1.CheckBox1.SetOnClick(func(sender vcl.IObject) {
		Form1.Button1.SetEnabled(Form1.CheckBox1.Checked())

	})

	Form1.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		fmt.Println("关闭。")
	})

	Form1.ActFileNew.SetOnExecute(func(vcl.IObject) {
		vcl.ShowMessage("ActFileNew Execute.")
	})
	Form1.Button2.SetOnClick(func(vcl.IObject) {
		result := Form2.ShowModal()
		if result == types.MrOk {
			vcl.ShowMessage("Form2返回了OK")
		} else if result == types.MrClose || result == types.MrNone {
			vcl.ShowMessage("Form2返回了Close")
		} else if result == types.MrCancel {
			vcl.ShowMessage("Form2返回了Cancel")
		}
	})
	Form1.ActExit.SetOnExecute(func(vcl.IObject) {
		vcl.Application.Terminate()
	})

	vcl.Application.Run()

}
