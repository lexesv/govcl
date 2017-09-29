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
	// 双缓冲
	mainForm.SetDoubleBuffered(true)

	lv1 := vcl.NewListView(mainForm)
	lv1.SetParent(mainForm)
	lv1.SetAlign(api.AlClient)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetViewStyle(api.VsReport)
	lv1.SetGridLines(true)
	lv1.SetColumnClick(false)

	col := lv1.Columns().Add()
	col.SetCaption("序号")
	col.SetWidth(100)

	col = lv1.Columns().Add()
	col.SetCaption("项目1")
	col.SetWidth(200)

	lv1.SetOnClick(func(vcl.IObject) {
		if lv1.ItemIndex() != -1 {
			item := lv1.Selected() // lv1.Items().Item(lv1.ItemIndex())
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})

	//	lv1.Clear()
	lv1.Items().BeginUpdate()
	for i := 1; i <= 100; i++ {
		item := lv1.Items().Add()
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	lv1.Items().EndUpdate()

	vcl.Application.Run()
}
