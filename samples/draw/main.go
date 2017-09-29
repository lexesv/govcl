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
	mainForm.SetWidth(400)
	mainForm.SetHeight(600)

	mainForm.SetOnPaint(func(vcl.IObject) {

		canvas := mainForm.Canvas()

		canvas.MoveTo(10, 10)
		canvas.LineTo(50, 10)
		s := "这是一段文字"
		canvas.Font().SetColor(types.ClRed) // red
		canvas.Font().SetSize(20)
		style := canvas.Font().Style()
		canvas.Font().SetStyle(types.TFontStyles(rtl.Include(uint32(style), types.FsBold, types.FsItalic)))
		canvas.TextOut(100, 30, s)

		r := types.TRect{0, 0, 80, 80}

		// 计算文字
		fmt.Println("TfSingleLine: ", types.TfSingleLine)
		s = "由于现有第三方的Go UI库不是太宠大就是用的不习惯，或者组件太少。"
		canvas.TextRect1(&r, &s, types.TTextFormat(rtl.Include(0,
			types.TfCenter, types.TfVerticalCenter, types.TfSingleLine)))
		fmt.Println("r: ", r, ", s: ", s)

		s = "测试输出"
		r = types.TRect{0, 0, 80, 80}
		// brush
		canvas.Brush().SetColor(types.ClGreen)
		canvas.FillRect(r)

		// font
		canvas.Font().SetStyle(0)
		canvas.Font().SetSize(9)
		canvas.Font().SetColor(types.ClBlue)

		// pen
		canvas.Pen().SetWidth(2)
		canvas.Pen().SetColor(types.ClFuchsia)
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		textFmt := rtl.Include(0, types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		fmt.Println("format: ", textFmt)
		//		canvas.TextRect2(r, 0, 0, s)
		canvas.TextRect1(&r, &s, types.TTextFormat(textFmt))

		jpgimg := vcl.NewJPEGImage()
		defer jpgimg.Free()
		jpgimg.LoadFromFile("..\\..\\imgs\\1.jpg")

		canvas.Draw1(0, 80, jpgimg)
		//canvas.Draw2(0, 200, jpgimg, 10)

	})

	vcl.Application.Run()
}
