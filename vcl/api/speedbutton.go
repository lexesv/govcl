
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func SpeedButton_Create(obj uintptr) uintptr {
    ret, _, _ := speedButton_Create.Call(obj)
    return ret
}

func SpeedButton_Free(obj uintptr) {
    speedButton_Free.Call(obj)
}

func SpeedButton_Click(obj uintptr)  {
    speedButton_Click.Call(obj)
}

func SpeedButton_BringToFront(obj uintptr)  {
    speedButton_BringToFront.Call(obj)
}

func SpeedButton_HasParent(obj uintptr) bool {
    ret, _, _ := speedButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_Hide(obj uintptr)  {
    speedButton_Hide.Call(obj)
}

func SpeedButton_Invalidate(obj uintptr)  {
    speedButton_Invalidate.Call(obj)
}

func SpeedButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := speedButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func SpeedButton_Refresh(obj uintptr)  {
    speedButton_Refresh.Call(obj)
}

func SpeedButton_Repaint(obj uintptr)  {
    speedButton_Repaint.Call(obj)
}

func SpeedButton_SendToBack(obj uintptr)  {
    speedButton_SendToBack.Call(obj)
}

func SpeedButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    speedButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func SpeedButton_Show(obj uintptr)  {
    speedButton_Show.Call(obj)
}

func SpeedButton_Update(obj uintptr)  {
    speedButton_Update.Call(obj)
}

func SpeedButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := speedButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func SpeedButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := speedButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SpeedButton_GetNamePath(obj uintptr) string {
    ret, _, _ := speedButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_Assign(obj uintptr, Source uintptr)  {
    speedButton_Assign.Call(obj, Source )
}

func SpeedButton_ClassName(obj uintptr) string {
    ret, _, _ := speedButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := speedButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SpeedButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := speedButton_GetHashCode.Call(obj)
    return int32(ret)
}

func SpeedButton_ToString(obj uintptr) string {
    ret, _, _ := speedButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetAction.Call(obj)
    return ret
}

func SpeedButton_SetAction(obj uintptr, value uintptr) {
   speedButton_SetAction.Call(obj, value)
}

func SpeedButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := speedButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func SpeedButton_SetAlign(obj uintptr, value TAlign) {
   speedButton_SetAlign.Call(obj, uintptr(value))
}

func SpeedButton_GetAllowAllUp(obj uintptr) bool {
    ret, _, _ := speedButton_GetAllowAllUp.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetAllowAllUp(obj uintptr, value bool) {
   speedButton_SetAllowAllUp.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := speedButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func SpeedButton_SetAnchors(obj uintptr, value TAnchors) {
   speedButton_SetAnchors.Call(obj, uintptr(value))
}

func SpeedButton_GetGroupIndex(obj uintptr) int32 {
    ret, _, _ := speedButton_GetGroupIndex.Call(obj)
    return int32(ret)
}

func SpeedButton_SetGroupIndex(obj uintptr, value int32) {
   speedButton_SetGroupIndex.Call(obj, uintptr(value))
}

func SpeedButton_GetDown(obj uintptr) bool {
    ret, _, _ := speedButton_GetDown.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetDown(obj uintptr, value bool) {
   speedButton_SetDown.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetCaption(obj uintptr) string {
    ret, _, _ := speedButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetCaption(obj uintptr, value string) {
   speedButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := speedButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetEnabled(obj uintptr, value bool) {
   speedButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFlat(obj uintptr) bool {
    ret, _, _ := speedButton_GetFlat.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetFlat(obj uintptr, value bool) {
   speedButton_SetFlat.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFont(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetFont.Call(obj)
    return ret
}

func SpeedButton_SetFont(obj uintptr, value uintptr) {
   speedButton_SetFont.Call(obj, value)
}

func SpeedButton_GetLayout(obj uintptr) TButtonLayout {
    ret, _, _ := speedButton_GetLayout.Call(obj)
    return TButtonLayout(ret)
}

func SpeedButton_SetLayout(obj uintptr, value TButtonLayout) {
   speedButton_SetLayout.Call(obj, uintptr(value))
}

func SpeedButton_GetParentFont(obj uintptr) bool {
    ret, _, _ := speedButton_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetParentFont(obj uintptr, value bool) {
   speedButton_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := speedButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetParentShowHint(obj uintptr, value bool) {
   speedButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetPopupMenu.Call(obj)
    return ret
}

func SpeedButton_SetPopupMenu(obj uintptr, value uintptr) {
   speedButton_SetPopupMenu.Call(obj, value)
}

func SpeedButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := speedButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetShowHint(obj uintptr, value bool) {
   speedButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetTransparent(obj uintptr) bool {
    ret, _, _ := speedButton_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetTransparent(obj uintptr, value bool) {
   speedButton_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetVisible(obj uintptr) bool {
    ret, _, _ := speedButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetVisible(obj uintptr, value bool) {
   speedButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_SetOnClick(obj uintptr, fn interface{}) {
    speedButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnDblClick(obj uintptr, fn interface{}) {
    speedButton_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func SpeedButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    speedButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func SpeedButton_SetBoundsRect(obj uintptr, value TRect) {
   speedButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func SpeedButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_SetClientHeight(obj uintptr, value int32) {
   speedButton_SetClientHeight.Call(obj, uintptr(value))
}

func SpeedButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    speedButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func SpeedButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_SetClientWidth(obj uintptr, value int32) {
   speedButton_SetClientWidth.Call(obj, uintptr(value))
}

func SpeedButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetParent.Call(obj)
    return ret
}

func SpeedButton_SetParent(obj uintptr, value uintptr) {
   speedButton_SetParent.Call(obj, value)
}

func SpeedButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := speedButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetAlignWithMargins(obj uintptr, value bool) {
   speedButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := speedButton_GetLeft.Call(obj)
    return int32(ret)
}

func SpeedButton_SetLeft(obj uintptr, value int32) {
   speedButton_SetLeft.Call(obj, uintptr(value))
}

func SpeedButton_GetTop(obj uintptr) int32 {
    ret, _, _ := speedButton_GetTop.Call(obj)
    return int32(ret)
}

func SpeedButton_SetTop(obj uintptr, value int32) {
   speedButton_SetTop.Call(obj, uintptr(value))
}

func SpeedButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_SetWidth(obj uintptr, value int32) {
   speedButton_SetWidth.Call(obj, uintptr(value))
}

func SpeedButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_SetHeight(obj uintptr, value int32) {
   speedButton_SetHeight.Call(obj, uintptr(value))
}

func SpeedButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := speedButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func SpeedButton_SetCursor(obj uintptr, value TCursor) {
   speedButton_SetCursor.Call(obj, uintptr(value))
}

func SpeedButton_GetHint(obj uintptr) string {
    ret, _, _ := speedButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetHint(obj uintptr, value string) {
   speedButton_SetHint.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetMargins.Call(obj)
    return ret
}

func SpeedButton_SetMargins(obj uintptr, value uintptr) {
   speedButton_SetMargins.Call(obj, value)
}

func SpeedButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := speedButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func SpeedButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := speedButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SpeedButton_SetComponentIndex(obj uintptr, value int32) {
   speedButton_SetComponentIndex.Call(obj, uintptr(value))
}

func SpeedButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetOwner.Call(obj)
    return ret
}

func SpeedButton_GetName(obj uintptr) string {
    ret, _, _ := speedButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetName(obj uintptr, value string) {
   speedButton_SetName.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetTag(obj uintptr) int {
    ret, _, _ := speedButton_GetTag.Call(obj)
    return int(ret)
}

func SpeedButton_SetTag(obj uintptr, value int) {
   speedButton_SetTag.Call(obj, uintptr(value))
}

func SpeedButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := speedButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

