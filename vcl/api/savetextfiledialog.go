
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func SaveTextFileDialog_Create(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_Create.Call(obj)
    return ret
}

func SaveTextFileDialog_Free(obj uintptr) {
    saveTextFileDialog_Free.Call(obj)
}

func SaveTextFileDialog_Execute(obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := saveTextFileDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SaveTextFileDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_HasParent(obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_Assign(obj uintptr, Source uintptr)  {
    saveTextFileDialog_Assign.Call(obj, Source )
}

func SaveTextFileDialog_ClassName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_ToString(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_GetFiles.Call(obj)
    return ret
}

func SaveTextFileDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetDefaultExt(obj uintptr, value string) {
   saveTextFileDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFileName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetFileName(obj uintptr, value string) {
   saveTextFileDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFilter(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetFilter(obj uintptr, value string) {
   saveTextFileDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_SetFilterIndex(obj uintptr, value int32) {
   saveTextFileDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetInitialDir(obj uintptr, value string) {
   saveTextFileDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := saveTextFileDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SaveTextFileDialog_SetOptions(obj uintptr, value TOpenOptions) {
   saveTextFileDialog_SetOptions.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := saveTextFileDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SaveTextFileDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   saveTextFileDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetTitle(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetTitle(obj uintptr, value string) {
   saveTextFileDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := saveTextFileDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SaveTextFileDialog_SetOnClose(obj uintptr, fn interface{}) {
    saveTextFileDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SaveTextFileDialog_SetOnShow(obj uintptr, fn interface{}) {
    saveTextFileDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SaveTextFileDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_SetComponentIndex(obj uintptr, value int32) {
   saveTextFileDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_GetOwner.Call(obj)
    return ret
}

func SaveTextFileDialog_GetName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetName(obj uintptr, value string) {
   saveTextFileDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetTag(obj uintptr) int {
    ret, _, _ := saveTextFileDialog_GetTag.Call(obj)
    return int(ret)
}

func SaveTextFileDialog_SetTag(obj uintptr, value int) {
   saveTextFileDialog_SetTag.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := saveTextFileDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

