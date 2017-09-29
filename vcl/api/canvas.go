
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Canvas_Create() uintptr {
    ret, _, _ := canvas_Create.Call()
    return ret
}

func Canvas_Free(obj uintptr) {
    canvas_Free.Call(obj)
}

func Canvas_Arc(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Arc.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_ArcTo(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_ArcTo.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_AngleArc(obj uintptr, X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32)  {
    canvas_AngleArc.Call(obj, uintptr(X) , uintptr(Y) , uintptr(Radius) , uintptr(unsafe.Pointer(&StartAngle)), uintptr(unsafe.Pointer(&SweepAngle)))
}

func Canvas_Chord(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Chord.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_Ellipse(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32)  {
    canvas_Ellipse.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) )
}

func Canvas_FloodFill(obj uintptr, X int32, Y int32, Color TColor, FillStyle TFillStyle)  {
    canvas_FloodFill.Call(obj, uintptr(X) , uintptr(Y) , uintptr(Color) , uintptr(FillStyle) )
}

func Canvas_HandleAllocated(obj uintptr) bool {
    ret, _, _ := canvas_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Canvas_LineTo(obj uintptr, X int32, Y int32)  {
    canvas_LineTo.Call(obj, uintptr(X) , uintptr(Y) )
}

func Canvas_MoveTo(obj uintptr, X int32, Y int32)  {
    canvas_MoveTo.Call(obj, uintptr(X) , uintptr(Y) )
}

func Canvas_Pie(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Pie.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_Rectangle(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32)  {
    canvas_Rectangle.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) )
}

func Canvas_Refresh(obj uintptr)  {
    canvas_Refresh.Call(obj)
}

func Canvas_RoundRect(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32)  {
    canvas_RoundRect.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) )
}

func Canvas_TextExtent(obj uintptr, Text string) TSize {
    var ret TSize
    canvas_TextExtent.Call(obj, GoStrToDStr(Text) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Canvas_TextOut(obj uintptr, X int32, Y int32, Text string)  {
    canvas_TextOut.Call(obj, uintptr(X) , uintptr(Y) , GoStrToDStr(Text) )
}

func Canvas_Assign(obj uintptr, Source uintptr)  {
    canvas_Assign.Call(obj, Source )
}

func Canvas_GetNamePath(obj uintptr) string {
    ret, _, _ := canvas_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_ClassName(obj uintptr) string {
    ret, _, _ := canvas_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := canvas_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Canvas_GetHashCode(obj uintptr) int32 {
    ret, _, _ := canvas_GetHashCode.Call(obj)
    return int32(ret)
}

func Canvas_ToString(obj uintptr) string {
    ret, _, _ := canvas_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_GetHandle(obj uintptr) HDC {
    ret, _, _ := canvas_GetHandle.Call(obj)
    return HDC(ret)
}

func Canvas_SetHandle(obj uintptr, value HDC) {
   canvas_SetHandle.Call(obj, uintptr(value))
}

func Canvas_GetBrush(obj uintptr) uintptr {
    ret, _, _ := canvas_GetBrush.Call(obj)
    return ret
}

func Canvas_SetBrush(obj uintptr, value uintptr) {
   canvas_SetBrush.Call(obj, value)
}

func Canvas_GetCopyMode(obj uintptr) int32 {
    ret, _, _ := canvas_GetCopyMode.Call(obj)
    return int32(ret)
}

func Canvas_SetCopyMode(obj uintptr, value int32) {
   canvas_SetCopyMode.Call(obj, uintptr(value))
}

func Canvas_GetFont(obj uintptr) uintptr {
    ret, _, _ := canvas_GetFont.Call(obj)
    return ret
}

func Canvas_SetFont(obj uintptr, value uintptr) {
   canvas_SetFont.Call(obj, value)
}

func Canvas_GetPen(obj uintptr) uintptr {
    ret, _, _ := canvas_GetPen.Call(obj)
    return ret
}

func Canvas_SetPen(obj uintptr, value uintptr) {
   canvas_SetPen.Call(obj, value)
}

func Canvas_SetOnChange(obj uintptr, fn interface{}) {
    canvas_SetOnChange.Call(obj, addEventToMap(fn))
}

