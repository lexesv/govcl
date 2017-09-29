
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func SaveDialog_Create(obj uintptr) uintptr {
    ret, _, _ := saveDialog_Create.Call(obj)
    return ret
}

func SaveDialog_Free(obj uintptr) {
    saveDialog_Free.Call(obj)
}

func SaveDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := saveDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func SaveDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := saveDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SaveDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := saveDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_HasParent(obj uintptr) bool {
    ret, _, _ := saveDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveDialog_Assign(obj uintptr, Source uintptr)  {
    saveDialog_Assign.Call(obj, Source )
}

func SaveDialog_ClassName(obj uintptr) string {
    ret, _, _ := saveDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := saveDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SaveDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SaveDialog_ToString(obj uintptr) string {
    ret, _, _ := saveDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := saveDialog_GetFiles.Call(obj)
    return ret
}

func SaveDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := saveDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetDefaultExt(obj uintptr, value string) {
   saveDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFileName(obj uintptr) string {
    ret, _, _ := saveDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetFileName(obj uintptr, value string) {
   saveDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilter(obj uintptr) string {
    ret, _, _ := saveDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetFilter(obj uintptr, value string) {
   saveDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SaveDialog_SetFilterIndex(obj uintptr, value int32) {
   saveDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SaveDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := saveDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetInitialDir(obj uintptr, value string) {
   saveDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := saveDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SaveDialog_SetOptions(obj uintptr, value TOpenOptions) {
   saveDialog_SetOptions.Call(obj, uintptr(value))
}

func SaveDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := saveDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SaveDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   saveDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SaveDialog_GetTitle(obj uintptr) string {
    ret, _, _ := saveDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetTitle(obj uintptr, value string) {
   saveDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := saveDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SaveDialog_SetOnClose(obj uintptr, fn interface{}) {
    saveDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SaveDialog_SetOnShow(obj uintptr, fn interface{}) {
    saveDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SaveDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SaveDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SaveDialog_SetComponentIndex(obj uintptr, value int32) {
   saveDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SaveDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := saveDialog_GetOwner.Call(obj)
    return ret
}

func SaveDialog_GetName(obj uintptr) string {
    ret, _, _ := saveDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetName(obj uintptr, value string) {
   saveDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetTag(obj uintptr) int {
    ret, _, _ := saveDialog_GetTag.Call(obj)
    return int(ret)
}

func SaveDialog_SetTag(obj uintptr, value int) {
   saveDialog_SetTag.Call(obj, uintptr(value))
}

func SaveDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := saveDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

