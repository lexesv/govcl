
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func MonthCalendar_Create(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_Create.Call(obj)
    return ret
}

func MonthCalendar_Free(obj uintptr) {
    monthCalendar_Free.Call(obj)
}

func MonthCalendar_CanFocus(obj uintptr) bool {
    ret, _, _ := monthCalendar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_FlipChildren(obj uintptr, AllLevels bool)  {
    monthCalendar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func MonthCalendar_Focused(obj uintptr) bool {
    ret, _, _ := monthCalendar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := monthCalendar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_Invalidate(obj uintptr)  {
    monthCalendar_Invalidate.Call(obj)
}

func MonthCalendar_Realign(obj uintptr)  {
    monthCalendar_Realign.Call(obj)
}

func MonthCalendar_Repaint(obj uintptr)  {
    monthCalendar_Repaint.Call(obj)
}

func MonthCalendar_ScaleBy(obj uintptr, M int32, D int32)  {
    monthCalendar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func MonthCalendar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    monthCalendar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func MonthCalendar_SetFocus(obj uintptr)  {
    monthCalendar_SetFocus.Call(obj)
}

func MonthCalendar_Update(obj uintptr)  {
    monthCalendar_Update.Call(obj)
}

func MonthCalendar_BringToFront(obj uintptr)  {
    monthCalendar_BringToFront.Call(obj)
}

func MonthCalendar_HasParent(obj uintptr) bool {
    ret, _, _ := monthCalendar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_Hide(obj uintptr)  {
    monthCalendar_Hide.Call(obj)
}

func MonthCalendar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := monthCalendar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func MonthCalendar_Refresh(obj uintptr)  {
    monthCalendar_Refresh.Call(obj)
}

func MonthCalendar_SendToBack(obj uintptr)  {
    monthCalendar_SendToBack.Call(obj)
}

func MonthCalendar_Show(obj uintptr)  {
    monthCalendar_Show.Call(obj)
}

func MonthCalendar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := monthCalendar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func MonthCalendar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := monthCalendar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MonthCalendar_GetNamePath(obj uintptr) string {
    ret, _, _ := monthCalendar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_Assign(obj uintptr, Source uintptr)  {
    monthCalendar_Assign.Call(obj, Source )
}

func MonthCalendar_ClassName(obj uintptr) string {
    ret, _, _ := monthCalendar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := monthCalendar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MonthCalendar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetHashCode.Call(obj)
    return int32(ret)
}

func MonthCalendar_ToString(obj uintptr) string {
    ret, _, _ := monthCalendar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := monthCalendar_GetAlign.Call(obj)
    return TAlign(ret)
}

func MonthCalendar_SetAlign(obj uintptr, value TAlign) {
   monthCalendar_SetAlign.Call(obj, uintptr(value))
}

func MonthCalendar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := monthCalendar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func MonthCalendar_SetAnchors(obj uintptr, value TAnchors) {
   monthCalendar_SetAnchors.Call(obj, uintptr(value))
}

func MonthCalendar_GetAutoSize(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetAutoSize(obj uintptr, value bool) {
   monthCalendar_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetBorderWidth(obj uintptr, value int32) {
   monthCalendar_SetBorderWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetMultiSelect(obj uintptr, value bool) {
   monthCalendar_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetDoubleBuffered(obj uintptr, value bool) {
   monthCalendar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetEnabled(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetEnabled(obj uintptr, value bool) {
   monthCalendar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetFont(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetFont.Call(obj)
    return ret
}

func MonthCalendar_SetFont(obj uintptr, value uintptr) {
   monthCalendar_SetFont.Call(obj, value)
}

func MonthCalendar_GetParentFont(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentFont(obj uintptr, value bool) {
   monthCalendar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentShowHint(obj uintptr, value bool) {
   monthCalendar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetPopupMenu.Call(obj)
    return ret
}

func MonthCalendar_SetPopupMenu(obj uintptr, value uintptr) {
   monthCalendar_SetPopupMenu.Call(obj, value)
}

func MonthCalendar_GetShowHint(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetShowHint(obj uintptr, value bool) {
   monthCalendar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := monthCalendar_GetTabOrder.Call(obj)
    return int16(ret)
}

func MonthCalendar_SetTabOrder(obj uintptr, value int16) {
   monthCalendar_SetTabOrder.Call(obj, uintptr(value))
}

func MonthCalendar_GetTabStop(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetTabStop(obj uintptr, value bool) {
   monthCalendar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetVisible(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetVisible(obj uintptr, value bool) {
   monthCalendar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_SetOnClick(obj uintptr, fn interface{}) {
    monthCalendar_SetOnClick.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnDblClick(obj uintptr, fn interface{}) {
    monthCalendar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnEnter(obj uintptr, fn interface{}) {
    monthCalendar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnExit(obj uintptr, fn interface{}) {
    monthCalendar_SetOnExit.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyDown(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyPress(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyUp(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    monthCalendar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    monthCalendar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func MonthCalendar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetBrush.Call(obj)
    return ret
}

func MonthCalendar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetControlCount.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetHandle(obj uintptr) HWND {
    ret, _, _ := monthCalendar_GetHandle.Call(obj)
    return HWND(ret)
}

func MonthCalendar_GetAction(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetAction.Call(obj)
    return ret
}

func MonthCalendar_SetAction(obj uintptr, value uintptr) {
   monthCalendar_SetAction.Call(obj, value)
}

func MonthCalendar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    monthCalendar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MonthCalendar_SetBoundsRect(obj uintptr, value TRect) {
   monthCalendar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MonthCalendar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetClientHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetClientHeight(obj uintptr, value int32) {
   monthCalendar_SetClientHeight.Call(obj, uintptr(value))
}

func MonthCalendar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    monthCalendar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MonthCalendar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetClientWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetClientWidth(obj uintptr, value int32) {
   monthCalendar_SetClientWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetParent(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetParent.Call(obj)
    return ret
}

func MonthCalendar_SetParent(obj uintptr, value uintptr) {
   monthCalendar_SetParent.Call(obj, value)
}

func MonthCalendar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetAlignWithMargins(obj uintptr, value bool) {
   monthCalendar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetLeft(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetLeft.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetLeft(obj uintptr, value int32) {
   monthCalendar_SetLeft.Call(obj, uintptr(value))
}

func MonthCalendar_GetTop(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetTop.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetTop(obj uintptr, value int32) {
   monthCalendar_SetTop.Call(obj, uintptr(value))
}

func MonthCalendar_GetWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetWidth(obj uintptr, value int32) {
   monthCalendar_SetWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetHeight(obj uintptr, value int32) {
   monthCalendar_SetHeight.Call(obj, uintptr(value))
}

func MonthCalendar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := monthCalendar_GetCursor.Call(obj)
    return TCursor(ret)
}

func MonthCalendar_SetCursor(obj uintptr, value TCursor) {
   monthCalendar_SetCursor.Call(obj, uintptr(value))
}

func MonthCalendar_GetHint(obj uintptr) string {
    ret, _, _ := monthCalendar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_SetHint(obj uintptr, value string) {
   monthCalendar_SetHint.Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetMargins.Call(obj)
    return ret
}

func MonthCalendar_SetMargins(obj uintptr, value uintptr) {
   monthCalendar_SetMargins.Call(obj, value)
}

func MonthCalendar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetComponentCount.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetComponentIndex(obj uintptr, value int32) {
   monthCalendar_SetComponentIndex.Call(obj, uintptr(value))
}

func MonthCalendar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetOwner.Call(obj)
    return ret
}

func MonthCalendar_GetName(obj uintptr) string {
    ret, _, _ := monthCalendar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_SetName(obj uintptr, value string) {
   monthCalendar_SetName.Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetTag(obj uintptr) int {
    ret, _, _ := monthCalendar_GetTag.Call(obj)
    return int(ret)
}

func MonthCalendar_SetTag(obj uintptr, value int) {
   monthCalendar_SetTag.Call(obj, uintptr(value))
}

func MonthCalendar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := monthCalendar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func MonthCalendar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := monthCalendar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

