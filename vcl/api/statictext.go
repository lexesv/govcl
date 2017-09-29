
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func StaticText_Create(obj uintptr) uintptr {
    ret, _, _ := staticText_Create.Call(obj)
    return ret
}

func StaticText_Free(obj uintptr) {
    staticText_Free.Call(obj)
}

func StaticText_CanFocus(obj uintptr) bool {
    ret, _, _ := staticText_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_FlipChildren(obj uintptr, AllLevels bool)  {
    staticText_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func StaticText_Focused(obj uintptr) bool {
    ret, _, _ := staticText_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_HandleAllocated(obj uintptr) bool {
    ret, _, _ := staticText_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_Invalidate(obj uintptr)  {
    staticText_Invalidate.Call(obj)
}

func StaticText_Realign(obj uintptr)  {
    staticText_Realign.Call(obj)
}

func StaticText_Repaint(obj uintptr)  {
    staticText_Repaint.Call(obj)
}

func StaticText_ScaleBy(obj uintptr, M int32, D int32)  {
    staticText_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func StaticText_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    staticText_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func StaticText_SetFocus(obj uintptr)  {
    staticText_SetFocus.Call(obj)
}

func StaticText_Update(obj uintptr)  {
    staticText_Update.Call(obj)
}

func StaticText_BringToFront(obj uintptr)  {
    staticText_BringToFront.Call(obj)
}

func StaticText_HasParent(obj uintptr) bool {
    ret, _, _ := staticText_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_Hide(obj uintptr)  {
    staticText_Hide.Call(obj)
}

func StaticText_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := staticText_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func StaticText_Refresh(obj uintptr)  {
    staticText_Refresh.Call(obj)
}

func StaticText_SendToBack(obj uintptr)  {
    staticText_SendToBack.Call(obj)
}

func StaticText_Show(obj uintptr)  {
    staticText_Show.Call(obj)
}

func StaticText_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := staticText_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func StaticText_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := staticText_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func StaticText_GetNamePath(obj uintptr) string {
    ret, _, _ := staticText_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_Assign(obj uintptr, Source uintptr)  {
    staticText_Assign.Call(obj, Source )
}

func StaticText_ClassName(obj uintptr) string {
    ret, _, _ := staticText_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := staticText_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StaticText_GetHashCode(obj uintptr) int32 {
    ret, _, _ := staticText_GetHashCode.Call(obj)
    return int32(ret)
}

func StaticText_ToString(obj uintptr) string {
    ret, _, _ := staticText_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_GetAlign(obj uintptr) TAlign {
    ret, _, _ := staticText_GetAlign.Call(obj)
    return TAlign(ret)
}

func StaticText_SetAlign(obj uintptr, value TAlign) {
   staticText_SetAlign.Call(obj, uintptr(value))
}

func StaticText_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := staticText_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func StaticText_SetAlignment(obj uintptr, value TAlignment) {
   staticText_SetAlignment.Call(obj, uintptr(value))
}

func StaticText_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := staticText_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func StaticText_SetAnchors(obj uintptr, value TAnchors) {
   staticText_SetAnchors.Call(obj, uintptr(value))
}

func StaticText_GetAutoSize(obj uintptr) bool {
    ret, _, _ := staticText_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetAutoSize(obj uintptr, value bool) {
   staticText_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetBorderStyle(obj uintptr) TStaticBorderStyle {
    ret, _, _ := staticText_GetBorderStyle.Call(obj)
    return TStaticBorderStyle(ret)
}

func StaticText_SetBorderStyle(obj uintptr, value TStaticBorderStyle) {
   staticText_SetBorderStyle.Call(obj, uintptr(value))
}

func StaticText_GetCaption(obj uintptr) string {
    ret, _, _ := staticText_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetCaption(obj uintptr, value string) {
   staticText_SetCaption.Call(obj, GoStrToDStr(value))
}

func StaticText_GetColor(obj uintptr) TColor {
    ret, _, _ := staticText_GetColor.Call(obj)
    return TColor(ret)
}

func StaticText_SetColor(obj uintptr, value TColor) {
   staticText_SetColor.Call(obj, uintptr(value))
}

func StaticText_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := staticText_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetDoubleBuffered(obj uintptr, value bool) {
   staticText_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetEnabled(obj uintptr) bool {
    ret, _, _ := staticText_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetEnabled(obj uintptr, value bool) {
   staticText_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetFont(obj uintptr) uintptr {
    ret, _, _ := staticText_GetFont.Call(obj)
    return ret
}

func StaticText_SetFont(obj uintptr, value uintptr) {
   staticText_SetFont.Call(obj, value)
}

func StaticText_GetParentColor(obj uintptr) bool {
    ret, _, _ := staticText_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentColor(obj uintptr, value bool) {
   staticText_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentFont(obj uintptr) bool {
    ret, _, _ := staticText_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentFont(obj uintptr, value bool) {
   staticText_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := staticText_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentShowHint(obj uintptr, value bool) {
   staticText_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := staticText_GetPopupMenu.Call(obj)
    return ret
}

func StaticText_SetPopupMenu(obj uintptr, value uintptr) {
   staticText_SetPopupMenu.Call(obj, value)
}

func StaticText_GetShowAccelChar(obj uintptr) bool {
    ret, _, _ := staticText_GetShowAccelChar.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetShowAccelChar(obj uintptr, value bool) {
   staticText_SetShowAccelChar.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetShowHint(obj uintptr) bool {
    ret, _, _ := staticText_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetShowHint(obj uintptr, value bool) {
   staticText_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := staticText_GetTabOrder.Call(obj)
    return int16(ret)
}

func StaticText_SetTabOrder(obj uintptr, value int16) {
   staticText_SetTabOrder.Call(obj, uintptr(value))
}

func StaticText_GetTabStop(obj uintptr) bool {
    ret, _, _ := staticText_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetTabStop(obj uintptr, value bool) {
   staticText_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTransparent(obj uintptr) bool {
    ret, _, _ := staticText_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetTransparent(obj uintptr, value bool) {
   staticText_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetVisible(obj uintptr) bool {
    ret, _, _ := staticText_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetVisible(obj uintptr, value bool) {
   staticText_SetVisible.Call(obj, GoBoolToDBool(value))
}

func StaticText_SetOnClick(obj uintptr, fn interface{}) {
    staticText_SetOnClick.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnDblClick(obj uintptr, fn interface{}) {
    staticText_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseDown(obj uintptr, fn interface{}) {
    staticText_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseEnter(obj uintptr, fn interface{}) {
    staticText_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseLeave(obj uintptr, fn interface{}) {
    staticText_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseMove(obj uintptr, fn interface{}) {
    staticText_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseUp(obj uintptr, fn interface{}) {
    staticText_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func StaticText_GetBrush(obj uintptr) uintptr {
    ret, _, _ := staticText_GetBrush.Call(obj)
    return ret
}

func StaticText_GetControlCount(obj uintptr) int32 {
    ret, _, _ := staticText_GetControlCount.Call(obj)
    return int32(ret)
}

func StaticText_GetHandle(obj uintptr) HWND {
    ret, _, _ := staticText_GetHandle.Call(obj)
    return HWND(ret)
}

func StaticText_GetAction(obj uintptr) uintptr {
    ret, _, _ := staticText_GetAction.Call(obj)
    return ret
}

func StaticText_SetAction(obj uintptr, value uintptr) {
   staticText_SetAction.Call(obj, value)
}

func StaticText_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    staticText_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StaticText_SetBoundsRect(obj uintptr, value TRect) {
   staticText_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StaticText_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetClientHeight.Call(obj)
    return int32(ret)
}

func StaticText_SetClientHeight(obj uintptr, value int32) {
   staticText_SetClientHeight.Call(obj, uintptr(value))
}

func StaticText_GetClientRect(obj uintptr) TRect {
    var ret TRect
    staticText_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StaticText_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetClientWidth.Call(obj)
    return int32(ret)
}

func StaticText_SetClientWidth(obj uintptr, value int32) {
   staticText_SetClientWidth.Call(obj, uintptr(value))
}

func StaticText_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitTop.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func StaticText_GetParent(obj uintptr) uintptr {
    ret, _, _ := staticText_GetParent.Call(obj)
    return ret
}

func StaticText_SetParent(obj uintptr, value uintptr) {
   staticText_SetParent.Call(obj, value)
}

func StaticText_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := staticText_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetAlignWithMargins(obj uintptr, value bool) {
   staticText_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetLeft(obj uintptr) int32 {
    ret, _, _ := staticText_GetLeft.Call(obj)
    return int32(ret)
}

func StaticText_SetLeft(obj uintptr, value int32) {
   staticText_SetLeft.Call(obj, uintptr(value))
}

func StaticText_GetTop(obj uintptr) int32 {
    ret, _, _ := staticText_GetTop.Call(obj)
    return int32(ret)
}

func StaticText_SetTop(obj uintptr, value int32) {
   staticText_SetTop.Call(obj, uintptr(value))
}

func StaticText_GetWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetWidth.Call(obj)
    return int32(ret)
}

func StaticText_SetWidth(obj uintptr, value int32) {
   staticText_SetWidth.Call(obj, uintptr(value))
}

func StaticText_GetHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetHeight.Call(obj)
    return int32(ret)
}

func StaticText_SetHeight(obj uintptr, value int32) {
   staticText_SetHeight.Call(obj, uintptr(value))
}

func StaticText_GetCursor(obj uintptr) TCursor {
    ret, _, _ := staticText_GetCursor.Call(obj)
    return TCursor(ret)
}

func StaticText_SetCursor(obj uintptr, value TCursor) {
   staticText_SetCursor.Call(obj, uintptr(value))
}

func StaticText_GetHint(obj uintptr) string {
    ret, _, _ := staticText_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetHint(obj uintptr, value string) {
   staticText_SetHint.Call(obj, GoStrToDStr(value))
}

func StaticText_GetMargins(obj uintptr) uintptr {
    ret, _, _ := staticText_GetMargins.Call(obj)
    return ret
}

func StaticText_SetMargins(obj uintptr, value uintptr) {
   staticText_SetMargins.Call(obj, value)
}

func StaticText_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := staticText_GetComponentCount.Call(obj)
    return int32(ret)
}

func StaticText_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := staticText_GetComponentIndex.Call(obj)
    return int32(ret)
}

func StaticText_SetComponentIndex(obj uintptr, value int32) {
   staticText_SetComponentIndex.Call(obj, uintptr(value))
}

func StaticText_GetOwner(obj uintptr) uintptr {
    ret, _, _ := staticText_GetOwner.Call(obj)
    return ret
}

func StaticText_GetName(obj uintptr) string {
    ret, _, _ := staticText_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetName(obj uintptr, value string) {
   staticText_SetName.Call(obj, GoStrToDStr(value))
}

func StaticText_GetTag(obj uintptr) int {
    ret, _, _ := staticText_GetTag.Call(obj)
    return int(ret)
}

func StaticText_SetTag(obj uintptr, value int) {
   staticText_SetTag.Call(obj, uintptr(value))
}

func StaticText_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := staticText_GetControls.Call(obj, uintptr(Index))
    return ret
}

func StaticText_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := staticText_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

