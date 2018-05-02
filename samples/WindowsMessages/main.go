package main

import (
	"syscall"

	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/win"
)

var (
	oldWndPrc uintptr
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Windows Messages")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	newWndProc := syscall.NewCallback(WndProc)

	oldWndPrc = win.SetWindowLongPtrW(mainForm.Handle(), win.GWL_WNDPROC, newWndProc)
	fmt.Println("newWndProc:", newWndProc)
	fmt.Println("oldWndPro:", oldWndPrc)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("按钮1")
	btn.SetLeft(50)
	btn.SetTop(50)

	vcl.Application.Run()

	// 完成后要恢复的, 先不管了
	win.SetWindowLongPtrW(mainForm.Handle(), win.GWL_WNDPROC, oldWndPrc)

}

func WndProc(hWnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case win.WM_SYSCOMMAND:
		switch wParam {
		case win.SC_MAXIMIZE:
			fmt.Println("最大化")
		case win.SC_MINIMIZE:
			fmt.Println("最小化")
		case win.SC_RESTORE:
			fmt.Println("恢复")
		case win.SC_CLOSE:
			fmt.Println("关闭")
		}
	}
	return win.CallWindowProcW(oldWndPrc, types.HWND(hWnd), message, wParam, lParam)
}
