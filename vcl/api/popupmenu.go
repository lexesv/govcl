
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func PopupMenu_Create(obj uintptr) uintptr {
    ret, _, _ := popupMenu_Create.Call(obj)
    return ret
}

func PopupMenu_Free(obj uintptr) {
    popupMenu_Free.Call(obj)
}

func PopupMenu_CloseMenu(obj uintptr)  {
    popupMenu_CloseMenu.Call(obj)
}

func PopupMenu_Popup(obj uintptr, X int32, Y int32)  {
    popupMenu_Popup.Call(obj, uintptr(X) , uintptr(Y) )
}

func PopupMenu_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := popupMenu_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PopupMenu_GetNamePath(obj uintptr) string {
    ret, _, _ := popupMenu_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_HasParent(obj uintptr) bool {
    ret, _, _ := popupMenu_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PopupMenu_Assign(obj uintptr, Source uintptr)  {
    popupMenu_Assign.Call(obj, Source )
}

func PopupMenu_ClassName(obj uintptr) string {
    ret, _, _ := popupMenu_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := popupMenu_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PopupMenu_GetHashCode(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetHashCode.Call(obj)
    return int32(ret)
}

func PopupMenu_ToString(obj uintptr) string {
    ret, _, _ := popupMenu_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_GetAlignment(obj uintptr) TPopupAlignment {
    ret, _, _ := popupMenu_GetAlignment.Call(obj)
    return TPopupAlignment(ret)
}

func PopupMenu_SetAlignment(obj uintptr, value TPopupAlignment) {
   popupMenu_SetAlignment.Call(obj, uintptr(value))
}

func PopupMenu_GetAutoHotkeys(obj uintptr) TMenuAutoFlag {
    ret, _, _ := popupMenu_GetAutoHotkeys.Call(obj)
    return TMenuAutoFlag(ret)
}

func PopupMenu_SetAutoHotkeys(obj uintptr, value TMenuAutoFlag) {
   popupMenu_SetAutoHotkeys.Call(obj, uintptr(value))
}

func PopupMenu_GetImages(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetImages.Call(obj)
    return ret
}

func PopupMenu_SetImages(obj uintptr, value uintptr) {
   popupMenu_SetImages.Call(obj, value)
}

func PopupMenu_SetOnChange(obj uintptr, fn interface{}) {
    popupMenu_SetOnChange.Call(obj, addEventToMap(fn))
}

func PopupMenu_SetOnPopup(obj uintptr, fn interface{}) {
    popupMenu_SetOnPopup.Call(obj, addEventToMap(fn))
}

func PopupMenu_GetHandle(obj uintptr) HMENU {
    ret, _, _ := popupMenu_GetHandle.Call(obj)
    return HMENU(ret)
}

func PopupMenu_GetWindowHandle(obj uintptr) HWND {
    ret, _, _ := popupMenu_GetWindowHandle.Call(obj)
    return HWND(ret)
}

func PopupMenu_SetWindowHandle(obj uintptr, value HWND) {
   popupMenu_SetWindowHandle.Call(obj, uintptr(value))
}

func PopupMenu_GetItems(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetItems.Call(obj)
    return ret
}

func PopupMenu_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetComponentCount.Call(obj)
    return int32(ret)
}

func PopupMenu_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PopupMenu_SetComponentIndex(obj uintptr, value int32) {
   popupMenu_SetComponentIndex.Call(obj, uintptr(value))
}

func PopupMenu_GetOwner(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetOwner.Call(obj)
    return ret
}

func PopupMenu_GetName(obj uintptr) string {
    ret, _, _ := popupMenu_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_SetName(obj uintptr, value string) {
   popupMenu_SetName.Call(obj, GoStrToDStr(value))
}

func PopupMenu_GetTag(obj uintptr) int {
    ret, _, _ := popupMenu_GetTag.Call(obj)
    return int(ret)
}

func PopupMenu_SetTag(obj uintptr, value int) {
   popupMenu_SetTag.Call(obj, uintptr(value))
}

func PopupMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := popupMenu_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

