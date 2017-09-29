
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Panel_Create(obj uintptr) uintptr {
    ret, _, _ := panel_Create.Call(obj)
    return ret
}

func Panel_Free(obj uintptr) {
    panel_Free.Call(obj)
}

func Panel_CanFocus(obj uintptr) bool {
    ret, _, _ := panel_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_FlipChildren(obj uintptr, AllLevels bool)  {
    panel_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Panel_Focused(obj uintptr) bool {
    ret, _, _ := panel_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_HandleAllocated(obj uintptr) bool {
    ret, _, _ := panel_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_Invalidate(obj uintptr)  {
    panel_Invalidate.Call(obj)
}

func Panel_Realign(obj uintptr)  {
    panel_Realign.Call(obj)
}

func Panel_Repaint(obj uintptr)  {
    panel_Repaint.Call(obj)
}

func Panel_ScaleBy(obj uintptr, M int32, D int32)  {
    panel_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Panel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    panel_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Panel_SetFocus(obj uintptr)  {
    panel_SetFocus.Call(obj)
}

func Panel_Update(obj uintptr)  {
    panel_Update.Call(obj)
}

func Panel_BringToFront(obj uintptr)  {
    panel_BringToFront.Call(obj)
}

func Panel_HasParent(obj uintptr) bool {
    ret, _, _ := panel_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_Hide(obj uintptr)  {
    panel_Hide.Call(obj)
}

func Panel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := panel_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Panel_Refresh(obj uintptr)  {
    panel_Refresh.Call(obj)
}

func Panel_SendToBack(obj uintptr)  {
    panel_SendToBack.Call(obj)
}

func Panel_Show(obj uintptr)  {
    panel_Show.Call(obj)
}

func Panel_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := panel_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Panel_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := panel_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Panel_GetNamePath(obj uintptr) string {
    ret, _, _ := panel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_Assign(obj uintptr, Source uintptr)  {
    panel_Assign.Call(obj, Source )
}

func Panel_ClassName(obj uintptr) string {
    ret, _, _ := panel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := panel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Panel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := panel_GetHashCode.Call(obj)
    return int32(ret)
}

func Panel_ToString(obj uintptr) string {
    ret, _, _ := panel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_GetAlign(obj uintptr) TAlign {
    ret, _, _ := panel_GetAlign.Call(obj)
    return TAlign(ret)
}

func Panel_SetAlign(obj uintptr, value TAlign) {
   panel_SetAlign.Call(obj, uintptr(value))
}

func Panel_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := panel_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Panel_SetAlignment(obj uintptr, value TAlignment) {
   panel_SetAlignment.Call(obj, uintptr(value))
}

func Panel_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := panel_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Panel_SetAnchors(obj uintptr, value TAnchors) {
   panel_SetAnchors.Call(obj, uintptr(value))
}

func Panel_GetAutoSize(obj uintptr) bool {
    ret, _, _ := panel_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetAutoSize(obj uintptr, value bool) {
   panel_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Panel_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetBorderWidth.Call(obj)
    return int32(ret)
}

func Panel_SetBorderWidth(obj uintptr, value int32) {
   panel_SetBorderWidth.Call(obj, uintptr(value))
}

func Panel_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := panel_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Panel_SetBorderStyle(obj uintptr, value TBorderStyle) {
   panel_SetBorderStyle.Call(obj, uintptr(value))
}

func Panel_GetCaption(obj uintptr) string {
    ret, _, _ := panel_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetCaption(obj uintptr, value string) {
   panel_SetCaption.Call(obj, GoStrToDStr(value))
}

func Panel_GetColor(obj uintptr) TColor {
    ret, _, _ := panel_GetColor.Call(obj)
    return TColor(ret)
}

func Panel_SetColor(obj uintptr, value TColor) {
   panel_SetColor.Call(obj, uintptr(value))
}

func Panel_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := panel_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetDoubleBuffered(obj uintptr, value bool) {
   panel_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Panel_GetEnabled(obj uintptr) bool {
    ret, _, _ := panel_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetEnabled(obj uintptr, value bool) {
   panel_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Panel_GetFont(obj uintptr) uintptr {
    ret, _, _ := panel_GetFont.Call(obj)
    return ret
}

func Panel_SetFont(obj uintptr, value uintptr) {
   panel_SetFont.Call(obj, value)
}

func Panel_GetParentColor(obj uintptr) bool {
    ret, _, _ := panel_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentColor(obj uintptr, value bool) {
   panel_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := panel_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentCtl3D(obj uintptr, value bool) {
   panel_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentFont(obj uintptr) bool {
    ret, _, _ := panel_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentFont(obj uintptr, value bool) {
   panel_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := panel_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentShowHint(obj uintptr, value bool) {
   panel_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Panel_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := panel_GetPopupMenu.Call(obj)
    return ret
}

func Panel_SetPopupMenu(obj uintptr, value uintptr) {
   panel_SetPopupMenu.Call(obj, value)
}

func Panel_GetShowCaption(obj uintptr) bool {
    ret, _, _ := panel_GetShowCaption.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetShowCaption(obj uintptr, value bool) {
   panel_SetShowCaption.Call(obj, GoBoolToDBool(value))
}

func Panel_GetShowHint(obj uintptr) bool {
    ret, _, _ := panel_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetShowHint(obj uintptr, value bool) {
   panel_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Panel_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := panel_GetTabOrder.Call(obj)
    return int16(ret)
}

func Panel_SetTabOrder(obj uintptr, value int16) {
   panel_SetTabOrder.Call(obj, uintptr(value))
}

func Panel_GetTabStop(obj uintptr) bool {
    ret, _, _ := panel_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetTabStop(obj uintptr, value bool) {
   panel_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Panel_GetVisible(obj uintptr) bool {
    ret, _, _ := panel_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetVisible(obj uintptr, value bool) {
   panel_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Panel_SetOnClick(obj uintptr, fn interface{}) {
    panel_SetOnClick.Call(obj, addEventToMap(fn))
}

func Panel_SetOnDblClick(obj uintptr, fn interface{}) {
    panel_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Panel_SetOnEnter(obj uintptr, fn interface{}) {
    panel_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Panel_SetOnExit(obj uintptr, fn interface{}) {
    panel_SetOnExit.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseDown(obj uintptr, fn interface{}) {
    panel_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseEnter(obj uintptr, fn interface{}) {
    panel_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseLeave(obj uintptr, fn interface{}) {
    panel_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseMove(obj uintptr, fn interface{}) {
    panel_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseUp(obj uintptr, fn interface{}) {
    panel_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Panel_SetOnResize(obj uintptr, fn interface{}) {
    panel_SetOnResize.Call(obj, addEventToMap(fn))
}

func Panel_GetBrush(obj uintptr) uintptr {
    ret, _, _ := panel_GetBrush.Call(obj)
    return ret
}

func Panel_GetControlCount(obj uintptr) int32 {
    ret, _, _ := panel_GetControlCount.Call(obj)
    return int32(ret)
}

func Panel_GetHandle(obj uintptr) HWND {
    ret, _, _ := panel_GetHandle.Call(obj)
    return HWND(ret)
}

func Panel_GetAction(obj uintptr) uintptr {
    ret, _, _ := panel_GetAction.Call(obj)
    return ret
}

func Panel_SetAction(obj uintptr, value uintptr) {
   panel_SetAction.Call(obj, value)
}

func Panel_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    panel_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Panel_SetBoundsRect(obj uintptr, value TRect) {
   panel_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Panel_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetClientHeight.Call(obj)
    return int32(ret)
}

func Panel_SetClientHeight(obj uintptr, value int32) {
   panel_SetClientHeight.Call(obj, uintptr(value))
}

func Panel_GetClientRect(obj uintptr) TRect {
    var ret TRect
    panel_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Panel_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetClientWidth.Call(obj)
    return int32(ret)
}

func Panel_SetClientWidth(obj uintptr, value int32) {
   panel_SetClientWidth.Call(obj, uintptr(value))
}

func Panel_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Panel_GetParent(obj uintptr) uintptr {
    ret, _, _ := panel_GetParent.Call(obj)
    return ret
}

func Panel_SetParent(obj uintptr, value uintptr) {
   panel_SetParent.Call(obj, value)
}

func Panel_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := panel_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetAlignWithMargins(obj uintptr, value bool) {
   panel_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Panel_GetLeft(obj uintptr) int32 {
    ret, _, _ := panel_GetLeft.Call(obj)
    return int32(ret)
}

func Panel_SetLeft(obj uintptr, value int32) {
   panel_SetLeft.Call(obj, uintptr(value))
}

func Panel_GetTop(obj uintptr) int32 {
    ret, _, _ := panel_GetTop.Call(obj)
    return int32(ret)
}

func Panel_SetTop(obj uintptr, value int32) {
   panel_SetTop.Call(obj, uintptr(value))
}

func Panel_GetWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetWidth.Call(obj)
    return int32(ret)
}

func Panel_SetWidth(obj uintptr, value int32) {
   panel_SetWidth.Call(obj, uintptr(value))
}

func Panel_GetHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetHeight.Call(obj)
    return int32(ret)
}

func Panel_SetHeight(obj uintptr, value int32) {
   panel_SetHeight.Call(obj, uintptr(value))
}

func Panel_GetCursor(obj uintptr) TCursor {
    ret, _, _ := panel_GetCursor.Call(obj)
    return TCursor(ret)
}

func Panel_SetCursor(obj uintptr, value TCursor) {
   panel_SetCursor.Call(obj, uintptr(value))
}

func Panel_GetHint(obj uintptr) string {
    ret, _, _ := panel_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetHint(obj uintptr, value string) {
   panel_SetHint.Call(obj, GoStrToDStr(value))
}

func Panel_GetMargins(obj uintptr) uintptr {
    ret, _, _ := panel_GetMargins.Call(obj)
    return ret
}

func Panel_SetMargins(obj uintptr, value uintptr) {
   panel_SetMargins.Call(obj, value)
}

func Panel_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := panel_GetComponentCount.Call(obj)
    return int32(ret)
}

func Panel_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := panel_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Panel_SetComponentIndex(obj uintptr, value int32) {
   panel_SetComponentIndex.Call(obj, uintptr(value))
}

func Panel_GetOwner(obj uintptr) uintptr {
    ret, _, _ := panel_GetOwner.Call(obj)
    return ret
}

func Panel_GetName(obj uintptr) string {
    ret, _, _ := panel_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetName(obj uintptr, value string) {
   panel_SetName.Call(obj, GoStrToDStr(value))
}

func Panel_GetTag(obj uintptr) int {
    ret, _, _ := panel_GetTag.Call(obj)
    return int(ret)
}

func Panel_SetTag(obj uintptr, value int) {
   panel_SetTag.Call(obj, uintptr(value))
}

func Panel_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := panel_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Panel_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := panel_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

