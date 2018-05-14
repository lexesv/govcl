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

func (f *TForm) ScaleForPPI(val int32) {
	Form_ScaleForPPI(f.instance, val)
}

func (f *TForm) ScaleControlsForDpi(val int32) {
	Form_ScaleControlsForDpi(f.instance, val)
}

func (f *TForm) SetAllowDropFiles(val bool) {
	Form_SetAllowDropFiles(f.instance, val)
}

func (f *TForm) AllowDropFiles() bool {
	return Form_GetAllowDropFiles(f.instance)
}

func (f *TForm) SetOnDropFiles(fn TDropFilesEvent) {
	Form_SetOnDropFiles(f.instance, fn)
}

func (f *TForm) SetOnDestroy(fn TNotifyEvent) {
	Form_SetOnDestroy(f.instance, fn)
}
