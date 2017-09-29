
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func HotKey_Create(obj uintptr) uintptr {
    ret, _, _ := hotKey_Create.Call(obj)
    return ret
}

func HotKey_Free(obj uintptr) {
    hotKey_Free.Call(obj)
}

func HotKey_CanFocus(obj uintptr) bool {
    ret, _, _ := hotKey_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_FlipChildren(obj uintptr, AllLevels bool)  {
    hotKey_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func HotKey_Focused(obj uintptr) bool {
    ret, _, _ := hotKey_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_HandleAllocated(obj uintptr) bool {
    ret, _, _ := hotKey_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_Invalidate(obj uintptr)  {
    hotKey_Invalidate.Call(obj)
}

func HotKey_Realign(obj uintptr)  {
    hotKey_Realign.Call(obj)
}

func HotKey_Repaint(obj uintptr)  {
    hotKey_Repaint.Call(obj)
}

func HotKey_ScaleBy(obj uintptr, M int32, D int32)  {
    hotKey_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func HotKey_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    hotKey_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func HotKey_SetFocus(obj uintptr)  {
    hotKey_SetFocus.Call(obj)
}

func HotKey_Update(obj uintptr)  {
    hotKey_Update.Call(obj)
}

func HotKey_BringToFront(obj uintptr)  {
    hotKey_BringToFront.Call(obj)
}

func HotKey_HasParent(obj uintptr) bool {
    ret, _, _ := hotKey_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_Hide(obj uintptr)  {
    hotKey_Hide.Call(obj)
}

func HotKey_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := hotKey_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func HotKey_Refresh(obj uintptr)  {
    hotKey_Refresh.Call(obj)
}

func HotKey_SendToBack(obj uintptr)  {
    hotKey_SendToBack.Call(obj)
}

func HotKey_Show(obj uintptr)  {
    hotKey_Show.Call(obj)
}

func HotKey_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := hotKey_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func HotKey_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := hotKey_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func HotKey_GetNamePath(obj uintptr) string {
    ret, _, _ := hotKey_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_Assign(obj uintptr, Source uintptr)  {
    hotKey_Assign.Call(obj, Source )
}

func HotKey_ClassName(obj uintptr) string {
    ret, _, _ := hotKey_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := hotKey_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func HotKey_GetHashCode(obj uintptr) int32 {
    ret, _, _ := hotKey_GetHashCode.Call(obj)
    return int32(ret)
}

func HotKey_ToString(obj uintptr) string {
    ret, _, _ := hotKey_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := hotKey_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func HotKey_SetAnchors(obj uintptr, value TAnchors) {
   hotKey_SetAnchors.Call(obj, uintptr(value))
}

func HotKey_GetAutoSize(obj uintptr) bool {
    ret, _, _ := hotKey_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetAutoSize(obj uintptr, value bool) {
   hotKey_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetEnabled(obj uintptr) bool {
    ret, _, _ := hotKey_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetEnabled(obj uintptr, value bool) {
   hotKey_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetHint(obj uintptr) string {
    ret, _, _ := hotKey_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_SetHint(obj uintptr, value string) {
   hotKey_SetHint.Call(obj, GoStrToDStr(value))
}

func HotKey_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := hotKey_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetParentShowHint(obj uintptr, value bool) {
   hotKey_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetPopupMenu.Call(obj)
    return ret
}

func HotKey_SetPopupMenu(obj uintptr, value uintptr) {
   hotKey_SetPopupMenu.Call(obj, value)
}

func HotKey_GetShowHint(obj uintptr) bool {
    ret, _, _ := hotKey_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetShowHint(obj uintptr, value bool) {
   hotKey_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := hotKey_GetTabOrder.Call(obj)
    return int16(ret)
}

func HotKey_SetTabOrder(obj uintptr, value int16) {
   hotKey_SetTabOrder.Call(obj, uintptr(value))
}

func HotKey_GetTabStop(obj uintptr) bool {
    ret, _, _ := hotKey_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetTabStop(obj uintptr, value bool) {
   hotKey_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetVisible(obj uintptr) bool {
    ret, _, _ := hotKey_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetVisible(obj uintptr, value bool) {
   hotKey_SetVisible.Call(obj, GoBoolToDBool(value))
}

func HotKey_SetOnChange(obj uintptr, fn interface{}) {
    hotKey_SetOnChange.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnEnter(obj uintptr, fn interface{}) {
    hotKey_SetOnEnter.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnExit(obj uintptr, fn interface{}) {
    hotKey_SetOnExit.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseDown(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseEnter(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseLeave(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseMove(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseUp(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func HotKey_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := hotKey_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetDoubleBuffered(obj uintptr, value bool) {
   hotKey_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetBrush(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetBrush.Call(obj)
    return ret
}

func HotKey_GetControlCount(obj uintptr) int32 {
    ret, _, _ := hotKey_GetControlCount.Call(obj)
    return int32(ret)
}

func HotKey_GetHandle(obj uintptr) HWND {
    ret, _, _ := hotKey_GetHandle.Call(obj)
    return HWND(ret)
}

func HotKey_GetAction(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetAction.Call(obj)
    return ret
}

func HotKey_SetAction(obj uintptr, value uintptr) {
   hotKey_SetAction.Call(obj, value)
}

func HotKey_GetAlign(obj uintptr) TAlign {
    ret, _, _ := hotKey_GetAlign.Call(obj)
    return TAlign(ret)
}

func HotKey_SetAlign(obj uintptr, value TAlign) {
   hotKey_SetAlign.Call(obj, uintptr(value))
}

func HotKey_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    hotKey_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func HotKey_SetBoundsRect(obj uintptr, value TRect) {
   hotKey_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func HotKey_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetClientHeight.Call(obj)
    return int32(ret)
}

func HotKey_SetClientHeight(obj uintptr, value int32) {
   hotKey_SetClientHeight.Call(obj, uintptr(value))
}

func HotKey_GetClientRect(obj uintptr) TRect {
    var ret TRect
    hotKey_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func HotKey_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetClientWidth.Call(obj)
    return int32(ret)
}

func HotKey_SetClientWidth(obj uintptr, value int32) {
   hotKey_SetClientWidth.Call(obj, uintptr(value))
}

func HotKey_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitTop.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func HotKey_GetParent(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetParent.Call(obj)
    return ret
}

func HotKey_SetParent(obj uintptr, value uintptr) {
   hotKey_SetParent.Call(obj, value)
}

func HotKey_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := hotKey_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetAlignWithMargins(obj uintptr, value bool) {
   hotKey_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetLeft(obj uintptr) int32 {
    ret, _, _ := hotKey_GetLeft.Call(obj)
    return int32(ret)
}

func HotKey_SetLeft(obj uintptr, value int32) {
   hotKey_SetLeft.Call(obj, uintptr(value))
}

func HotKey_GetTop(obj uintptr) int32 {
    ret, _, _ := hotKey_GetTop.Call(obj)
    return int32(ret)
}

func HotKey_SetTop(obj uintptr, value int32) {
   hotKey_SetTop.Call(obj, uintptr(value))
}

func HotKey_GetWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetWidth.Call(obj)
    return int32(ret)
}

func HotKey_SetWidth(obj uintptr, value int32) {
   hotKey_SetWidth.Call(obj, uintptr(value))
}

func HotKey_GetHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetHeight.Call(obj)
    return int32(ret)
}

func HotKey_SetHeight(obj uintptr, value int32) {
   hotKey_SetHeight.Call(obj, uintptr(value))
}

func HotKey_GetCursor(obj uintptr) TCursor {
    ret, _, _ := hotKey_GetCursor.Call(obj)
    return TCursor(ret)
}

func HotKey_SetCursor(obj uintptr, value TCursor) {
   hotKey_SetCursor.Call(obj, uintptr(value))
}

func HotKey_GetMargins(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetMargins.Call(obj)
    return ret
}

func HotKey_SetMargins(obj uintptr, value uintptr) {
   hotKey_SetMargins.Call(obj, value)
}

func HotKey_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := hotKey_GetComponentCount.Call(obj)
    return int32(ret)
}

func HotKey_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := hotKey_GetComponentIndex.Call(obj)
    return int32(ret)
}

func HotKey_SetComponentIndex(obj uintptr, value int32) {
   hotKey_SetComponentIndex.Call(obj, uintptr(value))
}

func HotKey_GetOwner(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetOwner.Call(obj)
    return ret
}

func HotKey_GetName(obj uintptr) string {
    ret, _, _ := hotKey_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_SetName(obj uintptr, value string) {
   hotKey_SetName.Call(obj, GoStrToDStr(value))
}

func HotKey_GetTag(obj uintptr) int {
    ret, _, _ := hotKey_GetTag.Call(obj)
    return int(ret)
}

func HotKey_SetTag(obj uintptr, value int) {
   hotKey_SetTag.Call(obj, uintptr(value))
}

func HotKey_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := hotKey_GetControls.Call(obj, uintptr(Index))
    return ret
}

func HotKey_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := hotKey_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

