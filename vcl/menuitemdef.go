package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func (m *TMenuItem) SetShortCutFromString(s string) {
	defer exceptionProc()
	MenuItem_SetShortCut(m.instance, DTextToShortCut(s))
}

func (m *TMenuItem) ShortCutFromString() string {
	defer exceptionProc()
	return DShortCutToText(MenuItem_GetShortCut(m.instance))
}
