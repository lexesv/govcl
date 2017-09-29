
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Splitter_Create(obj uintptr) uintptr {
    ret, _, _ := splitter_Create.Call(obj)
    return ret
}

func Splitter_Free(obj uintptr) {
    splitter_Free.Call(obj)
}

func Splitter_BringToFront(obj uintptr)  {
    splitter_BringToFront.Call(obj)
}

func Splitter_HasParent(obj uintptr) bool {
    ret, _, _ := splitter_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_Hide(obj uintptr)  {
    splitter_Hide.Call(obj)
}

func Splitter_Invalidate(obj uintptr)  {
    splitter_Invalidate.Call(obj)
}

func Splitter_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := splitter_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Splitter_Refresh(obj uintptr)  {
    splitter_Refresh.Call(obj)
}

func Splitter_Repaint(obj uintptr)  {
    splitter_Repaint.Call(obj)
}

func Splitter_SendToBack(obj uintptr)  {
    splitter_SendToBack.Call(obj)
}

func Splitter_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    splitter_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Splitter_Show(obj uintptr)  {
    splitter_Show.Call(obj)
}

func Splitter_Update(obj uintptr)  {
    splitter_Update.Call(obj)
}

func Splitter_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := splitter_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Splitter_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := splitter_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Splitter_GetNamePath(obj uintptr) string {
    ret, _, _ := splitter_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_Assign(obj uintptr, Source uintptr)  {
    splitter_Assign.Call(obj, Source )
}

func Splitter_ClassName(obj uintptr) string {
    ret, _, _ := splitter_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := splitter_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Splitter_GetHashCode(obj uintptr) int32 {
    ret, _, _ := splitter_GetHashCode.Call(obj)
    return int32(ret)
}

func Splitter_ToString(obj uintptr) string {
    ret, _, _ := splitter_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := splitter_GetCanvas.Call(obj)
    return ret
}

func Splitter_GetAlign(obj uintptr) TAlign {
    ret, _, _ := splitter_GetAlign.Call(obj)
    return TAlign(ret)
}

func Splitter_SetAlign(obj uintptr, value TAlign) {
   splitter_SetAlign.Call(obj, uintptr(value))
}

func Splitter_GetColor(obj uintptr) TColor {
    ret, _, _ := splitter_GetColor.Call(obj)
    return TColor(ret)
}

func Splitter_SetColor(obj uintptr, value TColor) {
   splitter_SetColor.Call(obj, uintptr(value))
}

func Splitter_GetCursor(obj uintptr) TCursor {
    ret, _, _ := splitter_GetCursor.Call(obj)
    return TCursor(ret)
}

func Splitter_SetCursor(obj uintptr, value TCursor) {
   splitter_SetCursor.Call(obj, uintptr(value))
}

func Splitter_GetParentColor(obj uintptr) bool {
    ret, _, _ := splitter_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetParentColor(obj uintptr, value bool) {
   splitter_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetVisible(obj uintptr) bool {
    ret, _, _ := splitter_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetVisible(obj uintptr, value bool) {
   splitter_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetWidth.Call(obj)
    return int32(ret)
}

func Splitter_SetWidth(obj uintptr, value int32) {
   splitter_SetWidth.Call(obj, uintptr(value))
}

func Splitter_SetOnPaint(obj uintptr, fn interface{}) {
    splitter_SetOnPaint.Call(obj, addEventToMap(fn))
}

func Splitter_GetEnabled(obj uintptr) bool {
    ret, _, _ := splitter_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetEnabled(obj uintptr, value bool) {
   splitter_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetAction(obj uintptr) uintptr {
    ret, _, _ := splitter_GetAction.Call(obj)
    return ret
}

func Splitter_SetAction(obj uintptr, value uintptr) {
   splitter_SetAction.Call(obj, value)
}

func Splitter_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := splitter_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Splitter_SetAnchors(obj uintptr, value TAnchors) {
   splitter_SetAnchors.Call(obj, uintptr(value))
}

func Splitter_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    splitter_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Splitter_SetBoundsRect(obj uintptr, value TRect) {
   splitter_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Splitter_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetClientHeight.Call(obj)
    return int32(ret)
}

func Splitter_SetClientHeight(obj uintptr, value int32) {
   splitter_SetClientHeight.Call(obj, uintptr(value))
}

func Splitter_GetClientRect(obj uintptr) TRect {
    var ret TRect
    splitter_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Splitter_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetClientWidth.Call(obj)
    return int32(ret)
}

func Splitter_SetClientWidth(obj uintptr, value int32) {
   splitter_SetClientWidth.Call(obj, uintptr(value))
}

func Splitter_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Splitter_GetShowHint(obj uintptr) bool {
    ret, _, _ := splitter_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetShowHint(obj uintptr, value bool) {
   splitter_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetParent(obj uintptr) uintptr {
    ret, _, _ := splitter_GetParent.Call(obj)
    return ret
}

func Splitter_SetParent(obj uintptr, value uintptr) {
   splitter_SetParent.Call(obj, value)
}

func Splitter_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := splitter_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetAlignWithMargins(obj uintptr, value bool) {
   splitter_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetLeft(obj uintptr) int32 {
    ret, _, _ := splitter_GetLeft.Call(obj)
    return int32(ret)
}

func Splitter_SetLeft(obj uintptr, value int32) {
   splitter_SetLeft.Call(obj, uintptr(value))
}

func Splitter_GetTop(obj uintptr) int32 {
    ret, _, _ := splitter_GetTop.Call(obj)
    return int32(ret)
}

func Splitter_SetTop(obj uintptr, value int32) {
   splitter_SetTop.Call(obj, uintptr(value))
}

func Splitter_GetHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetHeight.Call(obj)
    return int32(ret)
}

func Splitter_SetHeight(obj uintptr, value int32) {
   splitter_SetHeight.Call(obj, uintptr(value))
}

func Splitter_GetHint(obj uintptr) string {
    ret, _, _ := splitter_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_SetHint(obj uintptr, value string) {
   splitter_SetHint.Call(obj, GoStrToDStr(value))
}

func Splitter_GetMargins(obj uintptr) uintptr {
    ret, _, _ := splitter_GetMargins.Call(obj)
    return ret
}

func Splitter_SetMargins(obj uintptr, value uintptr) {
   splitter_SetMargins.Call(obj, value)
}

func Splitter_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := splitter_GetComponentCount.Call(obj)
    return int32(ret)
}

func Splitter_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := splitter_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Splitter_SetComponentIndex(obj uintptr, value int32) {
   splitter_SetComponentIndex.Call(obj, uintptr(value))
}

func Splitter_GetOwner(obj uintptr) uintptr {
    ret, _, _ := splitter_GetOwner.Call(obj)
    return ret
}

func Splitter_GetName(obj uintptr) string {
    ret, _, _ := splitter_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_SetName(obj uintptr, value string) {
   splitter_SetName.Call(obj, GoStrToDStr(value))
}

func Splitter_GetTag(obj uintptr) int {
    ret, _, _ := splitter_GetTag.Call(obj)
    return int(ret)
}

func Splitter_SetTag(obj uintptr, value int) {
   splitter_SetTag.Call(obj, uintptr(value))
}

func Splitter_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := splitter_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

