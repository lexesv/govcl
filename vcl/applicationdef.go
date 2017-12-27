package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/win"
)

// CreateForm 一般不建议使用 NewForm，而优先使用CreateForm
func (a *TApplication) CreateForm() *TForm {
	return FormFromInst(Application_CreateForm(a.instance))
}

// SetFormScaled 设置全局窗口的Scaled
func (a *TApplication) SetFormScaled(val bool) {
	SetGlobalFormScaled(val)
}

// SetIconResId 从资源中置图标的id
func (a *TApplication) SetIconResId(id int) {
	a.Icon().SetHandle(win.LoadIconId(win.GetSelfModuleHandle(), id))
}
