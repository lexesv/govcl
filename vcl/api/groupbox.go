
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func GroupBox_Create(obj uintptr) uintptr {
    ret, _, _ := groupBox_Create.Call(obj)
    return ret
}

func GroupBox_Free(obj uintptr) {
    groupBox_Free.Call(obj)
}

func GroupBox_CanFocus(obj uintptr) bool {
    ret, _, _ := groupBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_FlipChildren(obj uintptr, AllLevels bool)  {
    groupBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func GroupBox_Focused(obj uintptr) bool {
    ret, _, _ := groupBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := groupBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_Invalidate(obj uintptr)  {
    groupBox_Invalidate.Call(obj)
}

func GroupBox_Realign(obj uintptr)  {
    groupBox_Realign.Call(obj)
}

func GroupBox_Repaint(obj uintptr)  {
    groupBox_Repaint.Call(obj)
}

func GroupBox_ScaleBy(obj uintptr, M int32, D int32)  {
    groupBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func GroupBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    groupBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func GroupBox_SetFocus(obj uintptr)  {
    groupBox_SetFocus.Call(obj)
}

func GroupBox_Update(obj uintptr)  {
    groupBox_Update.Call(obj)
}

func GroupBox_BringToFront(obj uintptr)  {
    groupBox_BringToFront.Call(obj)
}

func GroupBox_HasParent(obj uintptr) bool {
    ret, _, _ := groupBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_Hide(obj uintptr)  {
    groupBox_Hide.Call(obj)
}

func GroupBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := groupBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func GroupBox_Refresh(obj uintptr)  {
    groupBox_Refresh.Call(obj)
}

func GroupBox_SendToBack(obj uintptr)  {
    groupBox_SendToBack.Call(obj)
}

func GroupBox_Show(obj uintptr)  {
    groupBox_Show.Call(obj)
}

func GroupBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := groupBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func GroupBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := groupBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func GroupBox_GetNamePath(obj uintptr) string {
    ret, _, _ := groupBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_Assign(obj uintptr, Source uintptr)  {
    groupBox_Assign.Call(obj, Source )
}

func GroupBox_ClassName(obj uintptr) string {
    ret, _, _ := groupBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := groupBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func GroupBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := groupBox_GetHashCode.Call(obj)
    return int32(ret)
}

func GroupBox_ToString(obj uintptr) string {
    ret, _, _ := groupBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := groupBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func GroupBox_SetAlign(obj uintptr, value TAlign) {
   groupBox_SetAlign.Call(obj, uintptr(value))
}

func GroupBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := groupBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func GroupBox_SetAnchors(obj uintptr, value TAnchors) {
   groupBox_SetAnchors.Call(obj, uintptr(value))
}

func GroupBox_GetCaption(obj uintptr) string {
    ret, _, _ := groupBox_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetCaption(obj uintptr, value string) {
   groupBox_SetCaption.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetColor(obj uintptr) TColor {
    ret, _, _ := groupBox_GetColor.Call(obj)
    return TColor(ret)
}

func GroupBox_SetColor(obj uintptr, value TColor) {
   groupBox_SetColor.Call(obj, uintptr(value))
}

func GroupBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := groupBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetDoubleBuffered(obj uintptr, value bool) {
   groupBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := groupBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetEnabled(obj uintptr, value bool) {
   groupBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetFont.Call(obj)
    return ret
}

func GroupBox_SetFont(obj uintptr, value uintptr) {
   groupBox_SetFont.Call(obj, value)
}

func GroupBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentColor(obj uintptr, value bool) {
   groupBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentCtl3D(obj uintptr, value bool) {
   groupBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentFont(obj uintptr, value bool) {
   groupBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentShowHint(obj uintptr, value bool) {
   groupBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetPopupMenu.Call(obj)
    return ret
}

func GroupBox_SetPopupMenu(obj uintptr, value uintptr) {
   groupBox_SetPopupMenu.Call(obj, value)
}

func GroupBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := groupBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetShowHint(obj uintptr, value bool) {
   groupBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := groupBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func GroupBox_SetTabOrder(obj uintptr, value int16) {
   groupBox_SetTabOrder.Call(obj, uintptr(value))
}

func GroupBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := groupBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetTabStop(obj uintptr, value bool) {
   groupBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetVisible(obj uintptr) bool {
    ret, _, _ := groupBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetVisible(obj uintptr, value bool) {
   groupBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func GroupBox_SetOnClick(obj uintptr, fn interface{}) {
    groupBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnDblClick(obj uintptr, fn interface{}) {
    groupBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnEnter(obj uintptr, fn interface{}) {
    groupBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnExit(obj uintptr, fn interface{}) {
    groupBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func GroupBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetBrush.Call(obj)
    return ret
}

func GroupBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := groupBox_GetControlCount.Call(obj)
    return int32(ret)
}

func GroupBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := groupBox_GetHandle.Call(obj)
    return HWND(ret)
}

func GroupBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetAction.Call(obj)
    return ret
}

func GroupBox_SetAction(obj uintptr, value uintptr) {
   groupBox_SetAction.Call(obj, value)
}

func GroupBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    groupBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GroupBox_SetBoundsRect(obj uintptr, value TRect) {
   groupBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func GroupBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func GroupBox_SetClientHeight(obj uintptr, value int32) {
   groupBox_SetClientHeight.Call(obj, uintptr(value))
}

func GroupBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    groupBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GroupBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func GroupBox_SetClientWidth(obj uintptr, value int32) {
   groupBox_SetClientWidth.Call(obj, uintptr(value))
}

func GroupBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func GroupBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetParent.Call(obj)
    return ret
}

func GroupBox_SetParent(obj uintptr, value uintptr) {
   groupBox_SetParent.Call(obj, value)
}

func GroupBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := groupBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetAlignWithMargins(obj uintptr, value bool) {
   groupBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := groupBox_GetLeft.Call(obj)
    return int32(ret)
}

func GroupBox_SetLeft(obj uintptr, value int32) {
   groupBox_SetLeft.Call(obj, uintptr(value))
}

func GroupBox_GetTop(obj uintptr) int32 {
    ret, _, _ := groupBox_GetTop.Call(obj)
    return int32(ret)
}

func GroupBox_SetTop(obj uintptr, value int32) {
   groupBox_SetTop.Call(obj, uintptr(value))
}

func GroupBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetWidth.Call(obj)
    return int32(ret)
}

func GroupBox_SetWidth(obj uintptr, value int32) {
   groupBox_SetWidth.Call(obj, uintptr(value))
}

func GroupBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetHeight.Call(obj)
    return int32(ret)
}

func GroupBox_SetHeight(obj uintptr, value int32) {
   groupBox_SetHeight.Call(obj, uintptr(value))
}

func GroupBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := groupBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func GroupBox_SetCursor(obj uintptr, value TCursor) {
   groupBox_SetCursor.Call(obj, uintptr(value))
}

func GroupBox_GetHint(obj uintptr) string {
    ret, _, _ := groupBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetHint(obj uintptr, value string) {
   groupBox_SetHint.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetMargins.Call(obj)
    return ret
}

func GroupBox_SetMargins(obj uintptr, value uintptr) {
   groupBox_SetMargins.Call(obj, value)
}

func GroupBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := groupBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func GroupBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := groupBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func GroupBox_SetComponentIndex(obj uintptr, value int32) {
   groupBox_SetComponentIndex.Call(obj, uintptr(value))
}

func GroupBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetOwner.Call(obj)
    return ret
}

func GroupBox_GetName(obj uintptr) string {
    ret, _, _ := groupBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetName(obj uintptr, value string) {
   groupBox_SetName.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetTag(obj uintptr) int {
    ret, _, _ := groupBox_GetTag.Call(obj)
    return int(ret)
}

func GroupBox_SetTag(obj uintptr, value int) {
   groupBox_SetTag.Call(obj, uintptr(value))
}

func GroupBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := groupBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func GroupBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := groupBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

