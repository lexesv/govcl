package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl/win"
)

func main() {

	// Exe运行在Administrator权限下检测
	fmt.Println("IsAdministrator:", win.IsAdministrator())

	// 注：只有当exe为Win32并在64位系统上运行时才返回true, 否则都会返回false
	fmt.Println("IsWow64:", win.IsWow64())

	// 在资源管理器中定位指定文件名
	// win.OpenInExplorer("F:\\Golang\\bin\\libvclx64.dll")
	var s string
	fmt.Scan(&s)

}
