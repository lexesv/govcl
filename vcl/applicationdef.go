package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

// CreateForm 一般不建议使用 NewForm，而优先使用CreateForm
func (a *TApplication) CreateForm() *TForm {
	return FormFromInst(Application_CreateForm(a.instance))
}
