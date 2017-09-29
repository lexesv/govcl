
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Mouse_Create() uintptr {
    ret, _, _ := mouse_Create.Call()
    return ret
}

func Mouse_Free(obj uintptr) {
    mouse_Free.Call(obj)
}

func Mouse_ClassName(obj uintptr) string {
    ret, _, _ := mouse_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Mouse_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := mouse_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Mouse_GetHashCode(obj uintptr) int32 {
    ret, _, _ := mouse_GetHashCode.Call(obj)
    return int32(ret)
}

func Mouse_ToString(obj uintptr) string {
    ret, _, _ := mouse_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Mouse_GetCapture(obj uintptr) HWND {
    ret, _, _ := mouse_GetCapture.Call(obj)
    return HWND(ret)
}

func Mouse_SetCapture(obj uintptr, value HWND) {
   mouse_SetCapture.Call(obj, uintptr(value))
}

func Mouse_GetCursorPos(obj uintptr) TPoint {
    var ret TPoint
    mouse_GetCursorPos.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Mouse_SetCursorPos(obj uintptr, value TPoint) {
   mouse_SetCursorPos.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Mouse_GetIsDragging(obj uintptr) bool {
    ret, _, _ := mouse_GetIsDragging.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetIsPanning(obj uintptr) bool {
    ret, _, _ := mouse_GetIsPanning.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetWheelPresent(obj uintptr) bool {
    ret, _, _ := mouse_GetWheelPresent.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetWheelScrollLines(obj uintptr) int32 {
    ret, _, _ := mouse_GetWheelScrollLines.Call(obj)
    return int32(ret)
}

