package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	//	"gitee.com/ying32/govcl/vcl/api"
	//	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/xui"
)

const uiXML = `<?xml encoding="utf-8" version="1.0" ?>
<Form name="mainForm" width="600" height="400" center="true" caption="这是一个测试" enabledmax="false">
  <Button name="btn1" left="10" top="10" width="75" height="25" caption="按钮1" align="alTop"></Button>
<Panel name="pnl1" align="alClient">
  <Button name="btn2" left="10" top="10" width="75" height="25" caption="消息" onclick="OnButtonClick"></Button>
</Panel>
</Form>`

type TEvents struct {
	form *vcl.TForm
}

func (e *TEvents) OnButtonClick(sender vcl.IObject) {
	fmt.Println("OnButtonClick.")
	vcl.ShowMessage("ButtonClick")
}

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	events := new(TEvents)
	m, err := xui.NewFormBytes([]byte(uiXML), events)
	if err == nil {
		events.form = m.Form
	} else {
		return
	}

	vcl.Application.Run()
}
