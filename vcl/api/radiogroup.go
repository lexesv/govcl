
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func RadioGroup_Create(obj uintptr) uintptr {
    ret, _, _ := radioGroup_Create.Call(obj)
    return ret
}

func RadioGroup_Free(obj uintptr) {
    radioGroup_Free.Call(obj)
}

func RadioGroup_FlipChildren(obj uintptr, AllLevels bool)  {
    radioGroup_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RadioGroup_CanFocus(obj uintptr) bool {
    ret, _, _ := radioGroup_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Focused(obj uintptr) bool {
    ret, _, _ := radioGroup_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_HandleAllocated(obj uintptr) bool {
    ret, _, _ := radioGroup_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Invalidate(obj uintptr)  {
    radioGroup_Invalidate.Call(obj)
}

func RadioGroup_Realign(obj uintptr)  {
    radioGroup_Realign.Call(obj)
}

func RadioGroup_Repaint(obj uintptr)  {
    radioGroup_Repaint.Call(obj)
}

func RadioGroup_ScaleBy(obj uintptr, M int32, D int32)  {
    radioGroup_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RadioGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    radioGroup_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RadioGroup_SetFocus(obj uintptr)  {
    radioGroup_SetFocus.Call(obj)
}

func RadioGroup_Update(obj uintptr)  {
    radioGroup_Update.Call(obj)
}

func RadioGroup_BringToFront(obj uintptr)  {
    radioGroup_BringToFront.Call(obj)
}

func RadioGroup_HasParent(obj uintptr) bool {
    ret, _, _ := radioGroup_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Hide(obj uintptr)  {
    radioGroup_Hide.Call(obj)
}

func RadioGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := radioGroup_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RadioGroup_Refresh(obj uintptr)  {
    radioGroup_Refresh.Call(obj)
}

func RadioGroup_SendToBack(obj uintptr)  {
    radioGroup_SendToBack.Call(obj)
}

func RadioGroup_Show(obj uintptr)  {
    radioGroup_Show.Call(obj)
}

func RadioGroup_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := radioGroup_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RadioGroup_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := radioGroup_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RadioGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := radioGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_Assign(obj uintptr, Source uintptr)  {
    radioGroup_Assign.Call(obj, Source )
}

func RadioGroup_ClassName(obj uintptr) string {
    ret, _, _ := radioGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := radioGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RadioGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func RadioGroup_ToString(obj uintptr) string {
    ret, _, _ := radioGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_GetAlign(obj uintptr) TAlign {
    ret, _, _ := radioGroup_GetAlign.Call(obj)
    return TAlign(ret)
}

func RadioGroup_SetAlign(obj uintptr, value TAlign) {
   radioGroup_SetAlign.Call(obj, uintptr(value))
}

func RadioGroup_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := radioGroup_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RadioGroup_SetAnchors(obj uintptr, value TAnchors) {
   radioGroup_SetAnchors.Call(obj, uintptr(value))
}

func RadioGroup_GetCaption(obj uintptr) string {
    ret, _, _ := radioGroup_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetCaption(obj uintptr, value string) {
   radioGroup_SetCaption.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetColor(obj uintptr) TColor {
    ret, _, _ := radioGroup_GetColor.Call(obj)
    return TColor(ret)
}

func RadioGroup_SetColor(obj uintptr, value TColor) {
   radioGroup_SetColor.Call(obj, uintptr(value))
}

func RadioGroup_GetColumns(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetColumns.Call(obj)
    return int32(ret)
}

func RadioGroup_SetColumns(obj uintptr, value int32) {
   radioGroup_SetColumns.Call(obj, uintptr(value))
}

func RadioGroup_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioGroup_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetDoubleBuffered(obj uintptr, value bool) {
   radioGroup_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetEnabled(obj uintptr) bool {
    ret, _, _ := radioGroup_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetEnabled(obj uintptr, value bool) {
   radioGroup_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetFont(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetFont.Call(obj)
    return ret
}

func RadioGroup_SetFont(obj uintptr, value uintptr) {
   radioGroup_SetFont.Call(obj, value)
}

func RadioGroup_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetItemIndex.Call(obj)
    return int32(ret)
}

func RadioGroup_SetItemIndex(obj uintptr, value int32) {
   radioGroup_SetItemIndex.Call(obj, uintptr(value))
}

func RadioGroup_GetItems(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetItems.Call(obj)
    return ret
}

func RadioGroup_SetItems(obj uintptr, value uintptr) {
   radioGroup_SetItems.Call(obj, value)
}

func RadioGroup_GetParentColor(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentColor(obj uintptr, value bool) {
   radioGroup_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentCtl3D(obj uintptr, value bool) {
   radioGroup_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentFont(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentFont(obj uintptr, value bool) {
   radioGroup_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentShowHint(obj uintptr, value bool) {
   radioGroup_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetPopupMenu.Call(obj)
    return ret
}

func RadioGroup_SetPopupMenu(obj uintptr, value uintptr) {
   radioGroup_SetPopupMenu.Call(obj, value)
}

func RadioGroup_GetShowHint(obj uintptr) bool {
    ret, _, _ := radioGroup_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetShowHint(obj uintptr, value bool) {
   radioGroup_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := radioGroup_GetTabOrder.Call(obj)
    return int16(ret)
}

func RadioGroup_SetTabOrder(obj uintptr, value int16) {
   radioGroup_SetTabOrder.Call(obj, uintptr(value))
}

func RadioGroup_GetTabStop(obj uintptr) bool {
    ret, _, _ := radioGroup_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetTabStop(obj uintptr, value bool) {
   radioGroup_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetVisible(obj uintptr) bool {
    ret, _, _ := radioGroup_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetVisible(obj uintptr, value bool) {
   radioGroup_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetWordWrap(obj uintptr) bool {
    ret, _, _ := radioGroup_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetWordWrap(obj uintptr, value bool) {
   radioGroup_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_SetOnClick(obj uintptr, fn interface{}) {
    radioGroup_SetOnClick.Call(obj, addEventToMap(fn))
}

func RadioGroup_SetOnEnter(obj uintptr, fn interface{}) {
    radioGroup_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RadioGroup_SetOnExit(obj uintptr, fn interface{}) {
    radioGroup_SetOnExit.Call(obj, addEventToMap(fn))
}

func RadioGroup_GetBrush(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetBrush.Call(obj)
    return ret
}

func RadioGroup_GetControlCount(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetControlCount.Call(obj)
    return int32(ret)
}

func RadioGroup_GetHandle(obj uintptr) HWND {
    ret, _, _ := radioGroup_GetHandle.Call(obj)
    return HWND(ret)
}

func RadioGroup_GetAction(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetAction.Call(obj)
    return ret
}

func RadioGroup_SetAction(obj uintptr, value uintptr) {
   radioGroup_SetAction.Call(obj, value)
}

func RadioGroup_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    radioGroup_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioGroup_SetBoundsRect(obj uintptr, value TRect) {
   radioGroup_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioGroup_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetClientHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_SetClientHeight(obj uintptr, value int32) {
   radioGroup_SetClientHeight.Call(obj, uintptr(value))
}

func RadioGroup_GetClientRect(obj uintptr) TRect {
    var ret TRect
    radioGroup_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioGroup_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetClientWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_SetClientWidth(obj uintptr, value int32) {
   radioGroup_SetClientWidth.Call(obj, uintptr(value))
}

func RadioGroup_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_GetParent(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetParent.Call(obj)
    return ret
}

func RadioGroup_SetParent(obj uintptr, value uintptr) {
   radioGroup_SetParent.Call(obj, value)
}

func RadioGroup_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := radioGroup_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetAlignWithMargins(obj uintptr, value bool) {
   radioGroup_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetLeft(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetLeft.Call(obj)
    return int32(ret)
}

func RadioGroup_SetLeft(obj uintptr, value int32) {
   radioGroup_SetLeft.Call(obj, uintptr(value))
}

func RadioGroup_GetTop(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetTop.Call(obj)
    return int32(ret)
}

func RadioGroup_SetTop(obj uintptr, value int32) {
   radioGroup_SetTop.Call(obj, uintptr(value))
}

func RadioGroup_GetWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_SetWidth(obj uintptr, value int32) {
   radioGroup_SetWidth.Call(obj, uintptr(value))
}

func RadioGroup_GetHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_SetHeight(obj uintptr, value int32) {
   radioGroup_SetHeight.Call(obj, uintptr(value))
}

func RadioGroup_GetCursor(obj uintptr) TCursor {
    ret, _, _ := radioGroup_GetCursor.Call(obj)
    return TCursor(ret)
}

func RadioGroup_SetCursor(obj uintptr, value TCursor) {
   radioGroup_SetCursor.Call(obj, uintptr(value))
}

func RadioGroup_GetHint(obj uintptr) string {
    ret, _, _ := radioGroup_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetHint(obj uintptr, value string) {
   radioGroup_SetHint.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetMargins(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetMargins.Call(obj)
    return ret
}

func RadioGroup_SetMargins(obj uintptr, value uintptr) {
   radioGroup_SetMargins.Call(obj, value)
}

func RadioGroup_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetComponentCount.Call(obj)
    return int32(ret)
}

func RadioGroup_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RadioGroup_SetComponentIndex(obj uintptr, value int32) {
   radioGroup_SetComponentIndex.Call(obj, uintptr(value))
}

func RadioGroup_GetOwner(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetOwner.Call(obj)
    return ret
}

func RadioGroup_GetName(obj uintptr) string {
    ret, _, _ := radioGroup_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetName(obj uintptr, value string) {
   radioGroup_SetName.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetTag(obj uintptr) int {
    ret, _, _ := radioGroup_GetTag.Call(obj)
    return int(ret)
}

func RadioGroup_SetTag(obj uintptr, value int) {
   radioGroup_SetTag.Call(obj, uintptr(value))
}

func RadioGroup_GetButtons(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioGroup_GetButtons.Call(obj, uintptr(Index))
    return ret
}

func RadioGroup_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioGroup_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RadioGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := radioGroup_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

