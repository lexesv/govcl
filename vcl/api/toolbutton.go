
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ToolButton_Create(obj uintptr) uintptr {
    ret, _, _ := toolButton_Create.Call(obj)
    return ret
}

func ToolButton_Free(obj uintptr) {
    toolButton_Free.Call(obj)
}

func ToolButton_CheckMenuDropdown(obj uintptr) bool {
    ret, _, _ := toolButton_CheckMenuDropdown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_Click(obj uintptr)  {
    toolButton_Click.Call(obj)
}

func ToolButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    toolButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ToolButton_BringToFront(obj uintptr)  {
    toolButton_BringToFront.Call(obj)
}

func ToolButton_HasParent(obj uintptr) bool {
    ret, _, _ := toolButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_Hide(obj uintptr)  {
    toolButton_Hide.Call(obj)
}

func ToolButton_Invalidate(obj uintptr)  {
    toolButton_Invalidate.Call(obj)
}

func ToolButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := toolButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ToolButton_Refresh(obj uintptr)  {
    toolButton_Refresh.Call(obj)
}

func ToolButton_Repaint(obj uintptr)  {
    toolButton_Repaint.Call(obj)
}

func ToolButton_SendToBack(obj uintptr)  {
    toolButton_SendToBack.Call(obj)
}

func ToolButton_Show(obj uintptr)  {
    toolButton_Show.Call(obj)
}

func ToolButton_Update(obj uintptr)  {
    toolButton_Update.Call(obj)
}

func ToolButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := toolButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ToolButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := toolButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ToolButton_GetNamePath(obj uintptr) string {
    ret, _, _ := toolButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_Assign(obj uintptr, Source uintptr)  {
    toolButton_Assign.Call(obj, Source )
}

func ToolButton_ClassName(obj uintptr) string {
    ret, _, _ := toolButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := toolButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ToolButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := toolButton_GetHashCode.Call(obj)
    return int32(ret)
}

func ToolButton_ToString(obj uintptr) string {
    ret, _, _ := toolButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_GetIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetIndex.Call(obj)
    return int32(ret)
}

func ToolButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetAction.Call(obj)
    return ret
}

func ToolButton_SetAction(obj uintptr, value uintptr) {
   toolButton_SetAction.Call(obj, value)
}

func ToolButton_GetAllowAllUp(obj uintptr) bool {
    ret, _, _ := toolButton_GetAllowAllUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAllowAllUp(obj uintptr, value bool) {
   toolButton_SetAllowAllUp.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetAutoSize(obj uintptr) bool {
    ret, _, _ := toolButton_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAutoSize(obj uintptr, value bool) {
   toolButton_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetCaption(obj uintptr) string {
    ret, _, _ := toolButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetCaption(obj uintptr, value string) {
   toolButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetDown(obj uintptr) bool {
    ret, _, _ := toolButton_GetDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetDown(obj uintptr, value bool) {
   toolButton_SetDown.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetDropdownMenu(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetDropdownMenu.Call(obj)
    return ret
}

func ToolButton_SetDropdownMenu(obj uintptr, value uintptr) {
   toolButton_SetDropdownMenu.Call(obj, value)
}

func ToolButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := toolButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetEnabled(obj uintptr, value bool) {
   toolButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetEnableDropdown(obj uintptr) bool {
    ret, _, _ := toolButton_GetEnableDropdown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetEnableDropdown(obj uintptr, value bool) {
   toolButton_SetEnableDropdown.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetGrouped(obj uintptr) bool {
    ret, _, _ := toolButton_GetGrouped.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetGrouped(obj uintptr, value bool) {
   toolButton_SetGrouped.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetHeight.Call(obj)
    return int32(ret)
}

func ToolButton_SetHeight(obj uintptr, value int32) {
   toolButton_SetHeight.Call(obj, uintptr(value))
}

func ToolButton_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetImageIndex.Call(obj)
    return int32(ret)
}

func ToolButton_SetImageIndex(obj uintptr, value int32) {
   toolButton_SetImageIndex.Call(obj, uintptr(value))
}

func ToolButton_GetIndeterminate(obj uintptr) bool {
    ret, _, _ := toolButton_GetIndeterminate.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetIndeterminate(obj uintptr, value bool) {
   toolButton_SetIndeterminate.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetMarked(obj uintptr) bool {
    ret, _, _ := toolButton_GetMarked.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetMarked(obj uintptr, value bool) {
   toolButton_SetMarked.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := toolButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetParentShowHint(obj uintptr, value bool) {
   toolButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetPopupMenu.Call(obj)
    return ret
}

func ToolButton_SetPopupMenu(obj uintptr, value uintptr) {
   toolButton_SetPopupMenu.Call(obj, value)
}

func ToolButton_GetWrap(obj uintptr) bool {
    ret, _, _ := toolButton_GetWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetWrap(obj uintptr, value bool) {
   toolButton_SetWrap.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := toolButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetShowHint(obj uintptr, value bool) {
   toolButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetStyle(obj uintptr) TToolButtonStyle {
    ret, _, _ := toolButton_GetStyle.Call(obj)
    return TToolButtonStyle(ret)
}

func ToolButton_SetStyle(obj uintptr, value TToolButtonStyle) {
   toolButton_SetStyle.Call(obj, uintptr(value))
}

func ToolButton_GetVisible(obj uintptr) bool {
    ret, _, _ := toolButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetVisible(obj uintptr, value bool) {
   toolButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetWidth.Call(obj)
    return int32(ret)
}

func ToolButton_SetWidth(obj uintptr, value int32) {
   toolButton_SetWidth.Call(obj, uintptr(value))
}

func ToolButton_SetOnClick(obj uintptr, fn interface{}) {
    toolButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ToolButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := toolButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func ToolButton_SetAlign(obj uintptr, value TAlign) {
   toolButton_SetAlign.Call(obj, uintptr(value))
}

func ToolButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := toolButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ToolButton_SetAnchors(obj uintptr, value TAnchors) {
   toolButton_SetAnchors.Call(obj, uintptr(value))
}

func ToolButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    toolButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolButton_SetBoundsRect(obj uintptr, value TRect) {
   toolButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func ToolButton_SetClientHeight(obj uintptr, value int32) {
   toolButton_SetClientHeight.Call(obj, uintptr(value))
}

func ToolButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    toolButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func ToolButton_SetClientWidth(obj uintptr, value int32) {
   toolButton_SetClientWidth.Call(obj, uintptr(value))
}

func ToolButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ToolButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetParent.Call(obj)
    return ret
}

func ToolButton_SetParent(obj uintptr, value uintptr) {
   toolButton_SetParent.Call(obj, value)
}

func ToolButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := toolButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAlignWithMargins(obj uintptr, value bool) {
   toolButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := toolButton_GetLeft.Call(obj)
    return int32(ret)
}

func ToolButton_SetLeft(obj uintptr, value int32) {
   toolButton_SetLeft.Call(obj, uintptr(value))
}

func ToolButton_GetTop(obj uintptr) int32 {
    ret, _, _ := toolButton_GetTop.Call(obj)
    return int32(ret)
}

func ToolButton_SetTop(obj uintptr, value int32) {
   toolButton_SetTop.Call(obj, uintptr(value))
}

func ToolButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := toolButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func ToolButton_SetCursor(obj uintptr, value TCursor) {
   toolButton_SetCursor.Call(obj, uintptr(value))
}

func ToolButton_GetHint(obj uintptr) string {
    ret, _, _ := toolButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetHint(obj uintptr, value string) {
   toolButton_SetHint.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetMargins.Call(obj)
    return ret
}

func ToolButton_SetMargins(obj uintptr, value uintptr) {
   toolButton_SetMargins.Call(obj, value)
}

func ToolButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := toolButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func ToolButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ToolButton_SetComponentIndex(obj uintptr, value int32) {
   toolButton_SetComponentIndex.Call(obj, uintptr(value))
}

func ToolButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetOwner.Call(obj)
    return ret
}

func ToolButton_GetName(obj uintptr) string {
    ret, _, _ := toolButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetName(obj uintptr, value string) {
   toolButton_SetName.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetTag(obj uintptr) int {
    ret, _, _ := toolButton_GetTag.Call(obj)
    return int(ret)
}

func ToolButton_SetTag(obj uintptr, value int) {
   toolButton_SetTag.Call(obj, uintptr(value))
}

func ToolButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := toolButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

