
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Button_Create(obj uintptr) uintptr {
    ret, _, _ := button_Create.Call(obj)
    return ret
}

func Button_Free(obj uintptr) {
    button_Free.Call(obj)
}

func Button_Click(obj uintptr)  {
    button_Click.Call(obj)
}

func Button_CanFocus(obj uintptr) bool {
    ret, _, _ := button_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_FlipChildren(obj uintptr, AllLevels bool)  {
    button_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Button_Focused(obj uintptr) bool {
    ret, _, _ := button_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_HandleAllocated(obj uintptr) bool {
    ret, _, _ := button_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_Invalidate(obj uintptr)  {
    button_Invalidate.Call(obj)
}

func Button_Realign(obj uintptr)  {
    button_Realign.Call(obj)
}

func Button_Repaint(obj uintptr)  {
    button_Repaint.Call(obj)
}

func Button_ScaleBy(obj uintptr, M int32, D int32)  {
    button_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Button_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    button_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Button_SetFocus(obj uintptr)  {
    button_SetFocus.Call(obj)
}

func Button_Update(obj uintptr)  {
    button_Update.Call(obj)
}

func Button_BringToFront(obj uintptr)  {
    button_BringToFront.Call(obj)
}

func Button_HasParent(obj uintptr) bool {
    ret, _, _ := button_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_Hide(obj uintptr)  {
    button_Hide.Call(obj)
}

func Button_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := button_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Button_Refresh(obj uintptr)  {
    button_Refresh.Call(obj)
}

func Button_SendToBack(obj uintptr)  {
    button_SendToBack.Call(obj)
}

func Button_Show(obj uintptr)  {
    button_Show.Call(obj)
}

func Button_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := button_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Button_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := button_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Button_GetNamePath(obj uintptr) string {
    ret, _, _ := button_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Button_Assign(obj uintptr, Source uintptr)  {
    button_Assign.Call(obj, Source )
}

func Button_ClassName(obj uintptr) string {
    ret, _, _ := button_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Button_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := button_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Button_GetHashCode(obj uintptr) int32 {
    ret, _, _ := button_GetHashCode.Call(obj)
    return int32(ret)
}

func Button_ToString(obj uintptr) string {
    ret, _, _ := button_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Button_GetAction(obj uintptr) uintptr {
    ret, _, _ := button_GetAction.Call(obj)
    return ret
}

func Button_SetAction(obj uintptr, value uintptr) {
   button_SetAction.Call(obj, value)
}

func Button_GetAlign(obj uintptr) TAlign {
    ret, _, _ := button_GetAlign.Call(obj)
    return TAlign(ret)
}

func Button_SetAlign(obj uintptr, value TAlign) {
   button_SetAlign.Call(obj, uintptr(value))
}

func Button_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := button_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Button_SetAnchors(obj uintptr, value TAnchors) {
   button_SetAnchors.Call(obj, uintptr(value))
}

func Button_GetCancel(obj uintptr) bool {
    ret, _, _ := button_GetCancel.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetCancel(obj uintptr, value bool) {
   button_SetCancel.Call(obj, GoBoolToDBool(value))
}

func Button_GetCaption(obj uintptr) string {
    ret, _, _ := button_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetCaption(obj uintptr, value string) {
   button_SetCaption.Call(obj, GoStrToDStr(value))
}

func Button_GetDefault(obj uintptr) bool {
    ret, _, _ := button_GetDefault.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetDefault(obj uintptr, value bool) {
   button_SetDefault.Call(obj, GoBoolToDBool(value))
}

func Button_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := button_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetDoubleBuffered(obj uintptr, value bool) {
   button_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Button_GetEnabled(obj uintptr) bool {
    ret, _, _ := button_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetEnabled(obj uintptr, value bool) {
   button_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Button_GetFont(obj uintptr) uintptr {
    ret, _, _ := button_GetFont.Call(obj)
    return ret
}

func Button_SetFont(obj uintptr, value uintptr) {
   button_SetFont.Call(obj, value)
}

func Button_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetImageIndex(obj uintptr, value int32) {
   button_SetImageIndex.Call(obj, uintptr(value))
}

func Button_GetImages(obj uintptr) uintptr {
    ret, _, _ := button_GetImages.Call(obj)
    return ret
}

func Button_SetImages(obj uintptr, value uintptr) {
   button_SetImages.Call(obj, value)
}

func Button_GetModalResult(obj uintptr) TModalResult {
    ret, _, _ := button_GetModalResult.Call(obj)
    return TModalResult(ret)
}

func Button_SetModalResult(obj uintptr, value TModalResult) {
   button_SetModalResult.Call(obj, uintptr(value))
}

func Button_GetParentFont(obj uintptr) bool {
    ret, _, _ := button_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetParentFont(obj uintptr, value bool) {
   button_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Button_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := button_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetParentShowHint(obj uintptr, value bool) {
   button_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Button_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := button_GetPopupMenu.Call(obj)
    return ret
}

func Button_SetPopupMenu(obj uintptr, value uintptr) {
   button_SetPopupMenu.Call(obj, value)
}

func Button_GetShowHint(obj uintptr) bool {
    ret, _, _ := button_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetShowHint(obj uintptr, value bool) {
   button_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Button_GetStyle(obj uintptr) TButtonStyle {
    ret, _, _ := button_GetStyle.Call(obj)
    return TButtonStyle(ret)
}

func Button_SetStyle(obj uintptr, value TButtonStyle) {
   button_SetStyle.Call(obj, uintptr(value))
}

func Button_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := button_GetTabOrder.Call(obj)
    return int16(ret)
}

func Button_SetTabOrder(obj uintptr, value int16) {
   button_SetTabOrder.Call(obj, uintptr(value))
}

func Button_GetTabStop(obj uintptr) bool {
    ret, _, _ := button_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetTabStop(obj uintptr, value bool) {
   button_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Button_GetVisible(obj uintptr) bool {
    ret, _, _ := button_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetVisible(obj uintptr, value bool) {
   button_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Button_GetWordWrap(obj uintptr) bool {
    ret, _, _ := button_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetWordWrap(obj uintptr, value bool) {
   button_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Button_SetOnClick(obj uintptr, fn interface{}) {
    button_SetOnClick.Call(obj, addEventToMap(fn))
}

func Button_SetOnEnter(obj uintptr, fn interface{}) {
    button_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Button_SetOnExit(obj uintptr, fn interface{}) {
    button_SetOnExit.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyDown(obj uintptr, fn interface{}) {
    button_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyPress(obj uintptr, fn interface{}) {
    button_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyUp(obj uintptr, fn interface{}) {
    button_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseDown(obj uintptr, fn interface{}) {
    button_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseEnter(obj uintptr, fn interface{}) {
    button_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseLeave(obj uintptr, fn interface{}) {
    button_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseMove(obj uintptr, fn interface{}) {
    button_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseUp(obj uintptr, fn interface{}) {
    button_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Button_GetBrush(obj uintptr) uintptr {
    ret, _, _ := button_GetBrush.Call(obj)
    return ret
}

func Button_GetControlCount(obj uintptr) int32 {
    ret, _, _ := button_GetControlCount.Call(obj)
    return int32(ret)
}

func Button_GetHandle(obj uintptr) HWND {
    ret, _, _ := button_GetHandle.Call(obj)
    return HWND(ret)
}

func Button_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    button_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Button_SetBoundsRect(obj uintptr, value TRect) {
   button_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Button_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := button_GetClientHeight.Call(obj)
    return int32(ret)
}

func Button_SetClientHeight(obj uintptr, value int32) {
   button_SetClientHeight.Call(obj, uintptr(value))
}

func Button_GetClientRect(obj uintptr) TRect {
    var ret TRect
    button_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Button_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := button_GetClientWidth.Call(obj)
    return int32(ret)
}

func Button_SetClientWidth(obj uintptr, value int32) {
   button_SetClientWidth.Call(obj, uintptr(value))
}

func Button_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Button_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Button_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Button_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Button_GetParent(obj uintptr) uintptr {
    ret, _, _ := button_GetParent.Call(obj)
    return ret
}

func Button_SetParent(obj uintptr, value uintptr) {
   button_SetParent.Call(obj, value)
}

func Button_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := button_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetAlignWithMargins(obj uintptr, value bool) {
   button_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Button_GetLeft(obj uintptr) int32 {
    ret, _, _ := button_GetLeft.Call(obj)
    return int32(ret)
}

func Button_SetLeft(obj uintptr, value int32) {
   button_SetLeft.Call(obj, uintptr(value))
}

func Button_GetTop(obj uintptr) int32 {
    ret, _, _ := button_GetTop.Call(obj)
    return int32(ret)
}

func Button_SetTop(obj uintptr, value int32) {
   button_SetTop.Call(obj, uintptr(value))
}

func Button_GetWidth(obj uintptr) int32 {
    ret, _, _ := button_GetWidth.Call(obj)
    return int32(ret)
}

func Button_SetWidth(obj uintptr, value int32) {
   button_SetWidth.Call(obj, uintptr(value))
}

func Button_GetHeight(obj uintptr) int32 {
    ret, _, _ := button_GetHeight.Call(obj)
    return int32(ret)
}

func Button_SetHeight(obj uintptr, value int32) {
   button_SetHeight.Call(obj, uintptr(value))
}

func Button_GetCursor(obj uintptr) TCursor {
    ret, _, _ := button_GetCursor.Call(obj)
    return TCursor(ret)
}

func Button_SetCursor(obj uintptr, value TCursor) {
   button_SetCursor.Call(obj, uintptr(value))
}

func Button_GetHint(obj uintptr) string {
    ret, _, _ := button_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetHint(obj uintptr, value string) {
   button_SetHint.Call(obj, GoStrToDStr(value))
}

func Button_GetMargins(obj uintptr) uintptr {
    ret, _, _ := button_GetMargins.Call(obj)
    return ret
}

func Button_SetMargins(obj uintptr, value uintptr) {
   button_SetMargins.Call(obj, value)
}

func Button_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := button_GetComponentCount.Call(obj)
    return int32(ret)
}

func Button_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := button_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Button_SetComponentIndex(obj uintptr, value int32) {
   button_SetComponentIndex.Call(obj, uintptr(value))
}

func Button_GetOwner(obj uintptr) uintptr {
    ret, _, _ := button_GetOwner.Call(obj)
    return ret
}

func Button_GetName(obj uintptr) string {
    ret, _, _ := button_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetName(obj uintptr, value string) {
   button_SetName.Call(obj, GoStrToDStr(value))
}

func Button_GetTag(obj uintptr) int {
    ret, _, _ := button_GetTag.Call(obj)
    return int(ret)
}

func Button_SetTag(obj uintptr, value int) {
   button_SetTag.Call(obj, uintptr(value))
}

func Button_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := button_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Button_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := button_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

