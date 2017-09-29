
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func PrintDialog_Create(obj uintptr) uintptr {
    ret, _, _ := printDialog_Create.Call(obj)
    return ret
}

func PrintDialog_Free(obj uintptr) {
    printDialog_Free.Call(obj)
}

func PrintDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := printDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func PrintDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := printDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PrintDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := printDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_HasParent(obj uintptr) bool {
    ret, _, _ := printDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_Assign(obj uintptr, Source uintptr)  {
    printDialog_Assign.Call(obj, Source )
}

func PrintDialog_ClassName(obj uintptr) string {
    ret, _, _ := printDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := printDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PrintDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := printDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func PrintDialog_ToString(obj uintptr) string {
    ret, _, _ := printDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_GetCollate(obj uintptr) bool {
    ret, _, _ := printDialog_GetCollate.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_SetCollate(obj uintptr, value bool) {
   printDialog_SetCollate.Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetCopies(obj uintptr) int32 {
    ret, _, _ := printDialog_GetCopies.Call(obj)
    return int32(ret)
}

func PrintDialog_SetCopies(obj uintptr, value int32) {
   printDialog_SetCopies.Call(obj, uintptr(value))
}

func PrintDialog_GetFromPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetFromPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetFromPage(obj uintptr, value int32) {
   printDialog_SetFromPage.Call(obj, uintptr(value))
}

func PrintDialog_GetMinPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetMinPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetMinPage(obj uintptr, value int32) {
   printDialog_SetMinPage.Call(obj, uintptr(value))
}

func PrintDialog_GetMaxPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetMaxPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetMaxPage(obj uintptr, value int32) {
   printDialog_SetMaxPage.Call(obj, uintptr(value))
}

func PrintDialog_GetOptions(obj uintptr) TPrintDialogOptions {
    ret, _, _ := printDialog_GetOptions.Call(obj)
    return TPrintDialogOptions(ret)
}

func PrintDialog_SetOptions(obj uintptr, value TPrintDialogOptions) {
   printDialog_SetOptions.Call(obj, uintptr(value))
}

func PrintDialog_GetPrintToFile(obj uintptr) bool {
    ret, _, _ := printDialog_GetPrintToFile.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_SetPrintToFile(obj uintptr, value bool) {
   printDialog_SetPrintToFile.Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetPrintRange(obj uintptr) TPrintRange {
    ret, _, _ := printDialog_GetPrintRange.Call(obj)
    return TPrintRange(ret)
}

func PrintDialog_SetPrintRange(obj uintptr, value TPrintRange) {
   printDialog_SetPrintRange.Call(obj, uintptr(value))
}

func PrintDialog_GetToPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetToPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetToPage(obj uintptr, value int32) {
   printDialog_SetToPage.Call(obj, uintptr(value))
}

func PrintDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := printDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func PrintDialog_SetOnClose(obj uintptr, fn interface{}) {
    printDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func PrintDialog_SetOnShow(obj uintptr, fn interface{}) {
    printDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func PrintDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := printDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func PrintDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := printDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PrintDialog_SetComponentIndex(obj uintptr, value int32) {
   printDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func PrintDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := printDialog_GetOwner.Call(obj)
    return ret
}

func PrintDialog_GetName(obj uintptr) string {
    ret, _, _ := printDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_SetName(obj uintptr, value string) {
   printDialog_SetName.Call(obj, GoStrToDStr(value))
}

func PrintDialog_GetTag(obj uintptr) int {
    ret, _, _ := printDialog_GetTag.Call(obj)
    return int(ret)
}

func PrintDialog_SetTag(obj uintptr, value int) {
   printDialog_SetTag.Call(obj, uintptr(value))
}

func PrintDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := printDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

