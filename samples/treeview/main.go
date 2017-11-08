package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

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
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(500)

	// -----------TreeView 不同Node弹出不同菜单，两个右键例程不同使用

	tvpm1 := vcl.NewPopupMenu(mainForm)
	mItem := vcl.NewMenuItem(mainForm)
	mItem.SetCaption("第一种")
	tvpm1.Items().Add(mItem)

	tvpm2 := vcl.NewPopupMenu(mainForm)
	mItem = vcl.NewMenuItem(mainForm)
	mItem.SetCaption("第二种")
	tvpm2.Items().Add(mItem)

	tv := vcl.NewTreeView(mainForm)
	tv.SetParent(mainForm)
	tv.SetAlign(types.AlClient)

	// 自动展开
	//tv.SetAutoExpand(true)

	tv.SetOnClick(func(vcl.IObject) {
		node := tv.Selected()
		if node != nil && node.IsValid() {
			fmt.Println("Text:", node.Text(), ", Level:", node.Level(), ", Index:", node.Index(), ", hasChild:", node.HasChildren())
		}
	})

	tv.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		if button == types.MbRight {
			node := tv.GetNodeAt(x, y)
			if node != nil && node.IsValid() {
				// 自由决择是否选中
				node.SetSelected(true)
				// 根据Level来判断，这里只是做演示
				p := vcl.Mouse.CursorPos()
				switch node.Level() {
				case 0:
					tvpm1.Popup(p.X, p.Y)
				case 1:
					tvpm2.Popup(p.X, p.Y)
				}
				fmt.Println("node.Level():", node.Level(), ", text:", node.Text())
			}
		}
	})

	//	tv.Items().Clear()
	// 第一个节点
	node := tv.Items().AddChild(nil, "首个")

	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv.Items().EndUpdate()
	// 展开
	node.Expand(true)

	// 第二个节点
	node = tv.Items().AddChild(nil, "第二个节点")
	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv.Items().EndUpdate()

	vcl.Application.Run()
}
