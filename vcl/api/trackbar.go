
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func TrackBar_Create(obj uintptr) uintptr {
    ret, _, _ := trackBar_Create.Call(obj)
    return ret
}

func TrackBar_Free(obj uintptr) {
    trackBar_Free.Call(obj)
}

func TrackBar_CanFocus(obj uintptr) bool {
    ret, _, _ := trackBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_FlipChildren(obj uintptr, AllLevels bool)  {
    trackBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TrackBar_Focused(obj uintptr) bool {
    ret, _, _ := trackBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := trackBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_Invalidate(obj uintptr)  {
    trackBar_Invalidate.Call(obj)
}

func TrackBar_Realign(obj uintptr)  {
    trackBar_Realign.Call(obj)
}

func TrackBar_Repaint(obj uintptr)  {
    trackBar_Repaint.Call(obj)
}

func TrackBar_ScaleBy(obj uintptr, M int32, D int32)  {
    trackBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TrackBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    trackBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TrackBar_SetFocus(obj uintptr)  {
    trackBar_SetFocus.Call(obj)
}

func TrackBar_Update(obj uintptr)  {
    trackBar_Update.Call(obj)
}

func TrackBar_BringToFront(obj uintptr)  {
    trackBar_BringToFront.Call(obj)
}

func TrackBar_HasParent(obj uintptr) bool {
    ret, _, _ := trackBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_Hide(obj uintptr)  {
    trackBar_Hide.Call(obj)
}

func TrackBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := trackBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TrackBar_Refresh(obj uintptr)  {
    trackBar_Refresh.Call(obj)
}

func TrackBar_SendToBack(obj uintptr)  {
    trackBar_SendToBack.Call(obj)
}

func TrackBar_Show(obj uintptr)  {
    trackBar_Show.Call(obj)
}

func TrackBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := trackBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TrackBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := trackBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TrackBar_GetNamePath(obj uintptr) string {
    ret, _, _ := trackBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_Assign(obj uintptr, Source uintptr)  {
    trackBar_Assign.Call(obj, Source )
}

func TrackBar_ClassName(obj uintptr) string {
    ret, _, _ := trackBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := trackBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TrackBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := trackBar_GetHashCode.Call(obj)
    return int32(ret)
}

func TrackBar_ToString(obj uintptr) string {
    ret, _, _ := trackBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := trackBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func TrackBar_SetAlign(obj uintptr, value TAlign) {
   trackBar_SetAlign.Call(obj, uintptr(value))
}

func TrackBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := trackBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TrackBar_SetAnchors(obj uintptr, value TAnchors) {
   trackBar_SetAnchors.Call(obj, uintptr(value))
}

func TrackBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetBorderWidth(obj uintptr, value int32) {
   trackBar_SetBorderWidth.Call(obj, uintptr(value))
}

func TrackBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := trackBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetDoubleBuffered(obj uintptr, value bool) {
   trackBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := trackBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetEnabled(obj uintptr, value bool) {
   trackBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetOrientation(obj uintptr) TTrackBarOrientation {
    ret, _, _ := trackBar_GetOrientation.Call(obj)
    return TTrackBarOrientation(ret)
}

func TrackBar_SetOrientation(obj uintptr, value TTrackBarOrientation) {
   trackBar_SetOrientation.Call(obj, uintptr(value))
}

func TrackBar_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := trackBar_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetParentCtl3D(obj uintptr, value bool) {
   trackBar_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := trackBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetParentShowHint(obj uintptr, value bool) {
   trackBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetPopupMenu.Call(obj)
    return ret
}

func TrackBar_SetPopupMenu(obj uintptr, value uintptr) {
   trackBar_SetPopupMenu.Call(obj, value)
}

func TrackBar_GetPosition(obj uintptr) int32 {
    ret, _, _ := trackBar_GetPosition.Call(obj)
    return int32(ret)
}

func TrackBar_SetPosition(obj uintptr, value int32) {
   trackBar_SetPosition.Call(obj, uintptr(value))
}

func TrackBar_GetSelStart(obj uintptr) int32 {
    ret, _, _ := trackBar_GetSelStart.Call(obj)
    return int32(ret)
}

func TrackBar_SetSelStart(obj uintptr, value int32) {
   trackBar_SetSelStart.Call(obj, uintptr(value))
}

func TrackBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := trackBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetShowHint(obj uintptr, value bool) {
   trackBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := trackBar_GetTabOrder.Call(obj)
    return int16(ret)
}

func TrackBar_SetTabOrder(obj uintptr, value int16) {
   trackBar_SetTabOrder.Call(obj, uintptr(value))
}

func TrackBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := trackBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetTabStop(obj uintptr, value bool) {
   trackBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetVisible(obj uintptr) bool {
    ret, _, _ := trackBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetVisible(obj uintptr, value bool) {
   trackBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TrackBar_SetOnChange(obj uintptr, fn interface{}) {
    trackBar_SetOnChange.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnEnter(obj uintptr, fn interface{}) {
    trackBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnExit(obj uintptr, fn interface{}) {
    trackBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyDown(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyPress(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyUp(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func TrackBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetBrush.Call(obj)
    return ret
}

func TrackBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := trackBar_GetControlCount.Call(obj)
    return int32(ret)
}

func TrackBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := trackBar_GetHandle.Call(obj)
    return HWND(ret)
}

func TrackBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetAction.Call(obj)
    return ret
}

func TrackBar_SetAction(obj uintptr, value uintptr) {
   trackBar_SetAction.Call(obj, value)
}

func TrackBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    trackBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TrackBar_SetBoundsRect(obj uintptr, value TRect) {
   trackBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TrackBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func TrackBar_SetClientHeight(obj uintptr, value int32) {
   trackBar_SetClientHeight.Call(obj, uintptr(value))
}

func TrackBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    trackBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TrackBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetClientWidth(obj uintptr, value int32) {
   trackBar_SetClientWidth.Call(obj, uintptr(value))
}

func TrackBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TrackBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetParent.Call(obj)
    return ret
}

func TrackBar_SetParent(obj uintptr, value uintptr) {
   trackBar_SetParent.Call(obj, value)
}

func TrackBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := trackBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetAlignWithMargins(obj uintptr, value bool) {
   trackBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := trackBar_GetLeft.Call(obj)
    return int32(ret)
}

func TrackBar_SetLeft(obj uintptr, value int32) {
   trackBar_SetLeft.Call(obj, uintptr(value))
}

func TrackBar_GetTop(obj uintptr) int32 {
    ret, _, _ := trackBar_GetTop.Call(obj)
    return int32(ret)
}

func TrackBar_SetTop(obj uintptr, value int32) {
   trackBar_SetTop.Call(obj, uintptr(value))
}

func TrackBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetWidth(obj uintptr, value int32) {
   trackBar_SetWidth.Call(obj, uintptr(value))
}

func TrackBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetHeight.Call(obj)
    return int32(ret)
}

func TrackBar_SetHeight(obj uintptr, value int32) {
   trackBar_SetHeight.Call(obj, uintptr(value))
}

func TrackBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := trackBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func TrackBar_SetCursor(obj uintptr, value TCursor) {
   trackBar_SetCursor.Call(obj, uintptr(value))
}

func TrackBar_GetHint(obj uintptr) string {
    ret, _, _ := trackBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_SetHint(obj uintptr, value string) {
   trackBar_SetHint.Call(obj, GoStrToDStr(value))
}

func TrackBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetMargins.Call(obj)
    return ret
}

func TrackBar_SetMargins(obj uintptr, value uintptr) {
   trackBar_SetMargins.Call(obj, value)
}

func TrackBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := trackBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func TrackBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := trackBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TrackBar_SetComponentIndex(obj uintptr, value int32) {
   trackBar_SetComponentIndex.Call(obj, uintptr(value))
}

func TrackBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetOwner.Call(obj)
    return ret
}

func TrackBar_GetName(obj uintptr) string {
    ret, _, _ := trackBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_SetName(obj uintptr, value string) {
   trackBar_SetName.Call(obj, GoStrToDStr(value))
}

func TrackBar_GetTag(obj uintptr) int {
    ret, _, _ := trackBar_GetTag.Call(obj)
    return int(ret)
}

func TrackBar_SetTag(obj uintptr, value int) {
   trackBar_SetTag.Call(obj, uintptr(value))
}

func TrackBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := trackBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TrackBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := trackBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

