package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl/rtl"
)

func main() {
	//fmt.Println(rtl.OSVersion)
	fmt.Println(rtl.OSVersion.Major)
	fmt.Println(rtl.OSVersion.Minor)
	fmt.Println(rtl.OSVersion.Name)
	fmt.Println(rtl.OSVersion.ToString())
	if rtl.OSVersion.CheckMajor(10) {
		fmt.Println("当前为Windows10")
	} else {
		fmt.Println("当前不是Windows10")
	}
}
