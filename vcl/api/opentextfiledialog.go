
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func OpenTextFileDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_Create.Call(obj)
    return ret
}

func OpenTextFileDialog_Free(obj uintptr) {
    openTextFileDialog_Free.Call(obj)
}

func OpenTextFileDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openTextFileDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openTextFileDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenTextFileDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openTextFileDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_Assign(obj uintptr, Source uintptr)  {
    openTextFileDialog_Assign.Call(obj, Source )
}

func OpenTextFileDialog_ClassName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openTextFileDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_ToString(obj uintptr) string {
    ret, _, _ := openTextFileDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_GetFiles.Call(obj)
    return ret
}

func OpenTextFileDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetDefaultExt(obj uintptr, value string) {
   openTextFileDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetFileName(obj uintptr, value string) {
   openTextFileDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetFilter(obj uintptr, value string) {
   openTextFileDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_SetFilterIndex(obj uintptr, value int32) {
   openTextFileDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetInitialDir(obj uintptr, value string) {
   openTextFileDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openTextFileDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenTextFileDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openTextFileDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openTextFileDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenTextFileDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openTextFileDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetTitle(obj uintptr, value string) {
   openTextFileDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openTextFileDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenTextFileDialog_SetOnClose(obj uintptr, fn interface{}) {
    openTextFileDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenTextFileDialog_SetOnShow(obj uintptr, fn interface{}) {
    openTextFileDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenTextFileDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_SetComponentIndex(obj uintptr, value int32) {
   openTextFileDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_GetOwner.Call(obj)
    return ret
}

func OpenTextFileDialog_GetName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetName(obj uintptr, value string) {
   openTextFileDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetTag(obj uintptr) int {
    ret, _, _ := openTextFileDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenTextFileDialog_SetTag(obj uintptr, value int) {
   openTextFileDialog_SetTag.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openTextFileDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

