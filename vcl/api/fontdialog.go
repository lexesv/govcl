
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func FontDialog_Create(obj uintptr) uintptr {
    ret, _, _ := fontDialog_Create.Call(obj)
    return ret
}

func FontDialog_Free(obj uintptr) {
    fontDialog_Free.Call(obj)
}

func FontDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := fontDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func FontDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := fontDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func FontDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := fontDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_HasParent(obj uintptr) bool {
    ret, _, _ := fontDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func FontDialog_Assign(obj uintptr, Source uintptr)  {
    fontDialog_Assign.Call(obj, Source )
}

func FontDialog_ClassName(obj uintptr) string {
    ret, _, _ := fontDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := fontDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func FontDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func FontDialog_ToString(obj uintptr) string {
    ret, _, _ := fontDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_GetFont(obj uintptr) uintptr {
    ret, _, _ := fontDialog_GetFont.Call(obj)
    return ret
}

func FontDialog_SetFont(obj uintptr, value uintptr) {
   fontDialog_SetFont.Call(obj, value)
}

func FontDialog_GetOptions(obj uintptr) TFontDialogOptions {
    ret, _, _ := fontDialog_GetOptions.Call(obj)
    return TFontDialogOptions(ret)
}

func FontDialog_SetOptions(obj uintptr, value TFontDialogOptions) {
   fontDialog_SetOptions.Call(obj, uintptr(value))
}

func FontDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := fontDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func FontDialog_SetOnClose(obj uintptr, fn interface{}) {
    fontDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func FontDialog_SetOnShow(obj uintptr, fn interface{}) {
    fontDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func FontDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func FontDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func FontDialog_SetComponentIndex(obj uintptr, value int32) {
   fontDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func FontDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := fontDialog_GetOwner.Call(obj)
    return ret
}

func FontDialog_GetName(obj uintptr) string {
    ret, _, _ := fontDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_SetName(obj uintptr, value string) {
   fontDialog_SetName.Call(obj, GoStrToDStr(value))
}

func FontDialog_GetTag(obj uintptr) int {
    ret, _, _ := fontDialog_GetTag.Call(obj)
    return int(ret)
}

func FontDialog_SetTag(obj uintptr, value int) {
   fontDialog_SetTag.Call(obj, uintptr(value))
}

func FontDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := fontDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

