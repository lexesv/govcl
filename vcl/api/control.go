
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Control_Create(obj uintptr) uintptr {
    ret, _, _ := control_Create.Call(obj)
    return ret
}

func Control_Free(obj uintptr) {
    control_Free.Call(obj)
}

func Control_BringToFront(obj uintptr)  {
    control_BringToFront.Call(obj)
}

func Control_HasParent(obj uintptr) bool {
    ret, _, _ := control_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_Hide(obj uintptr)  {
    control_Hide.Call(obj)
}

func Control_Invalidate(obj uintptr)  {
    control_Invalidate.Call(obj)
}

func Control_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := control_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Control_Refresh(obj uintptr)  {
    control_Refresh.Call(obj)
}

func Control_Repaint(obj uintptr)  {
    control_Repaint.Call(obj)
}

func Control_SendToBack(obj uintptr)  {
    control_SendToBack.Call(obj)
}

func Control_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    control_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Control_Show(obj uintptr)  {
    control_Show.Call(obj)
}

func Control_Update(obj uintptr)  {
    control_Update.Call(obj)
}

func Control_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := control_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Control_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := control_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Control_GetNamePath(obj uintptr) string {
    ret, _, _ := control_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Control_Assign(obj uintptr, Source uintptr)  {
    control_Assign.Call(obj, Source )
}

func Control_ClassName(obj uintptr) string {
    ret, _, _ := control_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Control_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := control_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Control_GetHashCode(obj uintptr) int32 {
    ret, _, _ := control_GetHashCode.Call(obj)
    return int32(ret)
}

func Control_ToString(obj uintptr) string {
    ret, _, _ := control_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Control_GetEnabled(obj uintptr) bool {
    ret, _, _ := control_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetEnabled(obj uintptr, value bool) {
   control_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Control_GetAction(obj uintptr) uintptr {
    ret, _, _ := control_GetAction.Call(obj)
    return ret
}

func Control_SetAction(obj uintptr, value uintptr) {
   control_SetAction.Call(obj, value)
}

func Control_GetAlign(obj uintptr) TAlign {
    ret, _, _ := control_GetAlign.Call(obj)
    return TAlign(ret)
}

func Control_SetAlign(obj uintptr, value TAlign) {
   control_SetAlign.Call(obj, uintptr(value))
}

func Control_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := control_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Control_SetAnchors(obj uintptr, value TAnchors) {
   control_SetAnchors.Call(obj, uintptr(value))
}

func Control_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    control_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Control_SetBoundsRect(obj uintptr, value TRect) {
   control_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Control_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := control_GetClientHeight.Call(obj)
    return int32(ret)
}

func Control_SetClientHeight(obj uintptr, value int32) {
   control_SetClientHeight.Call(obj, uintptr(value))
}

func Control_GetClientRect(obj uintptr) TRect {
    var ret TRect
    control_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Control_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := control_GetClientWidth.Call(obj)
    return int32(ret)
}

func Control_SetClientWidth(obj uintptr, value int32) {
   control_SetClientWidth.Call(obj, uintptr(value))
}

func Control_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Control_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Control_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Control_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Control_GetShowHint(obj uintptr) bool {
    ret, _, _ := control_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetShowHint(obj uintptr, value bool) {
   control_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Control_GetVisible(obj uintptr) bool {
    ret, _, _ := control_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetVisible(obj uintptr, value bool) {
   control_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Control_GetParent(obj uintptr) uintptr {
    ret, _, _ := control_GetParent.Call(obj)
    return ret
}

func Control_SetParent(obj uintptr, value uintptr) {
   control_SetParent.Call(obj, value)
}

func Control_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := control_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetAlignWithMargins(obj uintptr, value bool) {
   control_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Control_GetLeft(obj uintptr) int32 {
    ret, _, _ := control_GetLeft.Call(obj)
    return int32(ret)
}

func Control_SetLeft(obj uintptr, value int32) {
   control_SetLeft.Call(obj, uintptr(value))
}

func Control_GetTop(obj uintptr) int32 {
    ret, _, _ := control_GetTop.Call(obj)
    return int32(ret)
}

func Control_SetTop(obj uintptr, value int32) {
   control_SetTop.Call(obj, uintptr(value))
}

func Control_GetWidth(obj uintptr) int32 {
    ret, _, _ := control_GetWidth.Call(obj)
    return int32(ret)
}

func Control_SetWidth(obj uintptr, value int32) {
   control_SetWidth.Call(obj, uintptr(value))
}

func Control_GetHeight(obj uintptr) int32 {
    ret, _, _ := control_GetHeight.Call(obj)
    return int32(ret)
}

func Control_SetHeight(obj uintptr, value int32) {
   control_SetHeight.Call(obj, uintptr(value))
}

func Control_GetCursor(obj uintptr) TCursor {
    ret, _, _ := control_GetCursor.Call(obj)
    return TCursor(ret)
}

func Control_SetCursor(obj uintptr, value TCursor) {
   control_SetCursor.Call(obj, uintptr(value))
}

func Control_GetHint(obj uintptr) string {
    ret, _, _ := control_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Control_SetHint(obj uintptr, value string) {
   control_SetHint.Call(obj, GoStrToDStr(value))
}

func Control_GetMargins(obj uintptr) uintptr {
    ret, _, _ := control_GetMargins.Call(obj)
    return ret
}

func Control_SetMargins(obj uintptr, value uintptr) {
   control_SetMargins.Call(obj, value)
}

func Control_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := control_GetComponentCount.Call(obj)
    return int32(ret)
}

func Control_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := control_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Control_SetComponentIndex(obj uintptr, value int32) {
   control_SetComponentIndex.Call(obj, uintptr(value))
}

func Control_GetOwner(obj uintptr) uintptr {
    ret, _, _ := control_GetOwner.Call(obj)
    return ret
}

func Control_GetName(obj uintptr) string {
    ret, _, _ := control_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Control_SetName(obj uintptr, value string) {
   control_SetName.Call(obj, GoStrToDStr(value))
}

func Control_GetTag(obj uintptr) int {
    ret, _, _ := control_GetTag.Call(obj)
    return int(ret)
}

func Control_SetTag(obj uintptr, value int) {
   control_SetTag.Call(obj, uintptr(value))
}

func Control_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := control_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

