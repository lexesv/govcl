package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
	. "gitee.com/ying32/govcl/vcl/types"
)

func (f *TForm) ScreenCenter() {
	Form_SetPosition(f.instance, PoScreenCenter)
}

func (f *TForm) EnabledMaximize(val bool) {
	Form_EnabledMaximize(f.instance, val)
}

func (f *TForm) EnabledMinimize(val bool) {
	Form_EnabledMinimize(f.instance, val)
}

func (f *TForm) EnabledSystemMenu(val bool) {
	Form_EnabledSystemMenu(f.instance, val)
}
