package api

import (
	"reflect"
	"unsafe"

	. "gitee.com/ying32/govcl/vcl/types"
)

type TGoParam struct {
	Type  uint8
	Value uintptr
}

var (
	CallbackMap = map[uintptr]interface{}{}
)

func DBoolToGoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func GoBoolToDBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

func addEventToMap(f interface{}) uintptr {
	p := reflect.ValueOf(f).Pointer()
	CallbackMap[p] = f
	return p
}

func SetEventCallback(ptr uintptr) {
	setEventCallback.Call(ptr)
}

func DGetParam(index int, ptr uintptr) TGoParam {
	p := TGoParam{}
	dGetParam.Call(uintptr(index), ptr, uintptr(unsafe.Pointer(&p)))
	return p
}

func DStrLen(p uintptr) int {
	ret, _, _ := dStrLen.Call(p)
	return int(ret)
}

func DMove(src, dest uintptr, llen int) {
	dMove.Call(src, dest, uintptr(llen))
}

func DShowMessage(s string) {
	dShowMessage.Call(GoStrToDStr(s))
}

func DGetMainInstance() uintptr {
	ret, _, _ := dGetMainInstance.Call()
	return ret
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	ret, _, _ := dMessageDlg.Call(GoStrToDStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx))
	return int32(ret)
}

func DSetReportMemoryLeaksOnShutdown(v bool) {
	dSetReportMemoryLeaksOnShutdown.Call(GoBoolToDBool(v))
}

func DTextToShortCut(val string) TShortCut {
	ret, _, _ := dTextToShortCut.Call(GoStrToDStr(val))
	return TShortCut(ret)
}

func DShortCutToText(val TShortCut) string {
	ret, _, _ := dShortCutToText.Call(uintptr(val))
	return DStrToGoStr(ret)
}

func DSysOpen(filename string) {
	dSysOpen.Call(GoStrToDStr(filename))
}

func DExtractFilePath(filename string) string {
	r, _, _ := dExtractFilePath.Call(GoStrToDStr(filename))
	return DStrToGoStr(r)
}

func DFileExists(filename string) bool {
	r, _, _ := dFileExists.Call(GoStrToDStr(filename))
	return DBoolToGoBool(r)
}
