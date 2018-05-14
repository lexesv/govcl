package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func initMainMenu() {
	mainMenu := vcl.NewMainMenu(frmMain)
	mainMenu.SetAutoHotkeys(types.MaManual)
	// 第一列菜单
	item := vcl.NewMenuItem(frmMain)
	item.SetCaption("文件(&F)")
	mainMenu.Items().Add(item)
	// 第一列子菜单开始
	subMenu := vcl.NewMenuItem(frmMain)
	subMenu.SetAction(actExit)
	item.Add(subMenu)

	// 第二列菜单
	// 登录
	item = vcl.NewMenuItem(frmMain)
	item.SetAction(actLogin)
	mainMenu.Items().Add(item)

	// 第三列菜单
	// 注销
	item = vcl.NewMenuItem(frmMain)
	item.SetAction(actLogOff)
	mainMenu.Items().Add(item)

	// 第四列菜单
	// 菜鸟后台
	item = vcl.NewMenuItem(frmMain)
	item.SetCaption("菜鸟后台(&C)")
	item.SetOnClick(func(vcl.IObject) {
		rtl.SysOpen("https://cainiaoyizhan.com")
	})
	mainMenu.Items().Add(item)

	//关于列菜单
	item = vcl.NewMenuItem(frmMain)
	item.SetCaption("帮助(&H)")
	mainMenu.Items().Add(item)
	subMenu = vcl.NewMenuItem(frmMain)
	subMenu.SetCaption("关于")
	subMenu.SetOnClick(func(vcl.IObject) {
		frmAbout.ShowModal()
	})
	item.Add(subMenu)

}
