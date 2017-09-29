
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func UpDown_Create(obj uintptr) uintptr {
    ret, _, _ := upDown_Create.Call(obj)
    return ret
}

func UpDown_Free(obj uintptr) {
    upDown_Free.Call(obj)
}

func UpDown_CanFocus(obj uintptr) bool {
    ret, _, _ := upDown_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_FlipChildren(obj uintptr, AllLevels bool)  {
    upDown_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func UpDown_Focused(obj uintptr) bool {
    ret, _, _ := upDown_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_HandleAllocated(obj uintptr) bool {
    ret, _, _ := upDown_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_Invalidate(obj uintptr)  {
    upDown_Invalidate.Call(obj)
}

func UpDown_Realign(obj uintptr)  {
    upDown_Realign.Call(obj)
}

func UpDown_Repaint(obj uintptr)  {
    upDown_Repaint.Call(obj)
}

func UpDown_ScaleBy(obj uintptr, M int32, D int32)  {
    upDown_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func UpDown_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    upDown_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func UpDown_SetFocus(obj uintptr)  {
    upDown_SetFocus.Call(obj)
}

func UpDown_Update(obj uintptr)  {
    upDown_Update.Call(obj)
}

func UpDown_BringToFront(obj uintptr)  {
    upDown_BringToFront.Call(obj)
}

func UpDown_HasParent(obj uintptr) bool {
    ret, _, _ := upDown_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_Hide(obj uintptr)  {
    upDown_Hide.Call(obj)
}

func UpDown_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := upDown_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func UpDown_Refresh(obj uintptr)  {
    upDown_Refresh.Call(obj)
}

func UpDown_SendToBack(obj uintptr)  {
    upDown_SendToBack.Call(obj)
}

func UpDown_Show(obj uintptr)  {
    upDown_Show.Call(obj)
}

func UpDown_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := upDown_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func UpDown_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := upDown_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func UpDown_GetNamePath(obj uintptr) string {
    ret, _, _ := upDown_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_Assign(obj uintptr, Source uintptr)  {
    upDown_Assign.Call(obj, Source )
}

func UpDown_ClassName(obj uintptr) string {
    ret, _, _ := upDown_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := upDown_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func UpDown_GetHashCode(obj uintptr) int32 {
    ret, _, _ := upDown_GetHashCode.Call(obj)
    return int32(ret)
}

func UpDown_ToString(obj uintptr) string {
    ret, _, _ := upDown_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := upDown_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func UpDown_SetAnchors(obj uintptr, value TAnchors) {
   upDown_SetAnchors.Call(obj, uintptr(value))
}

func UpDown_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := upDown_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetDoubleBuffered(obj uintptr, value bool) {
   upDown_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetEnabled(obj uintptr) bool {
    ret, _, _ := upDown_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetEnabled(obj uintptr, value bool) {
   upDown_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetHint(obj uintptr) string {
    ret, _, _ := upDown_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_SetHint(obj uintptr, value string) {
   upDown_SetHint.Call(obj, GoStrToDStr(value))
}

func UpDown_GetOrientation(obj uintptr) TUDOrientation {
    ret, _, _ := upDown_GetOrientation.Call(obj)
    return TUDOrientation(ret)
}

func UpDown_SetOrientation(obj uintptr, value TUDOrientation) {
   upDown_SetOrientation.Call(obj, uintptr(value))
}

func UpDown_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := upDown_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetParentShowHint(obj uintptr, value bool) {
   upDown_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := upDown_GetPopupMenu.Call(obj)
    return ret
}

func UpDown_SetPopupMenu(obj uintptr, value uintptr) {
   upDown_SetPopupMenu.Call(obj, value)
}

func UpDown_GetPosition(obj uintptr) int32 {
    ret, _, _ := upDown_GetPosition.Call(obj)
    return int32(ret)
}

func UpDown_SetPosition(obj uintptr, value int32) {
   upDown_SetPosition.Call(obj, uintptr(value))
}

func UpDown_GetShowHint(obj uintptr) bool {
    ret, _, _ := upDown_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetShowHint(obj uintptr, value bool) {
   upDown_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := upDown_GetTabOrder.Call(obj)
    return int16(ret)
}

func UpDown_SetTabOrder(obj uintptr, value int16) {
   upDown_SetTabOrder.Call(obj, uintptr(value))
}

func UpDown_GetTabStop(obj uintptr) bool {
    ret, _, _ := upDown_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetTabStop(obj uintptr, value bool) {
   upDown_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetVisible(obj uintptr) bool {
    ret, _, _ := upDown_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetVisible(obj uintptr, value bool) {
   upDown_SetVisible.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetWrap(obj uintptr) bool {
    ret, _, _ := upDown_GetWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetWrap(obj uintptr, value bool) {
   upDown_SetWrap.Call(obj, GoBoolToDBool(value))
}

func UpDown_SetOnClick(obj uintptr, fn interface{}) {
    upDown_SetOnClick.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnEnter(obj uintptr, fn interface{}) {
    upDown_SetOnEnter.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnExit(obj uintptr, fn interface{}) {
    upDown_SetOnExit.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseDown(obj uintptr, fn interface{}) {
    upDown_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseEnter(obj uintptr, fn interface{}) {
    upDown_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseLeave(obj uintptr, fn interface{}) {
    upDown_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseMove(obj uintptr, fn interface{}) {
    upDown_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseUp(obj uintptr, fn interface{}) {
    upDown_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func UpDown_GetBrush(obj uintptr) uintptr {
    ret, _, _ := upDown_GetBrush.Call(obj)
    return ret
}

func UpDown_GetControlCount(obj uintptr) int32 {
    ret, _, _ := upDown_GetControlCount.Call(obj)
    return int32(ret)
}

func UpDown_GetHandle(obj uintptr) HWND {
    ret, _, _ := upDown_GetHandle.Call(obj)
    return HWND(ret)
}

func UpDown_GetAction(obj uintptr) uintptr {
    ret, _, _ := upDown_GetAction.Call(obj)
    return ret
}

func UpDown_SetAction(obj uintptr, value uintptr) {
   upDown_SetAction.Call(obj, value)
}

func UpDown_GetAlign(obj uintptr) TAlign {
    ret, _, _ := upDown_GetAlign.Call(obj)
    return TAlign(ret)
}

func UpDown_SetAlign(obj uintptr, value TAlign) {
   upDown_SetAlign.Call(obj, uintptr(value))
}

func UpDown_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    upDown_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func UpDown_SetBoundsRect(obj uintptr, value TRect) {
   upDown_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func UpDown_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetClientHeight.Call(obj)
    return int32(ret)
}

func UpDown_SetClientHeight(obj uintptr, value int32) {
   upDown_SetClientHeight.Call(obj, uintptr(value))
}

func UpDown_GetClientRect(obj uintptr) TRect {
    var ret TRect
    upDown_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func UpDown_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetClientWidth.Call(obj)
    return int32(ret)
}

func UpDown_SetClientWidth(obj uintptr, value int32) {
   upDown_SetClientWidth.Call(obj, uintptr(value))
}

func UpDown_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitTop.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func UpDown_GetParent(obj uintptr) uintptr {
    ret, _, _ := upDown_GetParent.Call(obj)
    return ret
}

func UpDown_SetParent(obj uintptr, value uintptr) {
   upDown_SetParent.Call(obj, value)
}

func UpDown_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := upDown_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetAlignWithMargins(obj uintptr, value bool) {
   upDown_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetLeft(obj uintptr) int32 {
    ret, _, _ := upDown_GetLeft.Call(obj)
    return int32(ret)
}

func UpDown_SetLeft(obj uintptr, value int32) {
   upDown_SetLeft.Call(obj, uintptr(value))
}

func UpDown_GetTop(obj uintptr) int32 {
    ret, _, _ := upDown_GetTop.Call(obj)
    return int32(ret)
}

func UpDown_SetTop(obj uintptr, value int32) {
   upDown_SetTop.Call(obj, uintptr(value))
}

func UpDown_GetWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetWidth.Call(obj)
    return int32(ret)
}

func UpDown_SetWidth(obj uintptr, value int32) {
   upDown_SetWidth.Call(obj, uintptr(value))
}

func UpDown_GetHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetHeight.Call(obj)
    return int32(ret)
}

func UpDown_SetHeight(obj uintptr, value int32) {
   upDown_SetHeight.Call(obj, uintptr(value))
}

func UpDown_GetCursor(obj uintptr) TCursor {
    ret, _, _ := upDown_GetCursor.Call(obj)
    return TCursor(ret)
}

func UpDown_SetCursor(obj uintptr, value TCursor) {
   upDown_SetCursor.Call(obj, uintptr(value))
}

func UpDown_GetMargins(obj uintptr) uintptr {
    ret, _, _ := upDown_GetMargins.Call(obj)
    return ret
}

func UpDown_SetMargins(obj uintptr, value uintptr) {
   upDown_SetMargins.Call(obj, value)
}

func UpDown_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := upDown_GetComponentCount.Call(obj)
    return int32(ret)
}

func UpDown_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := upDown_GetComponentIndex.Call(obj)
    return int32(ret)
}

func UpDown_SetComponentIndex(obj uintptr, value int32) {
   upDown_SetComponentIndex.Call(obj, uintptr(value))
}

func UpDown_GetOwner(obj uintptr) uintptr {
    ret, _, _ := upDown_GetOwner.Call(obj)
    return ret
}

func UpDown_GetName(obj uintptr) string {
    ret, _, _ := upDown_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_SetName(obj uintptr, value string) {
   upDown_SetName.Call(obj, GoStrToDStr(value))
}

func UpDown_GetTag(obj uintptr) int {
    ret, _, _ := upDown_GetTag.Call(obj)
    return int(ret)
}

func UpDown_SetTag(obj uintptr, value int) {
   upDown_SetTag.Call(obj, uintptr(value))
}

func UpDown_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := upDown_GetControls.Call(obj, uintptr(Index))
    return ret
}

func UpDown_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := upDown_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

