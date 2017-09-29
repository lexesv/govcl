
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func CheckBox_Create(obj uintptr) uintptr {
    ret, _, _ := checkBox_Create.Call(obj)
    return ret
}

func CheckBox_Free(obj uintptr) {
    checkBox_Free.Call(obj)
}

func CheckBox_CanFocus(obj uintptr) bool {
    ret, _, _ := checkBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_FlipChildren(obj uintptr, AllLevels bool)  {
    checkBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func CheckBox_Focused(obj uintptr) bool {
    ret, _, _ := checkBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := checkBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_Invalidate(obj uintptr)  {
    checkBox_Invalidate.Call(obj)
}

func CheckBox_Realign(obj uintptr)  {
    checkBox_Realign.Call(obj)
}

func CheckBox_Repaint(obj uintptr)  {
    checkBox_Repaint.Call(obj)
}

func CheckBox_ScaleBy(obj uintptr, M int32, D int32)  {
    checkBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func CheckBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    checkBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func CheckBox_SetFocus(obj uintptr)  {
    checkBox_SetFocus.Call(obj)
}

func CheckBox_Update(obj uintptr)  {
    checkBox_Update.Call(obj)
}

func CheckBox_BringToFront(obj uintptr)  {
    checkBox_BringToFront.Call(obj)
}

func CheckBox_HasParent(obj uintptr) bool {
    ret, _, _ := checkBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_Hide(obj uintptr)  {
    checkBox_Hide.Call(obj)
}

func CheckBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := checkBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func CheckBox_Refresh(obj uintptr)  {
    checkBox_Refresh.Call(obj)
}

func CheckBox_SendToBack(obj uintptr)  {
    checkBox_SendToBack.Call(obj)
}

func CheckBox_Show(obj uintptr)  {
    checkBox_Show.Call(obj)
}

func CheckBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := checkBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func CheckBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := checkBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func CheckBox_GetNamePath(obj uintptr) string {
    ret, _, _ := checkBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_Assign(obj uintptr, Source uintptr)  {
    checkBox_Assign.Call(obj, Source )
}

func CheckBox_ClassName(obj uintptr) string {
    ret, _, _ := checkBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := checkBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CheckBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := checkBox_GetHashCode.Call(obj)
    return int32(ret)
}

func CheckBox_ToString(obj uintptr) string {
    ret, _, _ := checkBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetAction.Call(obj)
    return ret
}

func CheckBox_SetAction(obj uintptr, value uintptr) {
   checkBox_SetAction.Call(obj, value)
}

func CheckBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := checkBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func CheckBox_SetAlign(obj uintptr, value TAlign) {
   checkBox_SetAlign.Call(obj, uintptr(value))
}

func CheckBox_GetAlignment(obj uintptr) TLeftRight {
    ret, _, _ := checkBox_GetAlignment.Call(obj)
    return TLeftRight(ret)
}

func CheckBox_SetAlignment(obj uintptr, value TLeftRight) {
   checkBox_SetAlignment.Call(obj, uintptr(value))
}

func CheckBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := checkBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func CheckBox_SetAnchors(obj uintptr, value TAnchors) {
   checkBox_SetAnchors.Call(obj, uintptr(value))
}

func CheckBox_GetCaption(obj uintptr) string {
    ret, _, _ := checkBox_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetCaption(obj uintptr, value string) {
   checkBox_SetCaption.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetChecked(obj uintptr) bool {
    ret, _, _ := checkBox_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetChecked(obj uintptr, value bool) {
   checkBox_SetChecked.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetColor(obj uintptr) TColor {
    ret, _, _ := checkBox_GetColor.Call(obj)
    return TColor(ret)
}

func CheckBox_SetColor(obj uintptr, value TColor) {
   checkBox_SetColor.Call(obj, uintptr(value))
}

func CheckBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := checkBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetDoubleBuffered(obj uintptr, value bool) {
   checkBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := checkBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetEnabled(obj uintptr, value bool) {
   checkBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetFont.Call(obj)
    return ret
}

func CheckBox_SetFont(obj uintptr, value uintptr) {
   checkBox_SetFont.Call(obj, value)
}

func CheckBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentColor(obj uintptr, value bool) {
   checkBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentCtl3D(obj uintptr, value bool) {
   checkBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentFont(obj uintptr, value bool) {
   checkBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentShowHint(obj uintptr, value bool) {
   checkBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetPopupMenu.Call(obj)
    return ret
}

func CheckBox_SetPopupMenu(obj uintptr, value uintptr) {
   checkBox_SetPopupMenu.Call(obj, value)
}

func CheckBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := checkBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetShowHint(obj uintptr, value bool) {
   checkBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetState(obj uintptr) TCheckBoxState {
    ret, _, _ := checkBox_GetState.Call(obj)
    return TCheckBoxState(ret)
}

func CheckBox_SetState(obj uintptr, value TCheckBoxState) {
   checkBox_SetState.Call(obj, uintptr(value))
}

func CheckBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := checkBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func CheckBox_SetTabOrder(obj uintptr, value int16) {
   checkBox_SetTabOrder.Call(obj, uintptr(value))
}

func CheckBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := checkBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetTabStop(obj uintptr, value bool) {
   checkBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetVisible(obj uintptr) bool {
    ret, _, _ := checkBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetVisible(obj uintptr, value bool) {
   checkBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetWordWrap(obj uintptr) bool {
    ret, _, _ := checkBox_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetWordWrap(obj uintptr, value bool) {
   checkBox_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func CheckBox_SetOnClick(obj uintptr, fn interface{}) {
    checkBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnEnter(obj uintptr, fn interface{}) {
    checkBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnExit(obj uintptr, fn interface{}) {
    checkBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func CheckBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetBrush.Call(obj)
    return ret
}

func CheckBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := checkBox_GetControlCount.Call(obj)
    return int32(ret)
}

func CheckBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := checkBox_GetHandle.Call(obj)
    return HWND(ret)
}

func CheckBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    checkBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CheckBox_SetBoundsRect(obj uintptr, value TRect) {
   checkBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func CheckBox_SetClientHeight(obj uintptr, value int32) {
   checkBox_SetClientHeight.Call(obj, uintptr(value))
}

func CheckBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    checkBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CheckBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func CheckBox_SetClientWidth(obj uintptr, value int32) {
   checkBox_SetClientWidth.Call(obj, uintptr(value))
}

func CheckBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func CheckBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetParent.Call(obj)
    return ret
}

func CheckBox_SetParent(obj uintptr, value uintptr) {
   checkBox_SetParent.Call(obj, value)
}

func CheckBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := checkBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetAlignWithMargins(obj uintptr, value bool) {
   checkBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := checkBox_GetLeft.Call(obj)
    return int32(ret)
}

func CheckBox_SetLeft(obj uintptr, value int32) {
   checkBox_SetLeft.Call(obj, uintptr(value))
}

func CheckBox_GetTop(obj uintptr) int32 {
    ret, _, _ := checkBox_GetTop.Call(obj)
    return int32(ret)
}

func CheckBox_SetTop(obj uintptr, value int32) {
   checkBox_SetTop.Call(obj, uintptr(value))
}

func CheckBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetWidth.Call(obj)
    return int32(ret)
}

func CheckBox_SetWidth(obj uintptr, value int32) {
   checkBox_SetWidth.Call(obj, uintptr(value))
}

func CheckBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetHeight.Call(obj)
    return int32(ret)
}

func CheckBox_SetHeight(obj uintptr, value int32) {
   checkBox_SetHeight.Call(obj, uintptr(value))
}

func CheckBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := checkBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func CheckBox_SetCursor(obj uintptr, value TCursor) {
   checkBox_SetCursor.Call(obj, uintptr(value))
}

func CheckBox_GetHint(obj uintptr) string {
    ret, _, _ := checkBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetHint(obj uintptr, value string) {
   checkBox_SetHint.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetMargins.Call(obj)
    return ret
}

func CheckBox_SetMargins(obj uintptr, value uintptr) {
   checkBox_SetMargins.Call(obj, value)
}

func CheckBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := checkBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func CheckBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := checkBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func CheckBox_SetComponentIndex(obj uintptr, value int32) {
   checkBox_SetComponentIndex.Call(obj, uintptr(value))
}

func CheckBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetOwner.Call(obj)
    return ret
}

func CheckBox_GetName(obj uintptr) string {
    ret, _, _ := checkBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetName(obj uintptr, value string) {
   checkBox_SetName.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetTag(obj uintptr) int {
    ret, _, _ := checkBox_GetTag.Call(obj)
    return int(ret)
}

func CheckBox_SetTag(obj uintptr, value int) {
   checkBox_SetTag.Call(obj, uintptr(value))
}

func CheckBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := checkBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func CheckBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := checkBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

