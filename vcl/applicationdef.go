package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func (a *TApplication) CreateForm() *TForm {
	return FormFromInst(Application_CreateForm(a.instance))
}
