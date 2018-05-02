package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl/rtl"
)

func main() {
	fmt.Println(rtl.OSVersion)
	fmt.Println(rtl.OSVersion.Name)
}
