package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
	. "gitee.com/ying32/govcl/vcl/types"
)

func (f *TForm) ScreenCenter() {
	defer exceptionProc()
	Form_SetPosition(f.instance, PoScreenCenter)
}

func (f *TForm) EnabledMaximize(val bool) {
	defer exceptionProc()
	Form_EnabledMaximize(f.instance, val)
}

func (f *TForm) EnabledMinimize(val bool) {
	defer exceptionProc()
	Form_EnabledMinimize(f.instance, val)
}

func (f *TForm) EnabledSystemMenu(val bool) {
	defer exceptionProc()
	Form_EnabledSystemMenu(f.instance, val)
}
