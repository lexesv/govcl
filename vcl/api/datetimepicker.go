
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func DateTimePicker_Create(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_Create.Call(obj)
    return ret
}

func DateTimePicker_Free(obj uintptr) {
    dateTimePicker_Free.Call(obj)
}

func DateTimePicker_CanFocus(obj uintptr) bool {
    ret, _, _ := dateTimePicker_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_FlipChildren(obj uintptr, AllLevels bool)  {
    dateTimePicker_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func DateTimePicker_Focused(obj uintptr) bool {
    ret, _, _ := dateTimePicker_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_HandleAllocated(obj uintptr) bool {
    ret, _, _ := dateTimePicker_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_Invalidate(obj uintptr)  {
    dateTimePicker_Invalidate.Call(obj)
}

func DateTimePicker_Realign(obj uintptr)  {
    dateTimePicker_Realign.Call(obj)
}

func DateTimePicker_Repaint(obj uintptr)  {
    dateTimePicker_Repaint.Call(obj)
}

func DateTimePicker_ScaleBy(obj uintptr, M int32, D int32)  {
    dateTimePicker_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func DateTimePicker_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    dateTimePicker_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func DateTimePicker_SetFocus(obj uintptr)  {
    dateTimePicker_SetFocus.Call(obj)
}

func DateTimePicker_Update(obj uintptr)  {
    dateTimePicker_Update.Call(obj)
}

func DateTimePicker_BringToFront(obj uintptr)  {
    dateTimePicker_BringToFront.Call(obj)
}

func DateTimePicker_HasParent(obj uintptr) bool {
    ret, _, _ := dateTimePicker_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_Hide(obj uintptr)  {
    dateTimePicker_Hide.Call(obj)
}

func DateTimePicker_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := dateTimePicker_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func DateTimePicker_Refresh(obj uintptr)  {
    dateTimePicker_Refresh.Call(obj)
}

func DateTimePicker_SendToBack(obj uintptr)  {
    dateTimePicker_SendToBack.Call(obj)
}

func DateTimePicker_Show(obj uintptr)  {
    dateTimePicker_Show.Call(obj)
}

func DateTimePicker_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := dateTimePicker_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func DateTimePicker_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := dateTimePicker_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func DateTimePicker_GetNamePath(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_Assign(obj uintptr, Source uintptr)  {
    dateTimePicker_Assign.Call(obj, Source )
}

func DateTimePicker_ClassName(obj uintptr) string {
    ret, _, _ := dateTimePicker_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := dateTimePicker_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func DateTimePicker_GetHashCode(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetHashCode.Call(obj)
    return int32(ret)
}

func DateTimePicker_ToString(obj uintptr) string {
    ret, _, _ := dateTimePicker_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_GetAlign(obj uintptr) TAlign {
    ret, _, _ := dateTimePicker_GetAlign.Call(obj)
    return TAlign(ret)
}

func DateTimePicker_SetAlign(obj uintptr, value TAlign) {
   dateTimePicker_SetAlign.Call(obj, uintptr(value))
}

func DateTimePicker_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := dateTimePicker_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func DateTimePicker_SetAnchors(obj uintptr, value TAnchors) {
   dateTimePicker_SetAnchors.Call(obj, uintptr(value))
}

func DateTimePicker_GetChecked(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetChecked(obj uintptr, value bool) {
   dateTimePicker_SetChecked.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetColor(obj uintptr) TColor {
    ret, _, _ := dateTimePicker_GetColor.Call(obj)
    return TColor(ret)
}

func DateTimePicker_SetColor(obj uintptr, value TColor) {
   dateTimePicker_SetColor.Call(obj, uintptr(value))
}

func DateTimePicker_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetDoubleBuffered(obj uintptr, value bool) {
   dateTimePicker_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetEnabled(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetEnabled(obj uintptr, value bool) {
   dateTimePicker_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetFont(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetFont.Call(obj)
    return ret
}

func DateTimePicker_SetFont(obj uintptr, value uintptr) {
   dateTimePicker_SetFont.Call(obj, value)
}

func DateTimePicker_GetParentColor(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentColor(obj uintptr, value bool) {
   dateTimePicker_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentFont(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentFont(obj uintptr, value bool) {
   dateTimePicker_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentShowHint(obj uintptr, value bool) {
   dateTimePicker_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetPopupMenu.Call(obj)
    return ret
}

func DateTimePicker_SetPopupMenu(obj uintptr, value uintptr) {
   dateTimePicker_SetPopupMenu.Call(obj, value)
}

func DateTimePicker_GetShowHint(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetShowHint(obj uintptr, value bool) {
   dateTimePicker_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := dateTimePicker_GetTabOrder.Call(obj)
    return int16(ret)
}

func DateTimePicker_SetTabOrder(obj uintptr, value int16) {
   dateTimePicker_SetTabOrder.Call(obj, uintptr(value))
}

func DateTimePicker_GetTabStop(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetTabStop(obj uintptr, value bool) {
   dateTimePicker_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetVisible(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetVisible(obj uintptr, value bool) {
   dateTimePicker_SetVisible.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_SetOnClick(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnClick.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnChange(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnChange.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnEnter(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnEnter.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnExit(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnExit.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyDown(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyPress(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyUp(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnMouseEnter(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnMouseLeave(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func DateTimePicker_GetBrush(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetBrush.Call(obj)
    return ret
}

func DateTimePicker_GetControlCount(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetControlCount.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetHandle(obj uintptr) HWND {
    ret, _, _ := dateTimePicker_GetHandle.Call(obj)
    return HWND(ret)
}

func DateTimePicker_GetAction(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetAction.Call(obj)
    return ret
}

func DateTimePicker_SetAction(obj uintptr, value uintptr) {
   dateTimePicker_SetAction.Call(obj, value)
}

func DateTimePicker_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    dateTimePicker_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func DateTimePicker_SetBoundsRect(obj uintptr, value TRect) {
   dateTimePicker_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DateTimePicker_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetClientHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetClientHeight(obj uintptr, value int32) {
   dateTimePicker_SetClientHeight.Call(obj, uintptr(value))
}

func DateTimePicker_GetClientRect(obj uintptr) TRect {
    var ret TRect
    dateTimePicker_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func DateTimePicker_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetClientWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetClientWidth(obj uintptr, value int32) {
   dateTimePicker_SetClientWidth.Call(obj, uintptr(value))
}

func DateTimePicker_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitTop.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetParent(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetParent.Call(obj)
    return ret
}

func DateTimePicker_SetParent(obj uintptr, value uintptr) {
   dateTimePicker_SetParent.Call(obj, value)
}

func DateTimePicker_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetAlignWithMargins(obj uintptr, value bool) {
   dateTimePicker_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetLeft(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetLeft.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetLeft(obj uintptr, value int32) {
   dateTimePicker_SetLeft.Call(obj, uintptr(value))
}

func DateTimePicker_GetTop(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetTop.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetTop(obj uintptr, value int32) {
   dateTimePicker_SetTop.Call(obj, uintptr(value))
}

func DateTimePicker_GetWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetWidth(obj uintptr, value int32) {
   dateTimePicker_SetWidth.Call(obj, uintptr(value))
}

func DateTimePicker_GetHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetHeight(obj uintptr, value int32) {
   dateTimePicker_SetHeight.Call(obj, uintptr(value))
}

func DateTimePicker_GetCursor(obj uintptr) TCursor {
    ret, _, _ := dateTimePicker_GetCursor.Call(obj)
    return TCursor(ret)
}

func DateTimePicker_SetCursor(obj uintptr, value TCursor) {
   dateTimePicker_SetCursor.Call(obj, uintptr(value))
}

func DateTimePicker_GetHint(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_SetHint(obj uintptr, value string) {
   dateTimePicker_SetHint.Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetMargins(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetMargins.Call(obj)
    return ret
}

func DateTimePicker_SetMargins(obj uintptr, value uintptr) {
   dateTimePicker_SetMargins.Call(obj, value)
}

func DateTimePicker_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetComponentCount.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetComponentIndex.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetComponentIndex(obj uintptr, value int32) {
   dateTimePicker_SetComponentIndex.Call(obj, uintptr(value))
}

func DateTimePicker_GetOwner(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetOwner.Call(obj)
    return ret
}

func DateTimePicker_GetName(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_SetName(obj uintptr, value string) {
   dateTimePicker_SetName.Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTag(obj uintptr) int {
    ret, _, _ := dateTimePicker_GetTag.Call(obj)
    return int(ret)
}

func DateTimePicker_SetTag(obj uintptr, value int) {
   dateTimePicker_SetTag.Call(obj, uintptr(value))
}

func DateTimePicker_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := dateTimePicker_GetControls.Call(obj, uintptr(Index))
    return ret
}

func DateTimePicker_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := dateTimePicker_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

