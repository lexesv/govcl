package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func (a *TApplication) CreateForm() *TForm {
	defer exceptionProc()
	return FormFromInst(Application_CreateForm(a.instance))
}
