
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Brush_Create() uintptr {
    ret, _, _ := brush_Create.Call()
    return ret
}

func Brush_Free(obj uintptr) {
    brush_Free.Call(obj)
}

func Brush_Assign(obj uintptr, Source uintptr)  {
    brush_Assign.Call(obj, Source )
}

func Brush_HandleAllocated(obj uintptr) bool {
    ret, _, _ := brush_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Brush_GetNamePath(obj uintptr) string {
    ret, _, _ := brush_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_ClassName(obj uintptr) string {
    ret, _, _ := brush_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := brush_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Brush_GetHashCode(obj uintptr) int32 {
    ret, _, _ := brush_GetHashCode.Call(obj)
    return int32(ret)
}

func Brush_ToString(obj uintptr) string {
    ret, _, _ := brush_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := brush_GetBitmap.Call(obj)
    return ret
}

func Brush_SetBitmap(obj uintptr, value uintptr) {
   brush_SetBitmap.Call(obj, value)
}

func Brush_GetHandle(obj uintptr) HBRUSH {
    ret, _, _ := brush_GetHandle.Call(obj)
    return HBRUSH(ret)
}

func Brush_SetHandle(obj uintptr, value HBRUSH) {
   brush_SetHandle.Call(obj, uintptr(value))
}

func Brush_GetColor(obj uintptr) TColor {
    ret, _, _ := brush_GetColor.Call(obj)
    return TColor(ret)
}

func Brush_SetColor(obj uintptr, value TColor) {
   brush_SetColor.Call(obj, uintptr(value))
}

func Brush_GetStyle(obj uintptr) TBrushStyle {
    ret, _, _ := brush_GetStyle.Call(obj)
    return TBrushStyle(ret)
}

func Brush_SetStyle(obj uintptr, value TBrushStyle) {
   brush_SetStyle.Call(obj, uintptr(value))
}

func Brush_SetOnChange(obj uintptr, fn interface{}) {
    brush_SetOnChange.Call(obj, addEventToMap(fn))
}

