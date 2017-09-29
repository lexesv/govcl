
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TBrush struct {
    IObject
    instance uintptr
}

func NewBrush() *TBrush {
    b := new(TBrush)
    b.instance = Brush_Create()
    return b
}

func BrushFromInst(inst uintptr) *TBrush {
    b := new(TBrush)
    b.instance = inst
    return b
}

func BrushFromObj(obj IObject) *TBrush {
    b := new(TBrush)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBrush) Free() {
    if b.instance != 0 {
        Brush_Free(b.instance)
    }
}

func (b *TBrush) Instance() uintptr {
    return b.instance
}

func (b *TBrush) IsValid() bool {
    return b.instance != 0
}

func (b *TBrush) Assign(Source IObject) {
    defer exceptionProc()
    Brush_Assign(b.instance, CheckPtr(Source))
}

func (b *TBrush) HandleAllocated() bool {
    defer exceptionProc()
    return Brush_HandleAllocated(b.instance)
}

func (b *TBrush) GetNamePath() string {
    defer exceptionProc()
    return Brush_GetNamePath(b.instance)
}

func (b *TBrush) ClassName() string {
    defer exceptionProc()
    return Brush_ClassName(b.instance)
}

func (b *TBrush) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Brush_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBrush) GetHashCode() int32 {
    defer exceptionProc()
    return Brush_GetHashCode(b.instance)
}

func (b *TBrush) ToString() string {
    defer exceptionProc()
    return Brush_ToString(b.instance)
}

func (b *TBrush) Bitmap() *TBitmap {
    defer exceptionProc()
    return BitmapFromInst(Brush_GetBitmap(b.instance))
}

func (b *TBrush) SetBitmap(value *TBitmap) {
    defer exceptionProc()
    Brush_SetBitmap(b.instance, CheckPtr(value))
}

func (b *TBrush) Handle() HBRUSH {
    defer exceptionProc()
    return Brush_GetHandle(b.instance)
}

func (b *TBrush) SetHandle(value HBRUSH) {
    defer exceptionProc()
    Brush_SetHandle(b.instance, value)
}

func (b *TBrush) Color() TColor {
    defer exceptionProc()
    return Brush_GetColor(b.instance)
}

func (b *TBrush) SetColor(value TColor) {
    defer exceptionProc()
    Brush_SetColor(b.instance, value)
}

func (b *TBrush) Style() TBrushStyle {
    defer exceptionProc()
    return Brush_GetStyle(b.instance)
}

func (b *TBrush) SetStyle(value TBrushStyle) {
    defer exceptionProc()
    Brush_SetStyle(b.instance, value)
}

func (b *TBrush) SetOnChange(fn TNotifyEvent) {
    Brush_SetOnChange(b.instance, fn)
}

