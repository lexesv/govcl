package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(600)

	img := vcl.NewImage(mainForm)
	img.SetParent(mainForm)
	// 本地加载
	mem := vcl.NewMemoryStream()
	defer mem.Free()
	mem.LoadFromFile("..\\..\\imgs\\1.jpg")
	img.Picture().LoadFromStream(mem)
	// 网络图片加载
	img2 := vcl.NewImage(mainForm)
	img2.SetParent(mainForm)
	img2.SetTop(img.Height() + 10)
	img2.SetAutoSize(true)
	// 清除memory
	mem.Clear()
	// 异步加载，不知道这里会有啥，一般来说不要在非线程中访问UI组件
	go func() {
		resp, err := http.Get("http://ww2.sinaimg.cn/large/df780e95jw1egxm06uxerj20cs05hjs8.jpg")
		if err == nil {
			defer resp.Body.Close()
			bs, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				mem.Write(bs)
				mem.SetPosition(0)
				img2.Picture().LoadFromStream(mem)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}()
	vcl.Application.Run()
}
