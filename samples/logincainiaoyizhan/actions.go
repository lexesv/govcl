package main

import (
	"gitee.com/ying32/govcl/vcl"
)

var (
	actExit   *vcl.TAction
	actLogin  *vcl.TAction
	actLogOff *vcl.TAction
)

func initActions() {
	actExit = vcl.NewAction(frmMain)
	actExit.SetCaption("退出(&Q)")
	actExit.SetOnExecute(func(vcl.IObject) {
		frmMain.Close()
	})

	// ----
	actLogin = vcl.NewAction(frmMain)
	actLogin.SetCaption("登录(&L)")
	actLogin.SetOnExecute(func(vcl.IObject) {
		frmLogin.ShowModal()
	})
	actLogin.SetOnUpdate(func(sender vcl.IObject) {
		vcl.ActionFromObj(sender).SetEnabled(!isLogin)
	})

	// ----
	actLogOff = vcl.NewAction(frmMain)
	actLogOff.SetCaption("注销(&O)")
	actLogOff.SetOnExecute(func(vcl.IObject) {
		isLogin = false
		session.ClearCookie()
	})
	actLogOff.SetOnUpdate(func(sender vcl.IObject) {
		vcl.ActionFromObj(sender).SetEnabled(isLogin)
	})
}
