
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

type TFontDialog struct {
    IComponent
    instance uintptr
}

func NewFontDialog(owner IComponent) *TFontDialog {
    f := new(TFontDialog)
    f.instance = FontDialog_Create(owner.Instance())
    return f
}

func FontDialogFromInst(inst uintptr) *TFontDialog {
    f := new(TFontDialog)
    f.instance = inst
    return f
}

func FontDialogFromObj(obj IObject) *TFontDialog {
    f := new(TFontDialog)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFontDialog) Free() {
    if f.instance != 0 {
        FontDialog_Free(f.instance)
    }
}

func (f *TFontDialog) Instance() uintptr {
    return f.instance
}

func (f *TFontDialog) IsValid() bool {
    return f.instance != 0
}

func (f *TFontDialog) Execute(ParentWnd HWND) bool {
    defer exceptionProc()
    return FontDialog_Execute(f.instance, ParentWnd )
}

func (f *TFontDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(FontDialog_FindComponent(f.instance, AName ))
}

func (f *TFontDialog) GetNamePath() string {
    defer exceptionProc()
    return FontDialog_GetNamePath(f.instance)
}

func (f *TFontDialog) HasParent() bool {
    defer exceptionProc()
    return FontDialog_HasParent(f.instance)
}

func (f *TFontDialog) Assign(Source IObject) {
    defer exceptionProc()
    FontDialog_Assign(f.instance, CheckPtr(Source))
}

func (f *TFontDialog) ClassName() string {
    defer exceptionProc()
    return FontDialog_ClassName(f.instance)
}

func (f *TFontDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return FontDialog_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFontDialog) GetHashCode() int32 {
    defer exceptionProc()
    return FontDialog_GetHashCode(f.instance)
}

func (f *TFontDialog) ToString() string {
    defer exceptionProc()
    return FontDialog_ToString(f.instance)
}

func (f *TFontDialog) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(FontDialog_GetFont(f.instance))
}

func (f *TFontDialog) SetFont(value *TFont) {
    defer exceptionProc()
    FontDialog_SetFont(f.instance, CheckPtr(value))
}

func (f *TFontDialog) Options() TFontDialogOptions {
    defer exceptionProc()
    return FontDialog_GetOptions(f.instance)
}

func (f *TFontDialog) SetOptions(value TFontDialogOptions) {
    defer exceptionProc()
    FontDialog_SetOptions(f.instance, value)
}

func (f *TFontDialog) Handle() HWND {
    defer exceptionProc()
    return FontDialog_GetHandle(f.instance)
}

func (f *TFontDialog) SetOnClose(fn TNotifyEvent) {
    FontDialog_SetOnClose(f.instance, fn)
}

func (f *TFontDialog) SetOnShow(fn TNotifyEvent) {
    FontDialog_SetOnShow(f.instance, fn)
}

func (f *TFontDialog) ComponentCount() int32 {
    defer exceptionProc()
    return FontDialog_GetComponentCount(f.instance)
}

func (f *TFontDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return FontDialog_GetComponentIndex(f.instance)
}

func (f *TFontDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    FontDialog_SetComponentIndex(f.instance, value)
}

func (f *TFontDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(FontDialog_GetOwner(f.instance))
}

func (f *TFontDialog) Name() string {
    defer exceptionProc()
    return FontDialog_GetName(f.instance)
}

func (f *TFontDialog) SetName(value string) {
    defer exceptionProc()
    FontDialog_SetName(f.instance, value)
}

func (f *TFontDialog) Tag() int {
    defer exceptionProc()
    return FontDialog_GetTag(f.instance)
}

func (f *TFontDialog) SetTag(value int) {
    defer exceptionProc()
    FontDialog_SetTag(f.instance, value)
}

func (f *TFontDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(FontDialog_GetComponents(f.instance, AIndex))
}

