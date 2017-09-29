
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Monitor_Create() uintptr {
    ret, _, _ := monitor_Create.Call()
    return ret
}

func Monitor_Free(obj uintptr) {
    monitor_Free.Call(obj)
}

func Monitor_ClassName(obj uintptr) string {
    ret, _, _ := monitor_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Monitor_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := monitor_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Monitor_GetHashCode(obj uintptr) int32 {
    ret, _, _ := monitor_GetHashCode.Call(obj)
    return int32(ret)
}

func Monitor_ToString(obj uintptr) string {
    ret, _, _ := monitor_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Monitor_GetHandle(obj uintptr) HMONITOR {
    ret, _, _ := monitor_GetHandle.Call(obj)
    return HMONITOR(ret)
}

func Monitor_GetMonitorNum(obj uintptr) int32 {
    ret, _, _ := monitor_GetMonitorNum.Call(obj)
    return int32(ret)
}

func Monitor_GetLeft(obj uintptr) int32 {
    ret, _, _ := monitor_GetLeft.Call(obj)
    return int32(ret)
}

func Monitor_GetHeight(obj uintptr) int32 {
    ret, _, _ := monitor_GetHeight.Call(obj)
    return int32(ret)
}

func Monitor_GetTop(obj uintptr) int32 {
    ret, _, _ := monitor_GetTop.Call(obj)
    return int32(ret)
}

func Monitor_GetWidth(obj uintptr) int32 {
    ret, _, _ := monitor_GetWidth.Call(obj)
    return int32(ret)
}

func Monitor_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    monitor_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Monitor_GetWorkareaRect(obj uintptr) TRect {
    var ret TRect
    monitor_GetWorkareaRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Monitor_GetPrimary(obj uintptr) bool {
    ret, _, _ := monitor_GetPrimary.Call(obj)
    return DBoolToGoBool(ret)
}

func Monitor_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := monitor_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

