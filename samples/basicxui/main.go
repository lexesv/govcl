package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/xui"
)

const uiXML = `<?xml encoding="utf-8" version="1.0" ?>
<Form width="800" height="600" center="true" caption="这是一个测试" enabledmax="true">
	<MainMenu>
		<MenuItem name="mn1" caption="文件">
			<MenuItem name="mn3" caption="新建" />
			<MenuItem caption="-" />
			<MenuItem name="mn4" caption="退出" onclick="OnMenuExit" />
		</MenuItem>
		<MenuItem name="mn2" caption="关于" />
	</MainMenu>
	<TrayIcon name="Tray1" hint="fdsfsdf" visible="true">
	
	</TrayIcon> 
	<Button name="Btn1" caption="按钮" align="alTop" hint="提示" showhint="true" />
	<Panel name="Pnl1" align="alClient">
		<PageControl name="PgcMain" activeindex="0" align="alClient">
			<TabSheet caption="第一页">
			   <Button name="Btn2" left="10" top="10" caption="消息2" onclick="OnButtonClick" />
			   <Combobox name="Cbb1" itemindex="0" top="70" left="10">
					<TextItem text="第1项" />
					<TextItem text="第2项" />
					<TextItem text="第3项" />
			   </Combobox>
			   <Memo name="Mmo1" top="100" left="10">
					<TextItem text="第1项" />
					<TextItem text="第2项" />
			   </Memo>
			   <Memo name="Mmo2" top="220" left="10" text="这是文本" />
			</TabSheet>
			<TabSheet caption="第二页"> 
			    <Button name="Btn4" left="100" top="50" caption="消息4" onclick="OnButtonClick" />
			    <Button name="Btn3" caption="消息3" onclick="OnButtonClick" />
			</TabSheet>
		</PageControl>
	</Panel>
</Form>`

type TEvents struct {
	xui.TEvents
	Btn1    *vcl.TButton
	Btn2    *vcl.TButton
	Pnl1    *vcl.TPanel
	Btn3    *vcl.TButton
	PgcMain *vcl.TPageControl
}

func (e *TEvents) OnButtonClick(sender vcl.IObject) {
	fmt.Println("OnButtonClick.")
	e.Btn1.SetCaption("changed")
	vcl.ShowMessage("ButtonClick")
}

func (e *TEvents) OnMenuExit(sender vcl.IObject) {
	fmt.Println("e.Form:", e.Form)
	e.Form.Close()
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
	if err != nil {
		return
	}
	fmt.Println("m:", m)
	fmt.Println("m.btn1", events.Btn3.Left(), events.Btn3.Top(),
		events.Btn3.Height(), events.Btn3.Width(),
		events.Btn3.Visible(), events.Btn3.Parent().ClassName())

	vcl.Application.Run()
}
