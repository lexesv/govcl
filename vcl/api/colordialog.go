
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func ColorDialog_Create(obj uintptr) uintptr {
    ret, _, _ := colorDialog_Create.Call(obj)
    return ret
}

func ColorDialog_Free(obj uintptr) {
    colorDialog_Free.Call(obj)
}

func ColorDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := colorDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func ColorDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := colorDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_HasParent(obj uintptr) bool {
    ret, _, _ := colorDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorDialog_Assign(obj uintptr, Source uintptr)  {
    colorDialog_Assign.Call(obj, Source )
}

func ColorDialog_ClassName(obj uintptr) string {
    ret, _, _ := colorDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorDialog_ToString(obj uintptr) string {
    ret, _, _ := colorDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_GetColor(obj uintptr) TColor {
    ret, _, _ := colorDialog_GetColor.Call(obj)
    return TColor(ret)
}

func ColorDialog_SetColor(obj uintptr, value TColor) {
   colorDialog_SetColor.Call(obj, uintptr(value))
}

func ColorDialog_GetOptions(obj uintptr) TColorDialogOptions {
    ret, _, _ := colorDialog_GetOptions.Call(obj)
    return TColorDialogOptions(ret)
}

func ColorDialog_SetOptions(obj uintptr, value TColorDialogOptions) {
   colorDialog_SetOptions.Call(obj, uintptr(value))
}

func ColorDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorDialog_SetOnClose(obj uintptr, fn interface{}) {
    colorDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func ColorDialog_SetOnShow(obj uintptr, fn interface{}) {
    colorDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func ColorDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorDialog_SetComponentIndex(obj uintptr, value int32) {
   colorDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorDialog_GetOwner.Call(obj)
    return ret
}

func ColorDialog_GetName(obj uintptr) string {
    ret, _, _ := colorDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_SetName(obj uintptr, value string) {
   colorDialog_SetName.Call(obj, GoStrToDStr(value))
}

func ColorDialog_GetTag(obj uintptr) int {
    ret, _, _ := colorDialog_GetTag.Call(obj)
    return int(ret)
}

func ColorDialog_SetTag(obj uintptr, value int) {
   colorDialog_SetTag.Call(obj, uintptr(value))
}

func ColorDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

