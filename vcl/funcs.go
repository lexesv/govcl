package vcl

import (
	"reflect"

	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
)

func ShowMessage(msg string) {
	api.DShowMessage(msg)
}

// MessageDlg   api.TMsgDlgButtons
func MessageDlg(Msg string, DlgType api.TMsgDlgType, Buttons ...uint8) int32 {
	return api.DMessageDlg(Msg, DlgType, api.TMsgDlgButtons(rtl.Include(0, Buttons...)), 0)
}

func CheckPtr(value IObject) uintptr {
	if reflect.ValueOf(value).Pointer() == 0 {
		return 0
	}
	return value.Instance()
}

// exceptionProc 公共的异常处理过程
func exceptionProc() {
	if err := recover(); err != nil {
		ShowMessage(err.(error).Error())
	}
}
