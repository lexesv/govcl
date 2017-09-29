
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPen struct {
    IObject
    instance uintptr
}

func NewPen() *TPen {
    p := new(TPen)
    p.instance = Pen_Create()
    return p
}

func PenFromInst(inst uintptr) *TPen {
    p := new(TPen)
    p.instance = inst
    return p
}

func PenFromObj(obj IObject) *TPen {
    p := new(TPen)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPen) Free() {
    if p.instance != 0 {
        Pen_Free(p.instance)
    }
}

func (p *TPen) Instance() uintptr {
    return p.instance
}

func (p *TPen) IsValid() bool {
    return p.instance != 0
}

func (p *TPen) Assign(Source IObject) {
    defer exceptionProc()
    Pen_Assign(p.instance, CheckPtr(Source))
}

func (p *TPen) HandleAllocated() bool {
    defer exceptionProc()
    return Pen_HandleAllocated(p.instance)
}

func (p *TPen) GetNamePath() string {
    defer exceptionProc()
    return Pen_GetNamePath(p.instance)
}

func (p *TPen) ClassName() string {
    defer exceptionProc()
    return Pen_ClassName(p.instance)
}

func (p *TPen) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Pen_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPen) GetHashCode() int32 {
    defer exceptionProc()
    return Pen_GetHashCode(p.instance)
}

func (p *TPen) ToString() string {
    defer exceptionProc()
    return Pen_ToString(p.instance)
}

func (p *TPen) Handle() HPEN {
    defer exceptionProc()
    return Pen_GetHandle(p.instance)
}

func (p *TPen) SetHandle(value HPEN) {
    defer exceptionProc()
    Pen_SetHandle(p.instance, value)
}

func (p *TPen) Color() TColor {
    defer exceptionProc()
    return Pen_GetColor(p.instance)
}

func (p *TPen) SetColor(value TColor) {
    defer exceptionProc()
    Pen_SetColor(p.instance, value)
}

func (p *TPen) Mode() TPenMode {
    defer exceptionProc()
    return Pen_GetMode(p.instance)
}

func (p *TPen) SetMode(value TPenMode) {
    defer exceptionProc()
    Pen_SetMode(p.instance, value)
}

func (p *TPen) Style() TPenStyle {
    defer exceptionProc()
    return Pen_GetStyle(p.instance)
}

func (p *TPen) SetStyle(value TPenStyle) {
    defer exceptionProc()
    Pen_SetStyle(p.instance, value)
}

func (p *TPen) Width() int32 {
    defer exceptionProc()
    return Pen_GetWidth(p.instance)
}

func (p *TPen) SetWidth(value int32) {
    defer exceptionProc()
    Pen_SetWidth(p.instance, value)
}

func (p *TPen) SetOnChange(fn TNotifyEvent) {
    Pen_SetOnChange(p.instance, fn)
}

