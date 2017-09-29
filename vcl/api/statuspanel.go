
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func StatusPanel_Create() uintptr {
    ret, _, _ := statusPanel_Create.Call()
    return ret
}

func StatusPanel_Free(obj uintptr) {
    statusPanel_Free.Call(obj)
}

func StatusPanel_Assign(obj uintptr, Source uintptr)  {
    statusPanel_Assign.Call(obj, Source )
}

func StatusPanel_GetNamePath(obj uintptr) string {
    ret, _, _ := statusPanel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_ClassName(obj uintptr) string {
    ret, _, _ := statusPanel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := statusPanel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StatusPanel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetHashCode.Call(obj)
    return int32(ret)
}

func StatusPanel_ToString(obj uintptr) string {
    ret, _, _ := statusPanel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := statusPanel_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func StatusPanel_SetAlignment(obj uintptr, value TAlignment) {
   statusPanel_SetAlignment.Call(obj, uintptr(value))
}

func StatusPanel_GetStyle(obj uintptr) TStatusPanelStyle {
    ret, _, _ := statusPanel_GetStyle.Call(obj)
    return TStatusPanelStyle(ret)
}

func StatusPanel_SetStyle(obj uintptr, value TStatusPanelStyle) {
   statusPanel_SetStyle.Call(obj, uintptr(value))
}

func StatusPanel_GetText(obj uintptr) string {
    ret, _, _ := statusPanel_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_SetText(obj uintptr, value string) {
   statusPanel_SetText.Call(obj, GoStrToDStr(value))
}

func StatusPanel_GetWidth(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetWidth.Call(obj)
    return int32(ret)
}

func StatusPanel_SetWidth(obj uintptr, value int32) {
   statusPanel_SetWidth.Call(obj, uintptr(value))
}

func StatusPanel_GetIndex(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetIndex.Call(obj)
    return int32(ret)
}

func StatusPanel_SetIndex(obj uintptr, value int32) {
   statusPanel_SetIndex.Call(obj, uintptr(value))
}

