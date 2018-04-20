package main

import (
	"gitee.com/ying32/govcl/vcl"

	"fmt"

	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("DropFiles测试")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)
	mainForm.SetAllowDropFiles(true)
	mainForm.SetOnDropFiles(func(sender *vcl.TObject, aFileNames []string) {
		fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
		for i, s := range aFileNames {
			fmt.Println("index:", i, ", filename:", s)
		}
	})
	vcl.Application.Run()
}
