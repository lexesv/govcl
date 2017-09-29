
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func RadioButton_Create(obj uintptr) uintptr {
    ret, _, _ := radioButton_Create.Call(obj)
    return ret
}

func RadioButton_Free(obj uintptr) {
    radioButton_Free.Call(obj)
}

func RadioButton_CanFocus(obj uintptr) bool {
    ret, _, _ := radioButton_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_FlipChildren(obj uintptr, AllLevels bool)  {
    radioButton_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RadioButton_Focused(obj uintptr) bool {
    ret, _, _ := radioButton_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_HandleAllocated(obj uintptr) bool {
    ret, _, _ := radioButton_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_Invalidate(obj uintptr)  {
    radioButton_Invalidate.Call(obj)
}

func RadioButton_Realign(obj uintptr)  {
    radioButton_Realign.Call(obj)
}

func RadioButton_Repaint(obj uintptr)  {
    radioButton_Repaint.Call(obj)
}

func RadioButton_ScaleBy(obj uintptr, M int32, D int32)  {
    radioButton_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RadioButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    radioButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RadioButton_SetFocus(obj uintptr)  {
    radioButton_SetFocus.Call(obj)
}

func RadioButton_Update(obj uintptr)  {
    radioButton_Update.Call(obj)
}

func RadioButton_BringToFront(obj uintptr)  {
    radioButton_BringToFront.Call(obj)
}

func RadioButton_HasParent(obj uintptr) bool {
    ret, _, _ := radioButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_Hide(obj uintptr)  {
    radioButton_Hide.Call(obj)
}

func RadioButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := radioButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RadioButton_Refresh(obj uintptr)  {
    radioButton_Refresh.Call(obj)
}

func RadioButton_SendToBack(obj uintptr)  {
    radioButton_SendToBack.Call(obj)
}

func RadioButton_Show(obj uintptr)  {
    radioButton_Show.Call(obj)
}

func RadioButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := radioButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RadioButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := radioButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RadioButton_GetNamePath(obj uintptr) string {
    ret, _, _ := radioButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_Assign(obj uintptr, Source uintptr)  {
    radioButton_Assign.Call(obj, Source )
}

func RadioButton_ClassName(obj uintptr) string {
    ret, _, _ := radioButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := radioButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RadioButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := radioButton_GetHashCode.Call(obj)
    return int32(ret)
}

func RadioButton_ToString(obj uintptr) string {
    ret, _, _ := radioButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetAction.Call(obj)
    return ret
}

func RadioButton_SetAction(obj uintptr, value uintptr) {
   radioButton_SetAction.Call(obj, value)
}

func RadioButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := radioButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func RadioButton_SetAlign(obj uintptr, value TAlign) {
   radioButton_SetAlign.Call(obj, uintptr(value))
}

func RadioButton_GetAlignment(obj uintptr) TLeftRight {
    ret, _, _ := radioButton_GetAlignment.Call(obj)
    return TLeftRight(ret)
}

func RadioButton_SetAlignment(obj uintptr, value TLeftRight) {
   radioButton_SetAlignment.Call(obj, uintptr(value))
}

func RadioButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := radioButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RadioButton_SetAnchors(obj uintptr, value TAnchors) {
   radioButton_SetAnchors.Call(obj, uintptr(value))
}

func RadioButton_GetCaption(obj uintptr) string {
    ret, _, _ := radioButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetCaption(obj uintptr, value string) {
   radioButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetChecked(obj uintptr) bool {
    ret, _, _ := radioButton_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetChecked(obj uintptr, value bool) {
   radioButton_SetChecked.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetColor(obj uintptr) TColor {
    ret, _, _ := radioButton_GetColor.Call(obj)
    return TColor(ret)
}

func RadioButton_SetColor(obj uintptr, value TColor) {
   radioButton_SetColor.Call(obj, uintptr(value))
}

func RadioButton_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioButton_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetDoubleBuffered(obj uintptr, value bool) {
   radioButton_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := radioButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetEnabled(obj uintptr, value bool) {
   radioButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetFont(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetFont.Call(obj)
    return ret
}

func RadioButton_SetFont(obj uintptr, value uintptr) {
   radioButton_SetFont.Call(obj, value)
}

func RadioButton_GetParentColor(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentColor(obj uintptr, value bool) {
   radioButton_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentCtl3D(obj uintptr, value bool) {
   radioButton_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentFont(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentFont(obj uintptr, value bool) {
   radioButton_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentShowHint(obj uintptr, value bool) {
   radioButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetPopupMenu.Call(obj)
    return ret
}

func RadioButton_SetPopupMenu(obj uintptr, value uintptr) {
   radioButton_SetPopupMenu.Call(obj, value)
}

func RadioButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := radioButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetShowHint(obj uintptr, value bool) {
   radioButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := radioButton_GetTabOrder.Call(obj)
    return int16(ret)
}

func RadioButton_SetTabOrder(obj uintptr, value int16) {
   radioButton_SetTabOrder.Call(obj, uintptr(value))
}

func RadioButton_GetTabStop(obj uintptr) bool {
    ret, _, _ := radioButton_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetTabStop(obj uintptr, value bool) {
   radioButton_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetVisible(obj uintptr) bool {
    ret, _, _ := radioButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetVisible(obj uintptr, value bool) {
   radioButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetWordWrap(obj uintptr) bool {
    ret, _, _ := radioButton_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetWordWrap(obj uintptr, value bool) {
   radioButton_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RadioButton_SetOnClick(obj uintptr, fn interface{}) {
    radioButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnDblClick(obj uintptr, fn interface{}) {
    radioButton_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnEnter(obj uintptr, fn interface{}) {
    radioButton_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnExit(obj uintptr, fn interface{}) {
    radioButton_SetOnExit.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyDown(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyPress(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyUp(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func RadioButton_GetBrush(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetBrush.Call(obj)
    return ret
}

func RadioButton_GetControlCount(obj uintptr) int32 {
    ret, _, _ := radioButton_GetControlCount.Call(obj)
    return int32(ret)
}

func RadioButton_GetHandle(obj uintptr) HWND {
    ret, _, _ := radioButton_GetHandle.Call(obj)
    return HWND(ret)
}

func RadioButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    radioButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioButton_SetBoundsRect(obj uintptr, value TRect) {
   radioButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func RadioButton_SetClientHeight(obj uintptr, value int32) {
   radioButton_SetClientHeight.Call(obj, uintptr(value))
}

func RadioButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    radioButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func RadioButton_SetClientWidth(obj uintptr, value int32) {
   radioButton_SetClientWidth.Call(obj, uintptr(value))
}

func RadioButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RadioButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetParent.Call(obj)
    return ret
}

func RadioButton_SetParent(obj uintptr, value uintptr) {
   radioButton_SetParent.Call(obj, value)
}

func RadioButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := radioButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetAlignWithMargins(obj uintptr, value bool) {
   radioButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := radioButton_GetLeft.Call(obj)
    return int32(ret)
}

func RadioButton_SetLeft(obj uintptr, value int32) {
   radioButton_SetLeft.Call(obj, uintptr(value))
}

func RadioButton_GetTop(obj uintptr) int32 {
    ret, _, _ := radioButton_GetTop.Call(obj)
    return int32(ret)
}

func RadioButton_SetTop(obj uintptr, value int32) {
   radioButton_SetTop.Call(obj, uintptr(value))
}

func RadioButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetWidth.Call(obj)
    return int32(ret)
}

func RadioButton_SetWidth(obj uintptr, value int32) {
   radioButton_SetWidth.Call(obj, uintptr(value))
}

func RadioButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetHeight.Call(obj)
    return int32(ret)
}

func RadioButton_SetHeight(obj uintptr, value int32) {
   radioButton_SetHeight.Call(obj, uintptr(value))
}

func RadioButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := radioButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func RadioButton_SetCursor(obj uintptr, value TCursor) {
   radioButton_SetCursor.Call(obj, uintptr(value))
}

func RadioButton_GetHint(obj uintptr) string {
    ret, _, _ := radioButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetHint(obj uintptr, value string) {
   radioButton_SetHint.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetMargins.Call(obj)
    return ret
}

func RadioButton_SetMargins(obj uintptr, value uintptr) {
   radioButton_SetMargins.Call(obj, value)
}

func RadioButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := radioButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func RadioButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := radioButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RadioButton_SetComponentIndex(obj uintptr, value int32) {
   radioButton_SetComponentIndex.Call(obj, uintptr(value))
}

func RadioButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetOwner.Call(obj)
    return ret
}

func RadioButton_GetName(obj uintptr) string {
    ret, _, _ := radioButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetName(obj uintptr, value string) {
   radioButton_SetName.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetTag(obj uintptr) int {
    ret, _, _ := radioButton_GetTag.Call(obj)
    return int(ret)
}

func RadioButton_SetTag(obj uintptr, value int) {
   radioButton_SetTag.Call(obj, uintptr(value))
}

func RadioButton_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioButton_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RadioButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := radioButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

