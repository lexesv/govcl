
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TColorDialog struct {
    IComponent
    instance uintptr
}

func NewColorDialog(owner IComponent) *TColorDialog {
    c := new(TColorDialog)
    c.instance = ColorDialog_Create(owner.Instance())
    return c
}

func ColorDialogFromInst(inst uintptr) *TColorDialog {
    c := new(TColorDialog)
    c.instance = inst
    return c
}

func ColorDialogFromObj(obj IObject) *TColorDialog {
    c := new(TColorDialog)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorDialog) Free() {
    if c.instance != 0 {
        ColorDialog_Free(c.instance)
    }
}

func (c *TColorDialog) Instance() uintptr {
    return c.instance
}

func (c *TColorDialog) IsValid() bool {
    return c.instance != 0
}

func (c *TColorDialog) Execute(ParentWnd HWND) bool {
    defer exceptionProc()
    return ColorDialog_Execute(c.instance, ParentWnd )
}

func (c *TColorDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorDialog_FindComponent(c.instance, AName ))
}

func (c *TColorDialog) GetNamePath() string {
    defer exceptionProc()
    return ColorDialog_GetNamePath(c.instance)
}

func (c *TColorDialog) HasParent() bool {
    defer exceptionProc()
    return ColorDialog_HasParent(c.instance)
}

func (c *TColorDialog) Assign(Source IObject) {
    defer exceptionProc()
    ColorDialog_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorDialog) ClassName() string {
    defer exceptionProc()
    return ColorDialog_ClassName(c.instance)
}

func (c *TColorDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ColorDialog_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorDialog) GetHashCode() int32 {
    defer exceptionProc()
    return ColorDialog_GetHashCode(c.instance)
}

func (c *TColorDialog) ToString() string {
    defer exceptionProc()
    return ColorDialog_ToString(c.instance)
}

func (c *TColorDialog) Color() TColor {
    defer exceptionProc()
    return ColorDialog_GetColor(c.instance)
}

func (c *TColorDialog) SetColor(value TColor) {
    defer exceptionProc()
    ColorDialog_SetColor(c.instance, value)
}

func (c *TColorDialog) Options() TColorDialogOptions {
    defer exceptionProc()
    return ColorDialog_GetOptions(c.instance)
}

func (c *TColorDialog) SetOptions(value TColorDialogOptions) {
    defer exceptionProc()
    ColorDialog_SetOptions(c.instance, value)
}

func (c *TColorDialog) Handle() HWND {
    defer exceptionProc()
    return ColorDialog_GetHandle(c.instance)
}

func (c *TColorDialog) SetOnClose(fn TNotifyEvent) {
    ColorDialog_SetOnClose(c.instance, fn)
}

func (c *TColorDialog) SetOnShow(fn TNotifyEvent) {
    ColorDialog_SetOnShow(c.instance, fn)
}

func (c *TColorDialog) ComponentCount() int32 {
    defer exceptionProc()
    return ColorDialog_GetComponentCount(c.instance)
}

func (c *TColorDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return ColorDialog_GetComponentIndex(c.instance)
}

func (c *TColorDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    ColorDialog_SetComponentIndex(c.instance, value)
}

func (c *TColorDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorDialog_GetOwner(c.instance))
}

func (c *TColorDialog) Name() string {
    defer exceptionProc()
    return ColorDialog_GetName(c.instance)
}

func (c *TColorDialog) SetName(value string) {
    defer exceptionProc()
    ColorDialog_SetName(c.instance, value)
}

func (c *TColorDialog) Tag() int {
    defer exceptionProc()
    return ColorDialog_GetTag(c.instance)
}

func (c *TColorDialog) SetTag(value int) {
    defer exceptionProc()
    ColorDialog_SetTag(c.instance, value)
}

func (c *TColorDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorDialog_GetComponents(c.instance, AIndex))
}

