package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

// 简单介绍下Delphi中控件的布局方式
func main() {

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)
	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
		*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
	})
	// 为方便演示，这里使用一个TPageControl作为多个布局的演示
	// Delphi中所有可视与非可视组件的owner都可设置为TForm的实例，
	// 本示例中，TForm的实例即mainForm，当owner设置为mainForm时
	// 在TForm销毁前会自动Free相关组件
	pgc := vcl.NewPageControl(mainForm)

	// 如需让一个可视控件显示到某个控件上即可使用SetParent，
	// 这里需要注意并非所有控件都可使用SetParent，Delphi中一般能接受子控件的只有
	// TWinControl，不过这里没有做区分。打个比方如TButton、TImage、TLabel是不可以作
	// 为父控件使用的，能作为父控件的常用的有TForm、TPanel、TPageControl、TTabSheet。
	pgc.SetParent(mainForm)
	// 这里将TPageControl设置为整个窗口客户区大小，并自动调整
	pgc.SetAlign(types.AlClient)

	//
	sheet := vcl.NewTabSheet(mainForm)
	// Delphi中有几个控件的设置Parent属性有些不同，这里需要使用SetPageControl才行。
	// TTabSheet控件默认的Align为alClient，即填充父控件客户区域
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-左-客户区")

	// 此时的pnl仅Height属性生效
	pnl := vcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClRed)
	pnl.SetParent(sheet)
	pnl.SetHeight(100)
	pnl.SetAlign(types.AlTop)

	// 此时的pnl仅Width属性生效
	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("左")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClGreen)
	pnl.SetParent(sheet)
	pnl.SetWidth(100)
	pnl.SetAlign(types.AlLeft)

	// 此时的pnl无法手动调整大小
	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("客户区")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)

	// --------------------------------------------------------------------

	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-客户区-底")

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClRed)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlTop)

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("客户区")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClGreen)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlClient)

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("底")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlBottom)

	//--------------------------------------------------------------------

	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("顶-客户区(左|-|右)-底")

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("顶")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClRed)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlTop)

	ppnl := vcl.NewPanel(mainForm)
	ppnl.SetCaption("客户区")
	ppnl.SetParentBackground(false)
	ppnl.SetColor(types.ClGreen)
	ppnl.SetParent(sheet)
	ppnl.SetAlign(types.AlClient)

	// 上个panel作为父控件
	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("左")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClAqua)
	pnl.SetParent(ppnl)
	pnl.SetAlign(types.AlLeft)

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("右")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClAzure)
	pnl.SetParent(ppnl)
	pnl.SetAlign(types.AlRight)

	pnl = vcl.NewPanel(mainForm)
	pnl.SetCaption("底")
	pnl.SetParentBackground(false)
	pnl.SetColor(types.ClBlue)
	pnl.SetParent(sheet)
	pnl.SetAlign(types.AlBottom)

	vcl.Application.Run()
}
