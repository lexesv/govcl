package main

import (
	"gitee.com/ying32/govcl/vcl"

	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetDoubleBuffered(true) // 最好开启，以免闪烁
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(400)

	btnClose := vcl.NewImageButton(mainForm)
	btnClose.SetParent(mainForm)

	btnClose.SetImageCount(4)
	btnClose.SetAutoSize(true)
	btnClose.Picture().LoadFromFile("./btn_close.png")
	btnClose.SetLeft(mainForm.ClientWidth() - btnClose.Width() - 3)

	btnClose.SetOnClick(func(vcl.IObject) {
		vcl.ShowMessage("close")
	})

	btnMax := vcl.NewImageButton(mainForm)
	btnMax.SetParent(mainForm)
	btnMax.SetImageCount(4)
	btnMax.SetAutoSize(true)
	btnMax.Picture().LoadFromFile("./btn_max.png")
	btnMax.SetLeft(btnClose.Left() - btnMax.Width())

	btnMin := vcl.NewImageButton(mainForm)
	btnMin.SetParent(mainForm)
	btnMin.SetImageCount(4)
	btnMin.SetAutoSize(true)
	btnMin.Picture().LoadFromFile("./btn_min.png")
	btnMin.SetLeft(btnMax.Left() - btnMin.Width())

	btnSkin := vcl.NewImageButton(mainForm)
	btnSkin.SetParent(mainForm)
	btnSkin.SetImageCount(3)
	btnSkin.SetAutoSize(true)
	btnSkin.Picture().LoadFromFile("./btn_skin.png")
	btnSkin.SetLeft(btnMin.Left() - btnSkin.Width())

	btnScan := vcl.NewImageButton(mainForm)
	btnScan.SetParent(mainForm)
	btnScan.SetImageCount(3)
	btnScan.SetAutoSize(true)
	btnScan.Picture().LoadFromFile("./btn_scan.png")
	btnScan.SetTop((mainForm.ClientHeight() - btnScan.Height()) / 2)
	btnScan.SetLeft((mainForm.ClientWidth() - btnScan.Width()) / 2)

	vcl.Application.Run()
}
