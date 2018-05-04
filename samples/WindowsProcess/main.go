package main

import (
	"unsafe"

	"path/filepath"

	"syscall"

	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/win"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Windows Process")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
	// 双缓冲
	mainForm.SetDoubleBuffered(true)

	lv1 := vcl.NewListView(mainForm)
	lv1.SetParent(mainForm)
	lv1.SetAlign(types.AlClient)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetViewStyle(types.VsReport)
	lv1.SetGridLines(true)
	//lv1.SetColumnClick(false)

	col := lv1.Columns().Add()
	col.SetCaption("进程名")
	col.SetAutoSize(true)

	col = lv1.Columns().Add()
	col.SetCaption("PID")
	col.SetAutoSize(true)

	lv1.Clear()
	fullListView(lv1)

	vcl.Application.Run()
}

func fullListView(lv *vcl.TListView) {
	var fProcessEntry32 win.TProcessEntry32W
	fProcessEntry32.DwSize = uint32(unsafe.Sizeof(fProcessEntry32))

	fSnapShotHandle := win.CreateToolhelp32SnapShot(win.TH32CS_SNAPPROCESS, 0)
	continueLoop := win.Process32FirstW(fSnapShotHandle, &fProcessEntry32)
	lv.Items().BeginUpdate()
	defer lv.Items().EndUpdate()
	for continueLoop {
		item := lv.Items().Add()
		item.SetCaption(filepath.Base(syscall.UTF16ToString(fProcessEntry32.SzExeFile[:])))
		item.SubItems().Add(fmt.Sprintf("%.4X", fProcessEntry32.Th32ProcessID))
		continueLoop = win.Process32NextW(fSnapShotHandle, &fProcessEntry32)
	}
	if fSnapShotHandle != 0 {
		win.CloseHandle(fSnapShotHandle)
	}
}
