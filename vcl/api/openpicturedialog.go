
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func OpenPictureDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_Create.Call(obj)
    return ret
}

func OpenPictureDialog_Free(obj uintptr) {
    openPictureDialog_Free.Call(obj)
}

func OpenPictureDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openPictureDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openPictureDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenPictureDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openPictureDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_Assign(obj uintptr, Source uintptr)  {
    openPictureDialog_Assign.Call(obj, Source )
}

func OpenPictureDialog_ClassName(obj uintptr) string {
    ret, _, _ := openPictureDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openPictureDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_ToString(obj uintptr) string {
    ret, _, _ := openPictureDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFilter(obj uintptr, value string) {
   openPictureDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_GetFiles.Call(obj)
    return ret
}

func OpenPictureDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetDefaultExt(obj uintptr, value string) {
   openPictureDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFileName(obj uintptr, value string) {
   openPictureDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_SetFilterIndex(obj uintptr, value int32) {
   openPictureDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetInitialDir(obj uintptr, value string) {
   openPictureDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openPictureDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenPictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openPictureDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openPictureDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenPictureDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openPictureDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetTitle(obj uintptr, value string) {
   openPictureDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openPictureDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenPictureDialog_SetOnClose(obj uintptr, fn interface{}) {
    openPictureDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenPictureDialog_SetOnShow(obj uintptr, fn interface{}) {
    openPictureDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenPictureDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_SetComponentIndex(obj uintptr, value int32) {
   openPictureDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_GetOwner.Call(obj)
    return ret
}

func OpenPictureDialog_GetName(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetName(obj uintptr, value string) {
   openPictureDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetTag(obj uintptr) int {
    ret, _, _ := openPictureDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenPictureDialog_SetTag(obj uintptr, value int) {
   openPictureDialog_SetTag.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openPictureDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

