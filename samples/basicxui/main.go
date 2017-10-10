package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	//	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/xui"
)

const uiXML = `<?xml encoding="utf-8" version="1.0" ?>
<Form name="mainForm" width="600" height="400" center="true" caption="这是一个测试" enabledmax="false">
	<MainMenu>
		<MenuItem name="mn1" caption="文件">
			<MenuItem name="mn3" caption="新建" />
			<MenuItem caption="-" />
			<MenuItem name="mn4" caption="退出" onclick="OnMenuExit" />
		</MenuItem>
		<MenuItem name="mn2" caption="关于" />
	</MainMenu>
	<Button name="btn1" left="10" top="10" caption="按钮" align="alTop"></Button>
	<Panel name="pnl1" align="alClient">
		<Button name="btn2" left="10" top="10" caption="消息" onclick="OnButtonClick"></Button>
	</Panel>
</Form>`

type TEvents struct {
	form *vcl.TForm
}

func (e *TEvents) OnButtonClick(sender vcl.IObject) {
	fmt.Println("OnButtonClick.")
	vcl.ShowMessage("ButtonClick")
}

func (e *TEvents) OnMenuExit(sender vcl.IObject) {
	e.form.Close()
}

func main() {
	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	events := new(TEvents)
	m, err := xui.NewFormBytes([]byte(uiXML), events)
	if err == nil {
		events.form = m.Form
	} else {
		return
	}

	vcl.Application.Run()
}
