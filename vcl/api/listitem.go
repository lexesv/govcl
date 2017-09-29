
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ListItem_Create() uintptr {
    ret, _, _ := listItem_Create.Call()
    return ret
}

func ListItem_Free(obj uintptr) {
    listItem_Free.Call(obj)
}

func ListItem_Assign(obj uintptr, Source uintptr)  {
    listItem_Assign.Call(obj, Source )
}

func ListItem_CancelEdit(obj uintptr)  {
    listItem_CancelEdit.Call(obj)
}

func ListItem_Delete(obj uintptr)  {
    listItem_Delete.Call(obj)
}

func ListItem_EditCaption(obj uintptr) bool {
    ret, _, _ := listItem_EditCaption.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_MakeVisible(obj uintptr, PartialOK bool)  {
    listItem_MakeVisible.Call(obj, GoBoolToDBool(PartialOK) )
}

func ListItem_Update(obj uintptr)  {
    listItem_Update.Call(obj)
}

func ListItem_WorkArea(obj uintptr) int32 {
    ret, _, _ := listItem_WorkArea.Call(obj)
    return int32(ret)
}

func ListItem_GetNamePath(obj uintptr) string {
    ret, _, _ := listItem_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_ClassName(obj uintptr) string {
    ret, _, _ := listItem_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listItem_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListItem_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listItem_GetHashCode.Call(obj)
    return int32(ret)
}

func ListItem_ToString(obj uintptr) string {
    ret, _, _ := listItem_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_GetCaption(obj uintptr) string {
    ret, _, _ := listItem_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_SetCaption(obj uintptr, value string) {
   listItem_SetCaption.Call(obj, GoStrToDStr(value))
}

func ListItem_GetChecked(obj uintptr) bool {
    ret, _, _ := listItem_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetChecked(obj uintptr, value bool) {
   listItem_SetChecked.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetCut(obj uintptr) bool {
    ret, _, _ := listItem_GetCut.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetCut(obj uintptr, value bool) {
   listItem_SetCut.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetData(obj uintptr) uintptr {
    ret, _, _ := listItem_GetData.Call(obj)
    return ret
}

func ListItem_SetData(obj uintptr, value uintptr) {
   listItem_SetData.Call(obj, value)
}

func ListItem_GetDeleting(obj uintptr) bool {
    ret, _, _ := listItem_GetDeleting.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_GetDropTarget(obj uintptr) bool {
    ret, _, _ := listItem_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetDropTarget(obj uintptr, value bool) {
   listItem_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetFocused(obj uintptr) bool {
    ret, _, _ := listItem_GetFocused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetFocused(obj uintptr, value bool) {
   listItem_SetFocused.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetGroupID(obj uintptr) int32 {
    ret, _, _ := listItem_GetGroupID.Call(obj)
    return int32(ret)
}

func ListItem_SetGroupID(obj uintptr, value int32) {
   listItem_SetGroupID.Call(obj, uintptr(value))
}

func ListItem_GetHandle(obj uintptr) HWND {
    ret, _, _ := listItem_GetHandle.Call(obj)
    return HWND(ret)
}

func ListItem_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetImageIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetImageIndex(obj uintptr, value int32) {
   listItem_SetImageIndex.Call(obj, uintptr(value))
}

func ListItem_GetIndent(obj uintptr) int32 {
    ret, _, _ := listItem_GetIndent.Call(obj)
    return int32(ret)
}

func ListItem_SetIndent(obj uintptr, value int32) {
   listItem_SetIndent.Call(obj, uintptr(value))
}

func ListItem_GetIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetIndex.Call(obj)
    return int32(ret)
}

func ListItem_GetLeft(obj uintptr) int32 {
    ret, _, _ := listItem_GetLeft.Call(obj)
    return int32(ret)
}

func ListItem_SetLeft(obj uintptr, value int32) {
   listItem_SetLeft.Call(obj, uintptr(value))
}

func ListItem_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listItem_GetOwner.Call(obj)
    return ret
}

func ListItem_GetOverlayIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetOverlayIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetOverlayIndex(obj uintptr, value int32) {
   listItem_SetOverlayIndex.Call(obj, uintptr(value))
}

func ListItem_GetPosition(obj uintptr) TPoint {
    var ret TPoint
    listItem_GetPosition.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListItem_SetPosition(obj uintptr, value TPoint) {
   listItem_SetPosition.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListItem_GetSelected(obj uintptr) bool {
    ret, _, _ := listItem_GetSelected.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetSelected(obj uintptr, value bool) {
   listItem_SetSelected.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetStateIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetStateIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetStateIndex(obj uintptr, value int32) {
   listItem_SetStateIndex.Call(obj, uintptr(value))
}

func ListItem_GetSubItems(obj uintptr) uintptr {
    ret, _, _ := listItem_GetSubItems.Call(obj)
    return ret
}

func ListItem_SetSubItems(obj uintptr, value uintptr) {
   listItem_SetSubItems.Call(obj, value)
}

func ListItem_GetTop(obj uintptr) int32 {
    ret, _, _ := listItem_GetTop.Call(obj)
    return int32(ret)
}

func ListItem_SetTop(obj uintptr, value int32) {
   listItem_SetTop.Call(obj, uintptr(value))
}

func ListItem_GetSubItemImages(obj uintptr, Index int32) int32 {
    ret, _, _ := listItem_GetSubItemImages.Call(obj, uintptr(Index))
    return int32(ret)
}

func ListItem_SetSubItemImages(obj uintptr, Index int32, value int32) {
   listItem_SetSubItemImages.Call(obj, uintptr(Index), uintptr(value))
}

