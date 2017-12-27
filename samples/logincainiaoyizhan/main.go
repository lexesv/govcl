// cnAdmin project main.go
package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
)

const (
	appTitle = "菜鸟后台"
)

func main() {
	//defer except()
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	vcl.Application.SetTitle(appTitle)
	vcl.Application.SetFormScaled(true)
	vcl.Application.SetOnMinimize(func(vcl.IObject) {
		frmMain.Hide()
	})
	//vcl.StyleManager.SetStyleFromFileName(".\\TabletLight.vsf")

	// 新建session
	session = NewCNSession()

	initfrmMain()
	initfrmLogin()
	initfrmAbout()

	vcl.Application.Run()
}

func except() {
	if err := recover(); err != nil {
		fmt.Println("Error:", err)
	}
}
