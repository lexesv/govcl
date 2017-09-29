
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Screen_Create(obj uintptr) uintptr {
    ret, _, _ := screen_Create.Call(obj)
    return ret
}

func Screen_Free(obj uintptr) {
    screen_Free.Call(obj)
}

func Screen_Realign(obj uintptr)  {
    screen_Realign.Call(obj)
}

func Screen_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := screen_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Screen_GetNamePath(obj uintptr) string {
    ret, _, _ := screen_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_HasParent(obj uintptr) bool {
    ret, _, _ := screen_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Screen_Assign(obj uintptr, Source uintptr)  {
    screen_Assign.Call(obj, Source )
}

func Screen_ClassName(obj uintptr) string {
    ret, _, _ := screen_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := screen_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Screen_GetHashCode(obj uintptr) int32 {
    ret, _, _ := screen_GetHashCode.Call(obj)
    return int32(ret)
}

func Screen_ToString(obj uintptr) string {
    ret, _, _ := screen_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_GetActiveForm(obj uintptr) uintptr {
    ret, _, _ := screen_GetActiveForm.Call(obj)
    return ret
}

func Screen_GetCustomFormCount(obj uintptr) int32 {
    ret, _, _ := screen_GetCustomFormCount.Call(obj)
    return int32(ret)
}

func Screen_GetCursorCount(obj uintptr) int32 {
    ret, _, _ := screen_GetCursorCount.Call(obj)
    return int32(ret)
}

func Screen_GetCursor(obj uintptr) TCursor {
    ret, _, _ := screen_GetCursor.Call(obj)
    return TCursor(ret)
}

func Screen_SetCursor(obj uintptr, value TCursor) {
   screen_SetCursor.Call(obj, uintptr(value))
}

func Screen_GetFocusedForm(obj uintptr) uintptr {
    ret, _, _ := screen_GetFocusedForm.Call(obj)
    return ret
}

func Screen_SetFocusedForm(obj uintptr, value uintptr) {
   screen_SetFocusedForm.Call(obj, value)
}

func Screen_GetMonitorCount(obj uintptr) int32 {
    ret, _, _ := screen_GetMonitorCount.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopRect(obj uintptr) TRect {
    var ret TRect
    screen_GetDesktopRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Screen_GetDesktopHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopHeight.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopLeft(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopLeft.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopTop(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopTop.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopWidth.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaRect(obj uintptr) TRect {
    var ret TRect
    screen_GetWorkAreaRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Screen_GetWorkAreaHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaHeight.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaLeft(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaLeft.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaTop(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaTop.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaWidth.Call(obj)
    return int32(ret)
}

func Screen_GetFonts(obj uintptr) uintptr {
    ret, _, _ := screen_GetFonts.Call(obj)
    return ret
}

func Screen_GetImes(obj uintptr) uintptr {
    ret, _, _ := screen_GetImes.Call(obj)
    return ret
}

func Screen_GetDefaultIme(obj uintptr) string {
    ret, _, _ := screen_GetDefaultIme.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_GetHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetHeight.Call(obj)
    return int32(ret)
}

func Screen_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := screen_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Screen_GetPrimaryMonitor(obj uintptr) uintptr {
    ret, _, _ := screen_GetPrimaryMonitor.Call(obj)
    return ret
}

func Screen_GetWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetWidth.Call(obj)
    return int32(ret)
}

func Screen_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := screen_GetComponentCount.Call(obj)
    return int32(ret)
}

func Screen_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := screen_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Screen_SetComponentIndex(obj uintptr, value int32) {
   screen_SetComponentIndex.Call(obj, uintptr(value))
}

func Screen_GetOwner(obj uintptr) uintptr {
    ret, _, _ := screen_GetOwner.Call(obj)
    return ret
}

func Screen_GetName(obj uintptr) string {
    ret, _, _ := screen_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_SetName(obj uintptr, value string) {
   screen_SetName.Call(obj, GoStrToDStr(value))
}

func Screen_GetTag(obj uintptr) int {
    ret, _, _ := screen_GetTag.Call(obj)
    return int(ret)
}

func Screen_SetTag(obj uintptr, value int) {
   screen_SetTag.Call(obj, uintptr(value))
}

func Screen_GetCursors(obj uintptr, Index int32) HICON {
    ret, _, _ := screen_GetCursors.Call(obj, uintptr(Index))
    return HICON(ret)
}

func Screen_SetCursors(obj uintptr, Index int32, value HICON) {
   screen_SetCursors.Call(obj, uintptr(Index), uintptr(value))
}

func Screen_GetMonitors(obj uintptr, Index int32) uintptr {
    ret, _, _ := screen_GetMonitors.Call(obj, uintptr(Index))
    return ret
}

func Screen_GetForms(obj uintptr, Index int32) uintptr {
    ret, _, _ := screen_GetForms.Call(obj, uintptr(Index))
    return ret
}

func Screen_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := screen_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

