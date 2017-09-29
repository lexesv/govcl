
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func SavePictureDialog_Create(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_Create.Call(obj)
    return ret
}

func SavePictureDialog_Free(obj uintptr) {
    savePictureDialog_Free.Call(obj)
}

func SavePictureDialog_Execute(obj uintptr) bool {
    ret, _, _ := savePictureDialog_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func SavePictureDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := savePictureDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SavePictureDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_HasParent(obj uintptr) bool {
    ret, _, _ := savePictureDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SavePictureDialog_Assign(obj uintptr, Source uintptr)  {
    savePictureDialog_Assign.Call(obj, Source )
}

func SavePictureDialog_ClassName(obj uintptr) string {
    ret, _, _ := savePictureDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := savePictureDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SavePictureDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SavePictureDialog_ToString(obj uintptr) string {
    ret, _, _ := savePictureDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_GetFilter(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetFilter(obj uintptr, value string) {
   savePictureDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_GetFiles.Call(obj)
    return ret
}

func SavePictureDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetDefaultExt(obj uintptr, value string) {
   savePictureDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFileName(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetFileName(obj uintptr, value string) {
   savePictureDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SavePictureDialog_SetFilterIndex(obj uintptr, value int32) {
   savePictureDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SavePictureDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetInitialDir(obj uintptr, value string) {
   savePictureDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := savePictureDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SavePictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
   savePictureDialog_SetOptions.Call(obj, uintptr(value))
}

func SavePictureDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := savePictureDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SavePictureDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   savePictureDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SavePictureDialog_GetTitle(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetTitle(obj uintptr, value string) {
   savePictureDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := savePictureDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SavePictureDialog_SetOnClose(obj uintptr, fn interface{}) {
    savePictureDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SavePictureDialog_SetOnShow(obj uintptr, fn interface{}) {
    savePictureDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SavePictureDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SavePictureDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SavePictureDialog_SetComponentIndex(obj uintptr, value int32) {
   savePictureDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SavePictureDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_GetOwner.Call(obj)
    return ret
}

func SavePictureDialog_GetName(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetName(obj uintptr, value string) {
   savePictureDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetTag(obj uintptr) int {
    ret, _, _ := savePictureDialog_GetTag.Call(obj)
    return int(ret)
}

func SavePictureDialog_SetTag(obj uintptr, value int) {
   savePictureDialog_SetTag.Call(obj, uintptr(value))
}

func SavePictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := savePictureDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

