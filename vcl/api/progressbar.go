
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ProgressBar_Create(obj uintptr) uintptr {
    ret, _, _ := progressBar_Create.Call(obj)
    return ret
}

func ProgressBar_Free(obj uintptr) {
    progressBar_Free.Call(obj)
}

func ProgressBar_CanFocus(obj uintptr) bool {
    ret, _, _ := progressBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_FlipChildren(obj uintptr, AllLevels bool)  {
    progressBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ProgressBar_Focused(obj uintptr) bool {
    ret, _, _ := progressBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := progressBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_Invalidate(obj uintptr)  {
    progressBar_Invalidate.Call(obj)
}

func ProgressBar_Realign(obj uintptr)  {
    progressBar_Realign.Call(obj)
}

func ProgressBar_Repaint(obj uintptr)  {
    progressBar_Repaint.Call(obj)
}

func ProgressBar_ScaleBy(obj uintptr, M int32, D int32)  {
    progressBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ProgressBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    progressBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ProgressBar_SetFocus(obj uintptr)  {
    progressBar_SetFocus.Call(obj)
}

func ProgressBar_Update(obj uintptr)  {
    progressBar_Update.Call(obj)
}

func ProgressBar_BringToFront(obj uintptr)  {
    progressBar_BringToFront.Call(obj)
}

func ProgressBar_HasParent(obj uintptr) bool {
    ret, _, _ := progressBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_Hide(obj uintptr)  {
    progressBar_Hide.Call(obj)
}

func ProgressBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := progressBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ProgressBar_Refresh(obj uintptr)  {
    progressBar_Refresh.Call(obj)
}

func ProgressBar_SendToBack(obj uintptr)  {
    progressBar_SendToBack.Call(obj)
}

func ProgressBar_Show(obj uintptr)  {
    progressBar_Show.Call(obj)
}

func ProgressBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := progressBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ProgressBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := progressBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ProgressBar_GetNamePath(obj uintptr) string {
    ret, _, _ := progressBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_Assign(obj uintptr, Source uintptr)  {
    progressBar_Assign.Call(obj, Source )
}

func ProgressBar_ClassName(obj uintptr) string {
    ret, _, _ := progressBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := progressBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ProgressBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := progressBar_GetHashCode.Call(obj)
    return int32(ret)
}

func ProgressBar_ToString(obj uintptr) string {
    ret, _, _ := progressBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := progressBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func ProgressBar_SetAlign(obj uintptr, value TAlign) {
   progressBar_SetAlign.Call(obj, uintptr(value))
}

func ProgressBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := progressBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ProgressBar_SetAnchors(obj uintptr, value TAnchors) {
   progressBar_SetAnchors.Call(obj, uintptr(value))
}

func ProgressBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetBorderWidth(obj uintptr, value int32) {
   progressBar_SetBorderWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := progressBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetDoubleBuffered(obj uintptr, value bool) {
   progressBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := progressBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetEnabled(obj uintptr, value bool) {
   progressBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetHint(obj uintptr) string {
    ret, _, _ := progressBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_SetHint(obj uintptr, value string) {
   progressBar_SetHint.Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetOrientation(obj uintptr) TProgressBarOrientation {
    ret, _, _ := progressBar_GetOrientation.Call(obj)
    return TProgressBarOrientation(ret)
}

func ProgressBar_SetOrientation(obj uintptr, value TProgressBarOrientation) {
   progressBar_SetOrientation.Call(obj, uintptr(value))
}

func ProgressBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := progressBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetParentShowHint(obj uintptr, value bool) {
   progressBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetPopupMenu.Call(obj)
    return ret
}

func ProgressBar_SetPopupMenu(obj uintptr, value uintptr) {
   progressBar_SetPopupMenu.Call(obj, value)
}

func ProgressBar_GetPosition(obj uintptr) int32 {
    ret, _, _ := progressBar_GetPosition.Call(obj)
    return int32(ret)
}

func ProgressBar_SetPosition(obj uintptr, value int32) {
   progressBar_SetPosition.Call(obj, uintptr(value))
}

func ProgressBar_GetStyle(obj uintptr) TProgressBarStyle {
    ret, _, _ := progressBar_GetStyle.Call(obj)
    return TProgressBarStyle(ret)
}

func ProgressBar_SetStyle(obj uintptr, value TProgressBarStyle) {
   progressBar_SetStyle.Call(obj, uintptr(value))
}

func ProgressBar_GetState(obj uintptr) TProgressBarState {
    ret, _, _ := progressBar_GetState.Call(obj)
    return TProgressBarState(ret)
}

func ProgressBar_SetState(obj uintptr, value TProgressBarState) {
   progressBar_SetState.Call(obj, uintptr(value))
}

func ProgressBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := progressBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetShowHint(obj uintptr, value bool) {
   progressBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := progressBar_GetTabOrder.Call(obj)
    return int16(ret)
}

func ProgressBar_SetTabOrder(obj uintptr, value int16) {
   progressBar_SetTabOrder.Call(obj, uintptr(value))
}

func ProgressBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := progressBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetTabStop(obj uintptr, value bool) {
   progressBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetVisible(obj uintptr) bool {
    ret, _, _ := progressBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetVisible(obj uintptr, value bool) {
   progressBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_SetOnEnter(obj uintptr, fn interface{}) {
    progressBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnExit(obj uintptr, fn interface{}) {
    progressBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ProgressBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetBrush.Call(obj)
    return ret
}

func ProgressBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := progressBar_GetControlCount.Call(obj)
    return int32(ret)
}

func ProgressBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := progressBar_GetHandle.Call(obj)
    return HWND(ret)
}

func ProgressBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetAction.Call(obj)
    return ret
}

func ProgressBar_SetAction(obj uintptr, value uintptr) {
   progressBar_SetAction.Call(obj, value)
}

func ProgressBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    progressBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ProgressBar_SetBoundsRect(obj uintptr, value TRect) {
   progressBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ProgressBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_SetClientHeight(obj uintptr, value int32) {
   progressBar_SetClientHeight.Call(obj, uintptr(value))
}

func ProgressBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    progressBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ProgressBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetClientWidth(obj uintptr, value int32) {
   progressBar_SetClientWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetParent.Call(obj)
    return ret
}

func ProgressBar_SetParent(obj uintptr, value uintptr) {
   progressBar_SetParent.Call(obj, value)
}

func ProgressBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := progressBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetAlignWithMargins(obj uintptr, value bool) {
   progressBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := progressBar_GetLeft.Call(obj)
    return int32(ret)
}

func ProgressBar_SetLeft(obj uintptr, value int32) {
   progressBar_SetLeft.Call(obj, uintptr(value))
}

func ProgressBar_GetTop(obj uintptr) int32 {
    ret, _, _ := progressBar_GetTop.Call(obj)
    return int32(ret)
}

func ProgressBar_SetTop(obj uintptr, value int32) {
   progressBar_SetTop.Call(obj, uintptr(value))
}

func ProgressBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetWidth(obj uintptr, value int32) {
   progressBar_SetWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_SetHeight(obj uintptr, value int32) {
   progressBar_SetHeight.Call(obj, uintptr(value))
}

func ProgressBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := progressBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func ProgressBar_SetCursor(obj uintptr, value TCursor) {
   progressBar_SetCursor.Call(obj, uintptr(value))
}

func ProgressBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetMargins.Call(obj)
    return ret
}

func ProgressBar_SetMargins(obj uintptr, value uintptr) {
   progressBar_SetMargins.Call(obj, value)
}

func ProgressBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := progressBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func ProgressBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := progressBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ProgressBar_SetComponentIndex(obj uintptr, value int32) {
   progressBar_SetComponentIndex.Call(obj, uintptr(value))
}

func ProgressBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetOwner.Call(obj)
    return ret
}

func ProgressBar_GetName(obj uintptr) string {
    ret, _, _ := progressBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_SetName(obj uintptr, value string) {
   progressBar_SetName.Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetTag(obj uintptr) int {
    ret, _, _ := progressBar_GetTag.Call(obj)
    return int(ret)
}

func ProgressBar_SetTag(obj uintptr, value int) {
   progressBar_SetTag.Call(obj, uintptr(value))
}

func ProgressBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := progressBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ProgressBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := progressBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

