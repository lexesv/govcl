
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Margins_Create() uintptr {
    ret, _, _ := margins_Create.Call()
    return ret
}

func Margins_Free(obj uintptr) {
    margins_Free.Call(obj)
}

func Margins_SetBounds(obj uintptr, ALeft int32, ATop int32, ARight int32, ABottom int32)  {
    margins_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(ARight) , uintptr(ABottom) )
}

func Margins_Assign(obj uintptr, Source uintptr)  {
    margins_Assign.Call(obj, Source )
}

func Margins_GetNamePath(obj uintptr) string {
    ret, _, _ := margins_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_ClassName(obj uintptr) string {
    ret, _, _ := margins_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := margins_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Margins_GetHashCode(obj uintptr) int32 {
    ret, _, _ := margins_GetHashCode.Call(obj)
    return int32(ret)
}

func Margins_ToString(obj uintptr) string {
    ret, _, _ := margins_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_GetControlLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetControlLeft.Call(obj)
    return int32(ret)
}

func Margins_GetControlTop(obj uintptr) int32 {
    ret, _, _ := margins_GetControlTop.Call(obj)
    return int32(ret)
}

func Margins_GetControlWidth(obj uintptr) int32 {
    ret, _, _ := margins_GetControlWidth.Call(obj)
    return int32(ret)
}

func Margins_GetControlHeight(obj uintptr) int32 {
    ret, _, _ := margins_GetControlHeight.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Margins_SetOnChange(obj uintptr, fn interface{}) {
    margins_SetOnChange.Call(obj, addEventToMap(fn))
}

func Margins_GetLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetLeft.Call(obj)
    return int32(ret)
}

func Margins_SetLeft(obj uintptr, value int32) {
   margins_SetLeft.Call(obj, uintptr(value))
}

func Margins_GetTop(obj uintptr) int32 {
    ret, _, _ := margins_GetTop.Call(obj)
    return int32(ret)
}

func Margins_SetTop(obj uintptr, value int32) {
   margins_SetTop.Call(obj, uintptr(value))
}

func Margins_GetRight(obj uintptr) int32 {
    ret, _, _ := margins_GetRight.Call(obj)
    return int32(ret)
}

func Margins_SetRight(obj uintptr, value int32) {
   margins_SetRight.Call(obj, uintptr(value))
}

func Margins_GetBottom(obj uintptr) int32 {
    ret, _, _ := margins_GetBottom.Call(obj)
    return int32(ret)
}

func Margins_SetBottom(obj uintptr, value int32) {
   margins_SetBottom.Call(obj, uintptr(value))
}

