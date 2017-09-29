
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func PaintBox_Create(obj uintptr) uintptr {
    ret, _, _ := paintBox_Create.Call(obj)
    return ret
}

func PaintBox_Free(obj uintptr) {
    paintBox_Free.Call(obj)
}

func PaintBox_BringToFront(obj uintptr)  {
    paintBox_BringToFront.Call(obj)
}

func PaintBox_HasParent(obj uintptr) bool {
    ret, _, _ := paintBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_Hide(obj uintptr)  {
    paintBox_Hide.Call(obj)
}

func PaintBox_Invalidate(obj uintptr)  {
    paintBox_Invalidate.Call(obj)
}

func PaintBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := paintBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func PaintBox_Refresh(obj uintptr)  {
    paintBox_Refresh.Call(obj)
}

func PaintBox_Repaint(obj uintptr)  {
    paintBox_Repaint.Call(obj)
}

func PaintBox_SendToBack(obj uintptr)  {
    paintBox_SendToBack.Call(obj)
}

func PaintBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    paintBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func PaintBox_Show(obj uintptr)  {
    paintBox_Show.Call(obj)
}

func PaintBox_Update(obj uintptr)  {
    paintBox_Update.Call(obj)
}

func PaintBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := paintBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func PaintBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := paintBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PaintBox_GetNamePath(obj uintptr) string {
    ret, _, _ := paintBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_Assign(obj uintptr, Source uintptr)  {
    paintBox_Assign.Call(obj, Source )
}

func PaintBox_ClassName(obj uintptr) string {
    ret, _, _ := paintBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := paintBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PaintBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := paintBox_GetHashCode.Call(obj)
    return int32(ret)
}

func PaintBox_ToString(obj uintptr) string {
    ret, _, _ := paintBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetCanvas.Call(obj)
    return ret
}

func PaintBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := paintBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func PaintBox_SetAlign(obj uintptr, value TAlign) {
   paintBox_SetAlign.Call(obj, uintptr(value))
}

func PaintBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := paintBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func PaintBox_SetAnchors(obj uintptr, value TAnchors) {
   paintBox_SetAnchors.Call(obj, uintptr(value))
}

func PaintBox_GetColor(obj uintptr) TColor {
    ret, _, _ := paintBox_GetColor.Call(obj)
    return TColor(ret)
}

func PaintBox_SetColor(obj uintptr, value TColor) {
   paintBox_SetColor.Call(obj, uintptr(value))
}

func PaintBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := paintBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetEnabled(obj uintptr, value bool) {
   paintBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetFont.Call(obj)
    return ret
}

func PaintBox_SetFont(obj uintptr, value uintptr) {
   paintBox_SetFont.Call(obj, value)
}

func PaintBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentColor(obj uintptr, value bool) {
   paintBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentFont(obj uintptr, value bool) {
   paintBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentShowHint(obj uintptr, value bool) {
   paintBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetPopupMenu.Call(obj)
    return ret
}

func PaintBox_SetPopupMenu(obj uintptr, value uintptr) {
   paintBox_SetPopupMenu.Call(obj, value)
}

func PaintBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := paintBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetShowHint(obj uintptr, value bool) {
   paintBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetVisible(obj uintptr) bool {
    ret, _, _ := paintBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetVisible(obj uintptr, value bool) {
   paintBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func PaintBox_SetOnClick(obj uintptr, fn interface{}) {
    paintBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnDblClick(obj uintptr, fn interface{}) {
    paintBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnPaint(obj uintptr, fn interface{}) {
    paintBox_SetOnPaint.Call(obj, addEventToMap(fn))
}

func PaintBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetAction.Call(obj)
    return ret
}

func PaintBox_SetAction(obj uintptr, value uintptr) {
   paintBox_SetAction.Call(obj, value)
}

func PaintBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    paintBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PaintBox_SetBoundsRect(obj uintptr, value TRect) {
   paintBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PaintBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func PaintBox_SetClientHeight(obj uintptr, value int32) {
   paintBox_SetClientHeight.Call(obj, uintptr(value))
}

func PaintBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    paintBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PaintBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func PaintBox_SetClientWidth(obj uintptr, value int32) {
   paintBox_SetClientWidth.Call(obj, uintptr(value))
}

func PaintBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func PaintBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetParent.Call(obj)
    return ret
}

func PaintBox_SetParent(obj uintptr, value uintptr) {
   paintBox_SetParent.Call(obj, value)
}

func PaintBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := paintBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetAlignWithMargins(obj uintptr, value bool) {
   paintBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := paintBox_GetLeft.Call(obj)
    return int32(ret)
}

func PaintBox_SetLeft(obj uintptr, value int32) {
   paintBox_SetLeft.Call(obj, uintptr(value))
}

func PaintBox_GetTop(obj uintptr) int32 {
    ret, _, _ := paintBox_GetTop.Call(obj)
    return int32(ret)
}

func PaintBox_SetTop(obj uintptr, value int32) {
   paintBox_SetTop.Call(obj, uintptr(value))
}

func PaintBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetWidth.Call(obj)
    return int32(ret)
}

func PaintBox_SetWidth(obj uintptr, value int32) {
   paintBox_SetWidth.Call(obj, uintptr(value))
}

func PaintBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetHeight.Call(obj)
    return int32(ret)
}

func PaintBox_SetHeight(obj uintptr, value int32) {
   paintBox_SetHeight.Call(obj, uintptr(value))
}

func PaintBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := paintBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func PaintBox_SetCursor(obj uintptr, value TCursor) {
   paintBox_SetCursor.Call(obj, uintptr(value))
}

func PaintBox_GetHint(obj uintptr) string {
    ret, _, _ := paintBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_SetHint(obj uintptr, value string) {
   paintBox_SetHint.Call(obj, GoStrToDStr(value))
}

func PaintBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetMargins.Call(obj)
    return ret
}

func PaintBox_SetMargins(obj uintptr, value uintptr) {
   paintBox_SetMargins.Call(obj, value)
}

func PaintBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := paintBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func PaintBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := paintBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PaintBox_SetComponentIndex(obj uintptr, value int32) {
   paintBox_SetComponentIndex.Call(obj, uintptr(value))
}

func PaintBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetOwner.Call(obj)
    return ret
}

func PaintBox_GetName(obj uintptr) string {
    ret, _, _ := paintBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_SetName(obj uintptr, value string) {
   paintBox_SetName.Call(obj, GoStrToDStr(value))
}

func PaintBox_GetTag(obj uintptr) int {
    ret, _, _ := paintBox_GetTag.Call(obj)
    return int(ret)
}

func PaintBox_SetTag(obj uintptr, value int) {
   paintBox_SetTag.Call(obj, uintptr(value))
}

func PaintBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := paintBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

