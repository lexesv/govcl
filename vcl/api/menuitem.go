
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func MenuItem_Create(obj uintptr) uintptr {
    ret, _, _ := menuItem_Create.Call(obj)
    return ret
}

func MenuItem_Free(obj uintptr) {
    menuItem_Free.Call(obj)
}

func MenuItem_Insert(obj uintptr, Index int32, Item uintptr)  {
    menuItem_Insert.Call(obj, uintptr(Index) , Item )
}

func MenuItem_Delete(obj uintptr, Index int32)  {
    menuItem_Delete.Call(obj, uintptr(Index) )
}

func MenuItem_Clear(obj uintptr)  {
    menuItem_Clear.Call(obj)
}

func MenuItem_Click(obj uintptr)  {
    menuItem_Click.Call(obj)
}

func MenuItem_IndexOf(obj uintptr, Item uintptr) int32 {
    ret, _, _ := menuItem_IndexOf.Call(obj, Item )
    return int32(ret)
}

func MenuItem_HasParent(obj uintptr) bool {
    ret, _, _ := menuItem_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_Add(obj uintptr, Item uintptr)  {
    menuItem_Add.Call(obj, Item )
}

func MenuItem_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := menuItem_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MenuItem_GetNamePath(obj uintptr) string {
    ret, _, _ := menuItem_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_Assign(obj uintptr, Source uintptr)  {
    menuItem_Assign.Call(obj, Source )
}

func MenuItem_ClassName(obj uintptr) string {
    ret, _, _ := menuItem_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := menuItem_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MenuItem_GetHashCode(obj uintptr) int32 {
    ret, _, _ := menuItem_GetHashCode.Call(obj)
    return int32(ret)
}

func MenuItem_ToString(obj uintptr) string {
    ret, _, _ := menuItem_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_GetHandle(obj uintptr) HMENU {
    ret, _, _ := menuItem_GetHandle.Call(obj)
    return HMENU(ret)
}

func MenuItem_GetParent(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetParent.Call(obj)
    return ret
}

func MenuItem_GetAction(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetAction.Call(obj)
    return ret
}

func MenuItem_SetAction(obj uintptr, value uintptr) {
   menuItem_SetAction.Call(obj, value)
}

func MenuItem_GetAutoHotkeys(obj uintptr) TMenuItemAutoFlag {
    ret, _, _ := menuItem_GetAutoHotkeys.Call(obj)
    return TMenuItemAutoFlag(ret)
}

func MenuItem_SetAutoHotkeys(obj uintptr, value TMenuItemAutoFlag) {
   menuItem_SetAutoHotkeys.Call(obj, uintptr(value))
}

func MenuItem_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetBitmap.Call(obj)
    return ret
}

func MenuItem_SetBitmap(obj uintptr, value uintptr) {
   menuItem_SetBitmap.Call(obj, value)
}

func MenuItem_GetCaption(obj uintptr) string {
    ret, _, _ := menuItem_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetCaption(obj uintptr, value string) {
   menuItem_SetCaption.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetChecked(obj uintptr) bool {
    ret, _, _ := menuItem_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetChecked(obj uintptr, value bool) {
   menuItem_SetChecked.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetDefault(obj uintptr) bool {
    ret, _, _ := menuItem_GetDefault.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetDefault(obj uintptr, value bool) {
   menuItem_SetDefault.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetEnabled(obj uintptr) bool {
    ret, _, _ := menuItem_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetEnabled(obj uintptr, value bool) {
   menuItem_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetGroupIndex(obj uintptr) uint8 {
    ret, _, _ := menuItem_GetGroupIndex.Call(obj)
    return uint8(ret)
}

func MenuItem_SetGroupIndex(obj uintptr, value uint8) {
   menuItem_SetGroupIndex.Call(obj, uintptr(value))
}

func MenuItem_GetHint(obj uintptr) string {
    ret, _, _ := menuItem_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetHint(obj uintptr, value string) {
   menuItem_SetHint.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := menuItem_GetImageIndex.Call(obj)
    return int32(ret)
}

func MenuItem_SetImageIndex(obj uintptr, value int32) {
   menuItem_SetImageIndex.Call(obj, uintptr(value))
}

func MenuItem_GetShortCut(obj uintptr) TShortCut {
    ret, _, _ := menuItem_GetShortCut.Call(obj)
    return TShortCut(ret)
}

func MenuItem_SetShortCut(obj uintptr, value TShortCut) {
   menuItem_SetShortCut.Call(obj, uintptr(value))
}

func MenuItem_GetVisible(obj uintptr) bool {
    ret, _, _ := menuItem_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetVisible(obj uintptr, value bool) {
   menuItem_SetVisible.Call(obj, GoBoolToDBool(value))
}

func MenuItem_SetOnClick(obj uintptr, fn interface{}) {
    menuItem_SetOnClick.Call(obj, addEventToMap(fn))
}

func MenuItem_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := menuItem_GetComponentCount.Call(obj)
    return int32(ret)
}

func MenuItem_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := menuItem_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MenuItem_SetComponentIndex(obj uintptr, value int32) {
   menuItem_SetComponentIndex.Call(obj, uintptr(value))
}

func MenuItem_GetOwner(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetOwner.Call(obj)
    return ret
}

func MenuItem_GetName(obj uintptr) string {
    ret, _, _ := menuItem_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetName(obj uintptr, value string) {
   menuItem_SetName.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetTag(obj uintptr) int {
    ret, _, _ := menuItem_GetTag.Call(obj)
    return int(ret)
}

func MenuItem_SetTag(obj uintptr, value int) {
   menuItem_SetTag.Call(obj, uintptr(value))
}

func MenuItem_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := menuItem_GetItems.Call(obj, uintptr(Index))
    return ret
}

func MenuItem_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := menuItem_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

