
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func OpenDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openDialog_Create.Call(obj)
    return ret
}

func OpenDialog_Free(obj uintptr) {
    openDialog_Free.Call(obj)
}

func OpenDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenDialog_Assign(obj uintptr, Source uintptr)  {
    openDialog_Assign.Call(obj, Source )
}

func OpenDialog_ClassName(obj uintptr) string {
    ret, _, _ := openDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenDialog_ToString(obj uintptr) string {
    ret, _, _ := openDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openDialog_GetFiles.Call(obj)
    return ret
}

func OpenDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetDefaultExt(obj uintptr, value string) {
   openDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetFileName(obj uintptr, value string) {
   openDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetFilter(obj uintptr, value string) {
   openDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenDialog_SetFilterIndex(obj uintptr, value int32) {
   openDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetInitialDir(obj uintptr, value string) {
   openDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetTitle(obj uintptr, value string) {
   openDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenDialog_SetOnClose(obj uintptr, fn interface{}) {
    openDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenDialog_SetOnShow(obj uintptr, fn interface{}) {
    openDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenDialog_SetComponentIndex(obj uintptr, value int32) {
   openDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openDialog_GetOwner.Call(obj)
    return ret
}

func OpenDialog_GetName(obj uintptr) string {
    ret, _, _ := openDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetName(obj uintptr, value string) {
   openDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetTag(obj uintptr) int {
    ret, _, _ := openDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenDialog_SetTag(obj uintptr, value int) {
   openDialog_SetTag.Call(obj, uintptr(value))
}

func OpenDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

