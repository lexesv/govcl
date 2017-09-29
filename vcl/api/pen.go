
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Pen_Create() uintptr {
    ret, _, _ := pen_Create.Call()
    return ret
}

func Pen_Free(obj uintptr) {
    pen_Free.Call(obj)
}

func Pen_Assign(obj uintptr, Source uintptr)  {
    pen_Assign.Call(obj, Source )
}

func Pen_HandleAllocated(obj uintptr) bool {
    ret, _, _ := pen_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Pen_GetNamePath(obj uintptr) string {
    ret, _, _ := pen_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_ClassName(obj uintptr) string {
    ret, _, _ := pen_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pen_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Pen_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pen_GetHashCode.Call(obj)
    return int32(ret)
}

func Pen_ToString(obj uintptr) string {
    ret, _, _ := pen_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_GetHandle(obj uintptr) HPEN {
    ret, _, _ := pen_GetHandle.Call(obj)
    return HPEN(ret)
}

func Pen_SetHandle(obj uintptr, value HPEN) {
   pen_SetHandle.Call(obj, uintptr(value))
}

func Pen_GetColor(obj uintptr) TColor {
    ret, _, _ := pen_GetColor.Call(obj)
    return TColor(ret)
}

func Pen_SetColor(obj uintptr, value TColor) {
   pen_SetColor.Call(obj, uintptr(value))
}

func Pen_GetMode(obj uintptr) TPenMode {
    ret, _, _ := pen_GetMode.Call(obj)
    return TPenMode(ret)
}

func Pen_SetMode(obj uintptr, value TPenMode) {
   pen_SetMode.Call(obj, uintptr(value))
}

func Pen_GetStyle(obj uintptr) TPenStyle {
    ret, _, _ := pen_GetStyle.Call(obj)
    return TPenStyle(ret)
}

func Pen_SetStyle(obj uintptr, value TPenStyle) {
   pen_SetStyle.Call(obj, uintptr(value))
}

func Pen_GetWidth(obj uintptr) int32 {
    ret, _, _ := pen_GetWidth.Call(obj)
    return int32(ret)
}

func Pen_SetWidth(obj uintptr, value int32) {
   pen_SetWidth.Call(obj, uintptr(value))
}

func Pen_SetOnChange(obj uintptr, fn interface{}) {
    pen_SetOnChange.Call(obj, addEventToMap(fn))
}

