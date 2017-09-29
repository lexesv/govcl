
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func LinkLabel_Create(obj uintptr) uintptr {
    ret, _, _ := linkLabel_Create.Call(obj)
    return ret
}

func LinkLabel_Free(obj uintptr) {
    linkLabel_Free.Call(obj)
}

func LinkLabel_CanFocus(obj uintptr) bool {
    ret, _, _ := linkLabel_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_FlipChildren(obj uintptr, AllLevels bool)  {
    linkLabel_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func LinkLabel_Focused(obj uintptr) bool {
    ret, _, _ := linkLabel_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_HandleAllocated(obj uintptr) bool {
    ret, _, _ := linkLabel_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_Invalidate(obj uintptr)  {
    linkLabel_Invalidate.Call(obj)
}

func LinkLabel_Realign(obj uintptr)  {
    linkLabel_Realign.Call(obj)
}

func LinkLabel_Repaint(obj uintptr)  {
    linkLabel_Repaint.Call(obj)
}

func LinkLabel_ScaleBy(obj uintptr, M int32, D int32)  {
    linkLabel_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func LinkLabel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    linkLabel_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func LinkLabel_SetFocus(obj uintptr)  {
    linkLabel_SetFocus.Call(obj)
}

func LinkLabel_Update(obj uintptr)  {
    linkLabel_Update.Call(obj)
}

func LinkLabel_BringToFront(obj uintptr)  {
    linkLabel_BringToFront.Call(obj)
}

func LinkLabel_HasParent(obj uintptr) bool {
    ret, _, _ := linkLabel_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_Hide(obj uintptr)  {
    linkLabel_Hide.Call(obj)
}

func LinkLabel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := linkLabel_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func LinkLabel_Refresh(obj uintptr)  {
    linkLabel_Refresh.Call(obj)
}

func LinkLabel_SendToBack(obj uintptr)  {
    linkLabel_SendToBack.Call(obj)
}

func LinkLabel_Show(obj uintptr)  {
    linkLabel_Show.Call(obj)
}

func LinkLabel_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := linkLabel_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func LinkLabel_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := linkLabel_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func LinkLabel_GetNamePath(obj uintptr) string {
    ret, _, _ := linkLabel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_Assign(obj uintptr, Source uintptr)  {
    linkLabel_Assign.Call(obj, Source )
}

func LinkLabel_ClassName(obj uintptr) string {
    ret, _, _ := linkLabel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := linkLabel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func LinkLabel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetHashCode.Call(obj)
    return int32(ret)
}

func LinkLabel_ToString(obj uintptr) string {
    ret, _, _ := linkLabel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_GetAlign(obj uintptr) TAlign {
    ret, _, _ := linkLabel_GetAlign.Call(obj)
    return TAlign(ret)
}

func LinkLabel_SetAlign(obj uintptr, value TAlign) {
   linkLabel_SetAlign.Call(obj, uintptr(value))
}

func LinkLabel_GetAlignment(obj uintptr) TLinkAlignment {
    ret, _, _ := linkLabel_GetAlignment.Call(obj)
    return TLinkAlignment(ret)
}

func LinkLabel_SetAlignment(obj uintptr, value TLinkAlignment) {
   linkLabel_SetAlignment.Call(obj, uintptr(value))
}

func LinkLabel_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := linkLabel_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func LinkLabel_SetAnchors(obj uintptr, value TAnchors) {
   linkLabel_SetAnchors.Call(obj, uintptr(value))
}

func LinkLabel_GetAutoSize(obj uintptr) bool {
    ret, _, _ := linkLabel_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetAutoSize(obj uintptr, value bool) {
   linkLabel_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetCaption(obj uintptr) string {
    ret, _, _ := linkLabel_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetCaption(obj uintptr, value string) {
   linkLabel_SetCaption.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetColor(obj uintptr) TColor {
    ret, _, _ := linkLabel_GetColor.Call(obj)
    return TColor(ret)
}

func LinkLabel_SetColor(obj uintptr, value TColor) {
   linkLabel_SetColor.Call(obj, uintptr(value))
}

func LinkLabel_GetEnabled(obj uintptr) bool {
    ret, _, _ := linkLabel_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetEnabled(obj uintptr, value bool) {
   linkLabel_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetFont(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetFont.Call(obj)
    return ret
}

func LinkLabel_SetFont(obj uintptr, value uintptr) {
   linkLabel_SetFont.Call(obj, value)
}

func LinkLabel_GetParentColor(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentColor(obj uintptr, value bool) {
   linkLabel_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentFont(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentFont(obj uintptr, value bool) {
   linkLabel_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentShowHint(obj uintptr, value bool) {
   linkLabel_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetPopupMenu.Call(obj)
    return ret
}

func LinkLabel_SetPopupMenu(obj uintptr, value uintptr) {
   linkLabel_SetPopupMenu.Call(obj, value)
}

func LinkLabel_GetShowHint(obj uintptr) bool {
    ret, _, _ := linkLabel_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetShowHint(obj uintptr, value bool) {
   linkLabel_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := linkLabel_GetTabOrder.Call(obj)
    return int16(ret)
}

func LinkLabel_SetTabOrder(obj uintptr, value int16) {
   linkLabel_SetTabOrder.Call(obj, uintptr(value))
}

func LinkLabel_GetTabStop(obj uintptr) bool {
    ret, _, _ := linkLabel_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetTabStop(obj uintptr, value bool) {
   linkLabel_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetUseVisualStyle(obj uintptr) bool {
    ret, _, _ := linkLabel_GetUseVisualStyle.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetUseVisualStyle(obj uintptr, value bool) {
   linkLabel_SetUseVisualStyle.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetVisible(obj uintptr) bool {
    ret, _, _ := linkLabel_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetVisible(obj uintptr, value bool) {
   linkLabel_SetVisible.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_SetOnClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnDblClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseDown(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseEnter(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseLeave(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseMove(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseUp(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnLinkClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnLinkClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := linkLabel_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetDoubleBuffered(obj uintptr, value bool) {
   linkLabel_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetBrush(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetBrush.Call(obj)
    return ret
}

func LinkLabel_GetControlCount(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetControlCount.Call(obj)
    return int32(ret)
}

func LinkLabel_GetHandle(obj uintptr) HWND {
    ret, _, _ := linkLabel_GetHandle.Call(obj)
    return HWND(ret)
}

func LinkLabel_GetAction(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetAction.Call(obj)
    return ret
}

func LinkLabel_SetAction(obj uintptr, value uintptr) {
   linkLabel_SetAction.Call(obj, value)
}

func LinkLabel_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    linkLabel_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func LinkLabel_SetBoundsRect(obj uintptr, value TRect) {
   linkLabel_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func LinkLabel_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetClientHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_SetClientHeight(obj uintptr, value int32) {
   linkLabel_SetClientHeight.Call(obj, uintptr(value))
}

func LinkLabel_GetClientRect(obj uintptr) TRect {
    var ret TRect
    linkLabel_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func LinkLabel_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetClientWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_SetClientWidth(obj uintptr, value int32) {
   linkLabel_SetClientWidth.Call(obj, uintptr(value))
}

func LinkLabel_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitTop.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_GetParent(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetParent.Call(obj)
    return ret
}

func LinkLabel_SetParent(obj uintptr, value uintptr) {
   linkLabel_SetParent.Call(obj, value)
}

func LinkLabel_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := linkLabel_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetAlignWithMargins(obj uintptr, value bool) {
   linkLabel_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetLeft(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetLeft.Call(obj)
    return int32(ret)
}

func LinkLabel_SetLeft(obj uintptr, value int32) {
   linkLabel_SetLeft.Call(obj, uintptr(value))
}

func LinkLabel_GetTop(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetTop.Call(obj)
    return int32(ret)
}

func LinkLabel_SetTop(obj uintptr, value int32) {
   linkLabel_SetTop.Call(obj, uintptr(value))
}

func LinkLabel_GetWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_SetWidth(obj uintptr, value int32) {
   linkLabel_SetWidth.Call(obj, uintptr(value))
}

func LinkLabel_GetHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_SetHeight(obj uintptr, value int32) {
   linkLabel_SetHeight.Call(obj, uintptr(value))
}

func LinkLabel_GetCursor(obj uintptr) TCursor {
    ret, _, _ := linkLabel_GetCursor.Call(obj)
    return TCursor(ret)
}

func LinkLabel_SetCursor(obj uintptr, value TCursor) {
   linkLabel_SetCursor.Call(obj, uintptr(value))
}

func LinkLabel_GetHint(obj uintptr) string {
    ret, _, _ := linkLabel_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetHint(obj uintptr, value string) {
   linkLabel_SetHint.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetMargins(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetMargins.Call(obj)
    return ret
}

func LinkLabel_SetMargins(obj uintptr, value uintptr) {
   linkLabel_SetMargins.Call(obj, value)
}

func LinkLabel_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetComponentCount.Call(obj)
    return int32(ret)
}

func LinkLabel_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetComponentIndex.Call(obj)
    return int32(ret)
}

func LinkLabel_SetComponentIndex(obj uintptr, value int32) {
   linkLabel_SetComponentIndex.Call(obj, uintptr(value))
}

func LinkLabel_GetOwner(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetOwner.Call(obj)
    return ret
}

func LinkLabel_GetName(obj uintptr) string {
    ret, _, _ := linkLabel_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetName(obj uintptr, value string) {
   linkLabel_SetName.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetTag(obj uintptr) int {
    ret, _, _ := linkLabel_GetTag.Call(obj)
    return int(ret)
}

func LinkLabel_SetTag(obj uintptr, value int) {
   linkLabel_SetTag.Call(obj, uintptr(value))
}

func LinkLabel_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := linkLabel_GetControls.Call(obj, uintptr(Index))
    return ret
}

func LinkLabel_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := linkLabel_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

