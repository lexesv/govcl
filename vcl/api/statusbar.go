
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func StatusBar_Create(obj uintptr) uintptr {
    ret, _, _ := statusBar_Create.Call(obj)
    return ret
}

func StatusBar_Free(obj uintptr) {
    statusBar_Free.Call(obj)
}

func StatusBar_FlipChildren(obj uintptr, AllLevels bool)  {
    statusBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func StatusBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    statusBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func StatusBar_CanFocus(obj uintptr) bool {
    ret, _, _ := statusBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Focused(obj uintptr) bool {
    ret, _, _ := statusBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := statusBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Invalidate(obj uintptr)  {
    statusBar_Invalidate.Call(obj)
}

func StatusBar_Realign(obj uintptr)  {
    statusBar_Realign.Call(obj)
}

func StatusBar_Repaint(obj uintptr)  {
    statusBar_Repaint.Call(obj)
}

func StatusBar_ScaleBy(obj uintptr, M int32, D int32)  {
    statusBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func StatusBar_SetFocus(obj uintptr)  {
    statusBar_SetFocus.Call(obj)
}

func StatusBar_Update(obj uintptr)  {
    statusBar_Update.Call(obj)
}

func StatusBar_BringToFront(obj uintptr)  {
    statusBar_BringToFront.Call(obj)
}

func StatusBar_HasParent(obj uintptr) bool {
    ret, _, _ := statusBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Hide(obj uintptr)  {
    statusBar_Hide.Call(obj)
}

func StatusBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := statusBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func StatusBar_Refresh(obj uintptr)  {
    statusBar_Refresh.Call(obj)
}

func StatusBar_SendToBack(obj uintptr)  {
    statusBar_SendToBack.Call(obj)
}

func StatusBar_Show(obj uintptr)  {
    statusBar_Show.Call(obj)
}

func StatusBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := statusBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func StatusBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := statusBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func StatusBar_GetNamePath(obj uintptr) string {
    ret, _, _ := statusBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_Assign(obj uintptr, Source uintptr)  {
    statusBar_Assign.Call(obj, Source )
}

func StatusBar_ClassName(obj uintptr) string {
    ret, _, _ := statusBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := statusBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StatusBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := statusBar_GetHashCode.Call(obj)
    return int32(ret)
}

func StatusBar_ToString(obj uintptr) string {
    ret, _, _ := statusBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetAction.Call(obj)
    return ret
}

func StatusBar_SetAction(obj uintptr, value uintptr) {
   statusBar_SetAction.Call(obj, value)
}

func StatusBar_GetAutoHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetAutoHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetAutoHint(obj uintptr, value bool) {
   statusBar_SetAutoHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := statusBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func StatusBar_SetAlign(obj uintptr, value TAlign) {
   statusBar_SetAlign.Call(obj, uintptr(value))
}

func StatusBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := statusBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func StatusBar_SetAnchors(obj uintptr, value TAnchors) {
   statusBar_SetAnchors.Call(obj, uintptr(value))
}

func StatusBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetBorderWidth(obj uintptr, value int32) {
   statusBar_SetBorderWidth.Call(obj, uintptr(value))
}

func StatusBar_GetColor(obj uintptr) TColor {
    ret, _, _ := statusBar_GetColor.Call(obj)
    return TColor(ret)
}

func StatusBar_SetColor(obj uintptr, value TColor) {
   statusBar_SetColor.Call(obj, uintptr(value))
}

func StatusBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := statusBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetDoubleBuffered(obj uintptr, value bool) {
   statusBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := statusBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetEnabled(obj uintptr, value bool) {
   statusBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetFont(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetFont.Call(obj)
    return ret
}

func StatusBar_SetFont(obj uintptr, value uintptr) {
   statusBar_SetFont.Call(obj, value)
}

func StatusBar_GetPanels(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetPanels.Call(obj)
    return ret
}

func StatusBar_SetPanels(obj uintptr, value uintptr) {
   statusBar_SetPanels.Call(obj, value)
}

func StatusBar_GetParentColor(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentColor(obj uintptr, value bool) {
   statusBar_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentFont(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentFont(obj uintptr, value bool) {
   statusBar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentShowHint(obj uintptr, value bool) {
   statusBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetPopupMenu.Call(obj)
    return ret
}

func StatusBar_SetPopupMenu(obj uintptr, value uintptr) {
   statusBar_SetPopupMenu.Call(obj, value)
}

func StatusBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetShowHint(obj uintptr, value bool) {
   statusBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimplePanel(obj uintptr) bool {
    ret, _, _ := statusBar_GetSimplePanel.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetSimplePanel(obj uintptr, value bool) {
   statusBar_SetSimplePanel.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimpleText(obj uintptr) string {
    ret, _, _ := statusBar_GetSimpleText.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetSimpleText(obj uintptr, value string) {
   statusBar_SetSimpleText.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetSizeGrip(obj uintptr) bool {
    ret, _, _ := statusBar_GetSizeGrip.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetSizeGrip(obj uintptr, value bool) {
   statusBar_SetSizeGrip.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetUseSystemFont(obj uintptr) bool {
    ret, _, _ := statusBar_GetUseSystemFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetUseSystemFont(obj uintptr, value bool) {
   statusBar_SetUseSystemFont.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetVisible(obj uintptr) bool {
    ret, _, _ := statusBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetVisible(obj uintptr, value bool) {
   statusBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func StatusBar_SetOnClick(obj uintptr, fn interface{}) {
    statusBar_SetOnClick.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnDblClick(obj uintptr, fn interface{}) {
    statusBar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnResize(obj uintptr, fn interface{}) {
    statusBar_SetOnResize.Call(obj, addEventToMap(fn))
}

func StatusBar_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetCanvas.Call(obj)
    return ret
}

func StatusBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetBrush.Call(obj)
    return ret
}

func StatusBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := statusBar_GetControlCount.Call(obj)
    return int32(ret)
}

func StatusBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := statusBar_GetHandle.Call(obj)
    return HWND(ret)
}

func StatusBar_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := statusBar_GetTabOrder.Call(obj)
    return int16(ret)
}

func StatusBar_SetTabOrder(obj uintptr, value int16) {
   statusBar_SetTabOrder.Call(obj, uintptr(value))
}

func StatusBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := statusBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetTabStop(obj uintptr, value bool) {
   statusBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    statusBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StatusBar_SetBoundsRect(obj uintptr, value TRect) {
   statusBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StatusBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func StatusBar_SetClientHeight(obj uintptr, value int32) {
   statusBar_SetClientHeight.Call(obj, uintptr(value))
}

func StatusBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    statusBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StatusBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetClientWidth(obj uintptr, value int32) {
   statusBar_SetClientWidth.Call(obj, uintptr(value))
}

func StatusBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func StatusBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetParent.Call(obj)
    return ret
}

func StatusBar_SetParent(obj uintptr, value uintptr) {
   statusBar_SetParent.Call(obj, value)
}

func StatusBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := statusBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetAlignWithMargins(obj uintptr, value bool) {
   statusBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := statusBar_GetLeft.Call(obj)
    return int32(ret)
}

func StatusBar_SetLeft(obj uintptr, value int32) {
   statusBar_SetLeft.Call(obj, uintptr(value))
}

func StatusBar_GetTop(obj uintptr) int32 {
    ret, _, _ := statusBar_GetTop.Call(obj)
    return int32(ret)
}

func StatusBar_SetTop(obj uintptr, value int32) {
   statusBar_SetTop.Call(obj, uintptr(value))
}

func StatusBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetWidth(obj uintptr, value int32) {
   statusBar_SetWidth.Call(obj, uintptr(value))
}

func StatusBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetHeight.Call(obj)
    return int32(ret)
}

func StatusBar_SetHeight(obj uintptr, value int32) {
   statusBar_SetHeight.Call(obj, uintptr(value))
}

func StatusBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := statusBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func StatusBar_SetCursor(obj uintptr, value TCursor) {
   statusBar_SetCursor.Call(obj, uintptr(value))
}

func StatusBar_GetHint(obj uintptr) string {
    ret, _, _ := statusBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetHint(obj uintptr, value string) {
   statusBar_SetHint.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetMargins.Call(obj)
    return ret
}

func StatusBar_SetMargins(obj uintptr, value uintptr) {
   statusBar_SetMargins.Call(obj, value)
}

func StatusBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := statusBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func StatusBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := statusBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func StatusBar_SetComponentIndex(obj uintptr, value int32) {
   statusBar_SetComponentIndex.Call(obj, uintptr(value))
}

func StatusBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetOwner.Call(obj)
    return ret
}

func StatusBar_GetName(obj uintptr) string {
    ret, _, _ := statusBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetName(obj uintptr, value string) {
   statusBar_SetName.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetTag(obj uintptr) int {
    ret, _, _ := statusBar_GetTag.Call(obj)
    return int(ret)
}

func StatusBar_SetTag(obj uintptr, value int) {
   statusBar_SetTag.Call(obj, uintptr(value))
}

func StatusBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := statusBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func StatusBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := statusBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

