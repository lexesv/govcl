
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func MainMenu_Create(obj uintptr) uintptr {
    ret, _, _ := mainMenu_Create.Call(obj)
    return ret
}

func MainMenu_Free(obj uintptr) {
    mainMenu_Free.Call(obj)
}

func MainMenu_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := mainMenu_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MainMenu_GetNamePath(obj uintptr) string {
    ret, _, _ := mainMenu_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_HasParent(obj uintptr) bool {
    ret, _, _ := mainMenu_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MainMenu_Assign(obj uintptr, Source uintptr)  {
    mainMenu_Assign.Call(obj, Source )
}

func MainMenu_ClassName(obj uintptr) string {
    ret, _, _ := mainMenu_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := mainMenu_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MainMenu_GetHashCode(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetHashCode.Call(obj)
    return int32(ret)
}

func MainMenu_ToString(obj uintptr) string {
    ret, _, _ := mainMenu_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_GetAutoHotkeys(obj uintptr) TMenuAutoFlag {
    ret, _, _ := mainMenu_GetAutoHotkeys.Call(obj)
    return TMenuAutoFlag(ret)
}

func MainMenu_SetAutoHotkeys(obj uintptr, value TMenuAutoFlag) {
   mainMenu_SetAutoHotkeys.Call(obj, uintptr(value))
}

func MainMenu_GetImages(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetImages.Call(obj)
    return ret
}

func MainMenu_SetImages(obj uintptr, value uintptr) {
   mainMenu_SetImages.Call(obj, value)
}

func MainMenu_SetOnChange(obj uintptr, fn interface{}) {
    mainMenu_SetOnChange.Call(obj, addEventToMap(fn))
}

func MainMenu_GetHandle(obj uintptr) HMENU {
    ret, _, _ := mainMenu_GetHandle.Call(obj)
    return HMENU(ret)
}

func MainMenu_GetWindowHandle(obj uintptr) HWND {
    ret, _, _ := mainMenu_GetWindowHandle.Call(obj)
    return HWND(ret)
}

func MainMenu_SetWindowHandle(obj uintptr, value HWND) {
   mainMenu_SetWindowHandle.Call(obj, uintptr(value))
}

func MainMenu_GetItems(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetItems.Call(obj)
    return ret
}

func MainMenu_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetComponentCount.Call(obj)
    return int32(ret)
}

func MainMenu_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MainMenu_SetComponentIndex(obj uintptr, value int32) {
   mainMenu_SetComponentIndex.Call(obj, uintptr(value))
}

func MainMenu_GetOwner(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetOwner.Call(obj)
    return ret
}

func MainMenu_GetName(obj uintptr) string {
    ret, _, _ := mainMenu_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_SetName(obj uintptr, value string) {
   mainMenu_SetName.Call(obj, GoStrToDStr(value))
}

func MainMenu_GetTag(obj uintptr) int {
    ret, _, _ := mainMenu_GetTag.Call(obj)
    return int(ret)
}

func MainMenu_SetTag(obj uintptr, value int) {
   mainMenu_SetTag.Call(obj, uintptr(value))
}

func MainMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := mainMenu_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

