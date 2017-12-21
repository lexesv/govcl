package main

import (
	"gitee.com/ying32/govcl/vcl"
	//	"gitee.com/ying32/govcl/vcl/types"
)

var (
	frmMain  *vcl.TForm
	trayIcon *vcl.TTrayIcon
	session  *TCNSession
)

func initfrmMain() {
	frmMain = vcl.Application.CreateForm()
	frmMain.SetWidth(1280)
	frmMain.SetHeight(720)
	frmMain.ScreenCenter()
	frmMain.SetCaption(appTitle)
	frmMain.SetDoubleBuffered(true)

	//	frmMain.SetOnCloseQuery(func(sender vcl.IObject, canClose uintptr) {
	//		rtl.SetFormCanClose(canClose,
	//			vcl.MessageDlg("是否退出？", types.MtConfirmation,
	//				types.MbOK, types.MbNo) == types.MrOk)
	//	})
	initActions()
	initMainMenu()
	initTrayIcon()

	//CategoryPanelGroup
	cPlnGroup := vcl.NewCategoryPanelGroup(frmMain)
	cPlnGroup.SetParent(frmMain)

	cPln := vcl.NewCategoryPanel(frmMain)
	cPln.SetCaption("站点信息")
	cPln.SetPanelGroup(cPlnGroup)
	cPln.SetAlignWithMargins(true)
	//cPln.SetAlign(types.AlTop)

	var top int32 = 10
	lbl := vcl.NewLabel(frmMain)
	lbl.SetParent(cPln)
	lbl.SetCaption("站点：")
	lbl.SetLeft(10)
	lbl.SetTop(top)

	top += lbl.Height() + 10
	lbl = vcl.NewLabel(frmMain)
	lbl.SetParent(cPln)
	lbl.SetCaption("代码：")
	lbl.SetLeft(10)
	lbl.SetTop(top)

	cPln = vcl.NewCategoryPanel(frmMain)
	cPln.SetCaption("主菜单")
	cPln.SetPanelGroup(cPlnGroup)
	cPln.SetAlignWithMargins(true)
	//cPln.SetAlign(types.AlClient)

}

func initTrayIcon() {
	trayIcon = vcl.NewTrayIcon(frmMain)
	trayIcon.SetHint(appTitle)
	trayIcon.SetVisible(true)
	trayIcon.SetOnDblClick(func(vcl.IObject) {
		frmMain.Show()
		frmMain.Perform(WM_SYSCOMMAND, SC_RESTORE, 0)
	})
	pm := vcl.NewPopupMenu(frmMain)
	item := vcl.NewMenuItem(frmMain)
	item.SetAction(actExit)
	pm.Items().Add(item)
	trayIcon.SetPopupMenu(pm)
}
