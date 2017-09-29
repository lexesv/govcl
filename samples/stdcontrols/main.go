package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
)

func main() {
	icon := vcl.NewIcon()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)
	defer icon.Free()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(api.PoScreenCenter)
	mainForm.SetWidth(200)
	mainForm.SetHeight(700)

	var top int32 = 40
	// TButton
	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetLeft(10)
	btn.SetTop(top)
	btn.SetCaption("按钮1")
	btn.SetOnClick(func(vcl.IObject) {
		fmt.Println("按钮1单击")
	})

	top += btn.Height() + 5

	// TEdit
	edit := vcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetTextHint("提示")
	//edit.SetText("文字")
	//	edit.SetReadOnly(true)
	edit.SetOnChange(func(vcl.IObject) {
		fmt.Println("文字改变了")
	})

	top += edit.Height() + 5
	// TEdit Password
	edit = vcl.NewEdit(mainForm)
	edit.SetParent(mainForm)
	edit.SetLeft(10)
	edit.SetTop(top)
	edit.SetText("文字")
	edit.SetPasswordChar('*')

	top += edit.Height() + 5

	// TLabel
	lbl := vcl.NewLabel(mainForm)
	lbl.SetParent(mainForm)
	lbl.SetLeft(10)
	lbl.SetTop(top)
	lbl.SetCaption("标签1")
	lbl.Font().SetColor(255)

	top += lbl.Height() + 5

	// TCheckBox
	chk := vcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetLeft(10)
	chk.SetTop(top)
	chk.SetCaption("选择框1")
	chk.SetOnClick(func(vcl.IObject) {
		fmt.Println("checked: ", chk.Checked())
	})

	// TStatusBar
	stat := vcl.NewStatusBar(mainForm)
	stat.SetParent(mainForm)
	//stat.SetSizeGrip(false) // 右解是否有可调的
	spnl := stat.Panels().Add()
	spnl.SetText("第一个")
	spnl.SetWidth(80)

	spnl = stat.Panels().Add()
	spnl.SetText("第二个")
	spnl.SetWidth(80)

	// TToolBar
	tlbar := vcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)
	tlbar.SetShowCaptions(true)

	// 倒过来创建
	tlbtn := vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("2")
	tlbtn.SetStyle(api.TbsDropDown)

	tlbtn = vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetStyle(api.TbsSeparator)

	tlbtn = vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)
	tlbtn.SetCaption("1")

	top += chk.Height() + 5
	// TRadioButton
	rd := vcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(10)
	rd.SetTop(top)
	rd.SetCaption("选项1")

	var left int32 = rd.Left() + rd.Width() + 5

	rd = vcl.NewRadioButton(mainForm)
	rd.SetParent(mainForm)
	rd.SetLeft(left)
	rd.SetTop(top)
	rd.SetCaption("选项2")

	top += rd.Height() + 5
	// TMemo
	mmo := vcl.NewMemo(mainForm)
	mmo.SetParent(mainForm)
	mmo.SetBounds(10, top, 167, 50)
	//    mmo.Text()
	mmo.Lines().Add("1")
	mmo.Lines().Add("2")

	top += mmo.Height() + 5
	// TComboBox
	cb := vcl.NewComboBox(mainForm)
	cb.SetParent(mainForm)
	cb.SetLeft(10)
	cb.SetTop(top)
	cb.SetStyle(api.CsDropDownList)
	cb.Items().Add("1")
	cb.Items().Add("2")
	cb.Items().Add("3")
	cb.SetItemIndex(0)
	cb.SetOnChange(func(vcl.IObject) {
		if cb.ItemIndex() != -1 {
			fmt.Println(cb.Items().Strings(cb.ItemIndex()))
		}
	})

	// TListBox
	top += cb.Height() + 5
	lst := vcl.NewListBox(mainForm)
	lst.SetParent(mainForm)
	lst.SetBounds(10, top, 167, 50)
	lst.Items().Add("1")
	lst.Items().Add("2")
	lst.Items().Add("3")

	// TPanel
	top += lst.Height() + 5
	pnl := vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetCaption("fff")
	//    pnl.SetShowCaption(false)
	pnl.SetBounds(10, top, 167, 50)

	// color
	top += pnl.Height() + 5
	clr := vcl.NewColorBox(mainForm)
	clr.SetParent(mainForm)
	clr.SetLeft(10)
	clr.SetTop(top)
	clr.SetOnChange(func(vcl.IObject) {
		if clr.ItemIndex() != -1 {
			lbl.Font().SetColor(clr.Selected())
		}
	})

	// TPageControl
	top += clr.Height() + 5
	pgc := vcl.NewPageControl(mainForm)
	pgc.SetParent(mainForm)
	pgc.SetBounds(10, top, 167, 100)
	sheet := vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("一")
	sheet = vcl.NewTabSheet(mainForm)
	sheet.SetPageControl(pgc)
	sheet.SetCaption("二")

	// TImage
	top += pgc.Height() + 5
	img := vcl.NewImage(mainForm)
	img.SetBounds(10, top, 167, 97)
	img.SetParent(mainForm)
	img.Picture().LoadFromFile("1.jpg")
	//img.SetStretch(true)
	img.SetProportional(true)

	// run
	vcl.Application.Run()
}
